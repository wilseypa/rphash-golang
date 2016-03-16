package main

import (
  "os"
  "log"
  "io"
  "strconv"
  "encoding/csv"
  "github.com/wilseypa/rphash-golang/stream"
  "github.com/wilseypa/rphash-golang/reader"
  "github.com/wilseypa/rphash-golang/types"
  "github.com/chrislusf/glow/flow"
)

var (
  dataFilePath = "/testData/rphash/labelled/HAR/dataset.csv"
  f = flow.New()
  expectedDimensions = -1
  numClusters = 4
)

type Vector struct {
	Data  []float64
}

// Read the CSV data into memory.
func ReadCSVData(filePath string) *csv.Reader {
  csvfile, err := os.Open(filePath)
  if err != nil {
    log.Println(err)
    os.Exit(1)
  }
  defer csvfile.Close()
  reader := csv.NewReader(csvfile)
  // Expected dimensions.
  reader.FieldsPerRecord = expectedDimensions
  return reader
}

func main() {

  var rphashObject *reader.StreamObject
  var rphashStream *stream.Stream

  // Split the data into shards and send them to the Agents to work on.
  f.Source(func(out chan Vector) {
    // Read the csv data file...
		CSVReader := ReadCSVData(dataFilePath)
    i := -1
    for {
      i++
      record, err := CSVReader.Read()
      if err == io.EOF {
        break
      }

      // For the first record, store the features.
      if i == 0 {
        rphashObject = reader.NewStreamObject(len(record), numClusters)
        rphashStream = stream.NewStream(rphashObject)
        rphashStream.RunCount = 1
      }

  		if err != nil {
  			log.Println(err)
        os.Exit(1)
  		}

      // Convert the record to standard floating points.
      data := make([]float64, len(record))
      for j, entry := range record {
        f, err := strconv.ParseFloat(entry, 64)
        if err != nil {
          log.Println(err)
          os.Exit(1);
        }
        data[j] = f
      }
      out <- Vector{data}
  	}
	}, 3).Map(func(vec Vector, out chan types.Centroid) {
    out <- rphashStream.AddVectorOnlineStep(vec.Data, nil)
	}).Map(func(cent types.Centroid, out chan int) {
    rphashStream.CentroidCounter.Add(cent)
    out <- 1
	}).Map(func(x int) [][]float64 {
    centroids := rphashStream.GetCentroids()
    log.Println(centroids)
    return centroids
  }).Run()
}
