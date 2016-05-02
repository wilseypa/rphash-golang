package main

import (
	"fmt"
	_ "github.com/chrislusf/glow/driver"
	"github.com/chrislusf/glow/flow"
	"github.com/chrislusf/glow/source/hdfs"
	clust "github.com/wilseypa/rphash-golang/clusterer"
	"github.com/wilseypa/rphash-golang/parse"
	"github.com/wilseypa/rphash-golang/reader"
	"github.com/wilseypa/rphash-golang/stream"
	"github.com/wilseypa/rphash-golang/types"
	"github.com/wilseypa/rphash-golang/utils"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strconv"
	"time"
)

var (
	f                            = flow.New()
	expectedDimensions           = -1
	app                          = kingpin.New("rphash-golang", "A command-line distributed clusterer.")
	clustererHdfsEnable          = app.Flag("hdfs.enable", "Enable hdfs").Default("false").Bool()
	clustererVerbose             = app.Flag("verbose", "Verbose output").Default("false").Bool()
	clustererHdfsDir             = app.Flag("hdfs.dir", "Path of hdfs data").Default("").String()
	clustererThreshold           = app.Flag("threshold", "Paint threshold").Default("0.5").Float()
	clustererCluster             = app.Flag("cluster", "Type of clusterer to use").Default("rphash").String()
	clustererClusters            = app.Flag("num.clusters", "Number of clusters").Default("4").Int()
	clustererShards              = app.Flag("num.shards", "Number of shards").Default("8").Int()
	clustererLocalInputFile      = app.Flag("local.file", "Path of the local input file").Required().String()
	clustererOutputPlots         = app.Flag("centroid.plots", "Output centroid dimension plots").Default("false").Bool()
	clustererOutputPlotsFileName = app.Flag("centroid.plots.file", "Output centroid dimension plots file name").Default("").String()
	clustererOutputPaintFile     = app.Flag("centroid.paint", "Output centroid dimension plots matrix file name").Default("").String()
	clustererOutputHeatFile      = app.Flag("centroid.heat", "Output centroid dimension plots heatmap file name").Default("").String()
)

type Vector struct {
	Data []float64
}

func main() {

	app.Parse(os.Args[1:])

	var normalizedResults [][]float64
	switch *clustererCluster {
	case "rphash":

		// Create a streaming rphash object
		var rphashObject *reader.StreamObject
		var rphashStream *stream.Stream
		var dataset *flow.Dataset
		var centroids []types.Centroid

		t1 := time.Now()
		// Split the data into shards and send them to the Agents to work on.
		if !*clustererHdfsEnable {

			dataset = f.Source(func(out chan Vector) {
				records, err := utils.ReadLines(*clustererLocalInputFile)
				if err != nil {
					panic(err)
				}

				// Convert the record to standard floating points.
				for i, record := range records {

					if i == 0 {
						// Create a new RPHash stream.
						rphashObject = reader.NewStreamObject(len(record), *clustererClusters)
						rphashStream = stream.NewStream(rphashObject)
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
			}, *clustererShards)
		} else {
			if *clustererHdfsDir == "" {
				panic("Must include --hdfs.dir")
			}
			dataset = hdfs.Source(
				f,
				*clustererHdfsDir,
				*clustererShards,
			).Map(func(record string, out chan Vector) {
				data := make([]float64, len(record))
				for j, entry := range record {
					f, err := strconv.ParseFloat(string(entry), 64)
					f = parse.Normalize(f)
					if err != nil {
						panic(err)
					}
					data[j] = f
				}
				out <- Vector{Data: data}
			})
		}

		dataset.Map(func(vec Vector) {
			centroids = append(centroids, rphashStream.AddVectorOnlineStep(vec.Data))
		}).Run()

		// Locally reduce
		for _, cent := range centroids {
			rphashStream.CentroidCounter.Add(cent)
		}

		// Get the centroids.
		normalizedResults = rphashStream.GetCentroids()
		if *clustererVerbose {
			fmt.Println("Streaming RPHash Time: ", time.Since(t1))
		}

	case "streaming-kmeans":

		var kmeansStream *clust.KMeansStream
		var dataset *flow.Dataset

		t1 := time.Now()

		dataset = f.Source(func(out chan Vector) {
			records, err := utils.ReadLines(*clustererLocalInputFile)
			if err != nil {
				panic(err)
			}

			// Convert the record to standard floating points.
			for i, record := range records {
				if i == 0 {
					// Create a new RPHash stream.
					kmeansStream = clust.NewKMeansStream(*clustererClusters, *clustererClusters, len(record))
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
		}, *clustererShards)

		dataset.Map(func(vec Vector) {
			kmeansStream.AddDataPoint(vec.Data)
		}).Run()

		normalizedResults = kmeansStream.GetCentroids()
    if *clustererVerbose {
  		fmt.Println("Streaming KMeans Time: ", time.Since(t1))
    }
	}

	// DeNormalize results
  if *clustererVerbose {
  	fmt.Println("Preparing results...")
  }
	denormalizedResults := make([][]float64, len(normalizedResults))
	for i, result := range normalizedResults {
		row := make([]float64, len(result))
		for j, dimension := range result {
			row[j] = parse.DeNormalize(dimension)
		}
		denormalizedResults[i] = row
	}

	if *clustererOutputPlots {
		// Prepare reults into plots
		if *clustererVerbose {
			fmt.Println("Plotting...")
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

			if *clustererOutputPaintFile != "" {
				utils.Paint(result, i, *clustererOutputPaintFile, *clustererThreshold)
			}
			if *clustererOutputHeatFile != "" {
				utils.HeatMap(result, i, *clustererOutputHeatFile)
			}

			labels[i] = "Centroid " + strconv.FormatInt(int64(i), 16)
		}
		utils.GeneratePlots(xPlotValues, yPlotValues, "Centroid", "Dimension", "Dimension Value", (*clustererOutputPlotsFileName)+"-", labels)
	} else {
		if *clustererVerbose {
			fmt.Println("Results: ", denormalizedResults)
		}

	}
}
