package main

import (
	"fmt"
	"github.com/chrislusf/glow/flow"
	"github.com/chrislusf/glow/source/hdfs"
	"github.com/wilseypa/rphash-golang/parse"
	"github.com/wilseypa/rphash-golang/reader"
	"github.com/wilseypa/rphash-golang/stream"
	"github.com/wilseypa/rphash-golang/types"
	"github.com/wilseypa/rphash-golang/utils"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"strconv"
	"time"
)

var (
	f                            = flow.New()
	expectedDimensions           = -1
	app                          = kingpin.New("rphash-golang", "A command-line distributed clusterer.")
	clusterer                    = app.Command("stream", "Start a streaming cluster process")
	clustererHdfsEnable          = clusterer.Flag("hdfs.enable", "Enable hdfs").Default("false").Bool()
	clustererHdfsDir             = clusterer.Flag("hdfs.dir", "Path of hdfs data").Default("").String()
	clustererCluster             = clusterer.Flag("cluster", "Type of clusterer to use").Default("rphash").String()
	clustererClusters            = clusterer.Flag("num.clusters", "Number of clusters").Default("4").Int()
	clustererShards              = clusterer.Flag("num.shards", "Number of shards").Default("8").Int()
	clustererLocalInputFile      = clusterer.Flag("local.file", "Path of the local input file").Default("").String()
	clustererOutputPlots         = clusterer.Flag("centroid.plots", "Output centroid dimension plots").Default("false").Bool()
	clustererOutputPlotsFileName = clusterer.Flag("centroid.plots.filename", "Output centroid dimension plots file name").Default("plots").String()
)

type Vector struct {
	Data []float64
}

func main() {

	// Parse commands
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case clusterer.FullCommand():

		var centroids []types.Centroid
    var normalizedResults [][]float64
		var dataset *flow.Dataset
    t1 := time.Now()

		switch *clustererCluster {
		case "rphash":

			// Create a streaming rphash object
			var rphashObject *reader.StreamObject
			var rphashStream *stream.Stream
			// Split the data into shards and send them to the Agents to work on.
			if !*clustererHdfsEnable {
        if *clustererLocalInputFile == "" {
          panic("Must include --local.filename")
        }

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
				}, *clustererShards)
			} else {
        if *clustererHdfsDir == "" {
          panic("Must include --hdfs.dir")
        }
				dataset := hdfs.Source(
					f,
					*clustererHdfsDir,
					*clustererShards,
				)
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
			fmt.Println("Streaming RPHash Time: ", time.Since(t1))

		case "streaming-kmeans":

      var kmeansStream

      dataset = f.Source(func(out chan Vector) {
        records, err := utils.ReadLines(*clustererLocalInputFile)
        if err != nil {
          panic(err)
        }

        // Convert the record to standard floating points.
        for i, record := range records {
          if i == 0 {
            // Create a new RPHash stream.
            kmeansStream = clusterer.NewKMeansStream(*clustererClusters, *clustererClusters, len(record))
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
				kmeansStream.AddDataPoint(vector.Data);
			}).Run()

      normalizedResults = kmeansStream.GetCentroids()
			fmt.Println("Streaming KMeans Time: ", time.Since(t1))
		}

    // DeNormalize results
    fmt.Println("Preparing results...")
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
      fmt.Println("Plotting...")
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
        labels[i] = "Centroid " + strconv.FormatInt(int64(i), 16)
      }
      utils.GeneratePlots(xPlotValues, yPlotValues, "Centroid", "Dimension", "Dimension Value", "plots/"+(*clustererOutputPlotsFileName)+"-", labels)
    } else {
      fmt.Println("Results: ", denormalizedResults)
    }
	}
}
