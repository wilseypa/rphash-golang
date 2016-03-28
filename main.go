package main

import (
  "encoding/csv"
  "github.com/chrislusf/glow/flow"
  "github.com/wilseypa/rphash-golang/reader"
  "github.com/wilseypa/rphash-golang/stream"
  "github.com/wilseypa/rphash-golang/types"
   _ "github.com/chrislusf/glow/driver"
  "log"
  "io"
  "os"
  "strconv"
)

var (
  dataFilePath       = "./dist/dataset.csv"
  f                  = flow.New()
  expectedDimensions = -1
  numClusters        = 4
  numShards          = 3
)

type Vector struct {
  Data []float64
}

func main() {
  var rphashObject *reader.StreamObject
  var rphashStream *stream.Stream
  var centroids []types.Centroid
  // Split the data into shards and send them to the Agents to work on.
  f.Source(func(out chan Vector) {
    // Read the csv data file...
    csvfile, err := os.OpenFile(dataFilePath, os.O_RDWR, os.ModeAppend)
    defer csvfile.Close()

    if err != nil {
      log.Println(err)
      os.Exit(1)
    }

    csvReader := csv.NewReader(csvfile)

    // Convert the record to standard floating points.
    i := -1
    for {
      i++;
      record, err := csvReader.Read()
      if err != nil {
        if err == io.EOF {
          break
        }
        log.Println(err)
        os.Exit(1)
      }

      if i == 0 {
        // Create a new RPHash stream.
        rphashObject = reader.NewStreamObject(len(record), numClusters)
        rphashStream = stream.NewStream(rphashObject)
        rphashStream.RunCount = 1
        continue
      }
      data := make([]float64, len(record))
      for j, entry := range record {
        f, err := strconv.ParseFloat(entry, 64)
        if err != nil {
          log.Println(err)
          os.Exit(1)
        }
        data[j] = f
      }
      out <- Vector{data}
    }
  }, numShards).Map(func(vec Vector) {
    centroids = append(centroids, rphashStream.AddVectorOnlineStep(vec.Data))
  }).Run()

  for _, cent := range centroids {
    rphashStream.CentroidCounter.Add(cent)
  }

  results := rphashStream.GetCentroids()
  log.Println(results)
}
