package main

import (
	"log"
	"math"
	"time"

	"github.com/chrislusf/glow/flow"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
	"github.com/wilseypa/rphash-golang/parse"
	"github.com/wilseypa/rphash-golang/reader"
	"github.com/wilseypa/rphash-golang/stream"
	"github.com/wilseypa/rphash-golang/types"
	"github.com/wilseypa/rphash-golang/utils"
	//  _ "github.com/chrislusf/glow/driver"

	"strconv"
)

var (
	dataFilePath       = "data/MNISTnumImages5000.txt"
	f                  = flow.New()
	expectedDimensions = -1
	numClusters        = 10
	numShards          = 8
)

type Vector struct {
	Data []float64
}

func GeneratePlots(x, y [][]float64, title, xLabel, yLabel, fileName string, legendLabel []string) {
	outPlotPoints := make([]plotter.XYs, len(x))
	outPlots := make([]*plot.Plot, len(x))

	for i, _ := range outPlotPoints {
		outPlot, err := plot.New()
		outPlots[i] = outPlot
		outPlots[i].Title.Text = title
		outPlots[i].X.Label.Text = xLabel
		outPlots[i].Y.Label.Text = yLabel
		outPlotPoints[i] = make(plotter.XYs, len(x[0]))
		for j, _ := range x[0] {
			outPlotPoints[i][j].X = x[i][j]
			outPlotPoints[i][j].Y = y[i][j]
		}
		err = plotutil.AddLines(outPlots[i],
			legendLabel[i], outPlotPoints[i])
		if err != nil {
			panic(err)
		}

		if err = outPlot.Save(6*vg.Inch, 6*vg.Inch, (fileName+strconv.FormatInt(int64(i), 16))+".png"); err != nil {
			panic(err)
		}
	}
}

// 784 Bits
func Paint(image []float64, imageId int) {
	outPlotPoints := make(plotter.XYs, len(image))
	outPlot, err := plot.New()
	if err != nil {
		panic(err)
	}
	x := 0
	y := 0
	for i, bit := range image {
		outPlotPoints[i].X = float64(x)
		if bit > 0.4 {
			outPlotPoints[i].Y = float64(y)
		} else {
			outPlotPoints[i].Y = 0
		}
		if i%int(math.Sqrt(float64(len(image)))) == 0 {
			x = 0
			y++
		} else {
			x++
		}
	}
	outPlot.Add(plotter.NewGrid())
	s, _ := plotter.NewScatter(outPlotPoints)
	outPlot.Add(s)
	if err = outPlot.Save(6*vg.Inch, 6*vg.Inch, "plots/centroid-drawing-"+strconv.FormatInt(int64(imageId), 16)+".png"); err != nil {
		panic(err)
	}
}

func main() {
	var rphashObject *reader.StreamObject
	var rphashStream *stream.Stream
	var centroids []types.Centroid
	t1 := time.Now()
	// Split the data into shards and send them to the Agents to work on.
	f.Source(func(out chan Vector) {
		records, err := utils.ReadLines(dataFilePath)
		if err != nil {
			panic(err)
		}
		// Convert the record to standard floating points.
		for i, record := range records {
			if i == 0 {
				// Create a new RPHash stream.
				rphashObject = reader.NewStreamObject(len(record), numClusters)
				rphashStream = stream.NewStream(rphashObject)
				rphashStream.RunCount = 1
			}
			data := make([]float64, len(record))
			for j, entry := range record {
				f, err := strconv.ParseFloat(entry, 64)
				f = parse.Normalize(f)
				if err != nil {
					panic(err)
				}
				data[j] = f
			}
			out <- Vector{Data: data}
		}
	}, numShards).Map(func(vec Vector) {
		centroids = append(centroids, rphashStream.AddVectorOnlineStep(vec.Data))
	}).Run()

	for _, cent := range centroids {
		rphashStream.CentroidCounter.Add(cent)
	}
	normalizedResults := rphashStream.GetCentroids()
	t2 := time.Now()
	log.Println("Time: ", t2.Sub(t1))

	denormalizedResults := make([][]float64, len(normalizedResults))
	for i, result := range normalizedResults {
		row := make([]float64, len(result))
		for j, dimension := range result {
			row[j] = parse.DeNormalize(dimension)
		}
		denormalizedResults[i] = row
	}
	labels := make([]string, len(denormalizedResults))
	xPlotValues := make([][]float64, len(denormalizedResults))
	yPlotValues := make([][]float64, len(denormalizedResults))
	for i, result := range denormalizedResults {
		xPlotValues[i] = make([]float64, len(result))
		yPlotValues[i] = make([]float64, len(result))
		for j, val := range result {
			xPlotValues[i][j] = float64(j)
			yPlotValues[i][j] = val
		}
		Paint(result, i)
		sI := strconv.FormatInt(int64(i), 16)
		labels[i] = "Digit " + sI + " (by Classifier Centroid)"
	}
	GeneratePlots(xPlotValues, yPlotValues, "High Dimension Handwritting Digits 0-9 Classification", "Dimension", "Strength of Visual Pixel Recognition (0-1000)", "plots/centroid-dimensions-", labels)
}
