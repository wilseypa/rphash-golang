package main

import (
  "encoding/csv"
  "github.com/chrislusf/glow/flow"
  "github.com/wilseypa/rphash-golang/reader"
  "github.com/wilseypa/rphash-golang/stream"
  "github.com/wilseypa/rphash-golang/types"
   _ "github.com/chrislusf/glow/driver"
  "log"
  "os"
  "strconv"
)

var (
  dataFilePath       = "/Users/sam.wenke/data/FL_insurance_sample_small.csv"
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
    records, err := csvReader.ReadAll()

    if err != nil {
      log.Println(err)
      os.Exit(1)
    }

    // Convert the record to standard floating points.
    for i, individualRecord := range records {
      individualRecord = individualRecord[3:14]
      if i == 0 {
        // Create a new RPHash stream.
        rphashObject = reader.NewStreamObject(len(individualRecord), numClusters)
        rphashStream = stream.NewStream(rphashObject)
        rphashStream.RunCount = 1
        continue
      }
      data := make([]float64, len(individualRecord))
      for j, entry := range individualRecord {
        f, err := strconv.ParseFloat(entry, 64)
        if err != nil {
          log.Println(err)
          os.Exit(1)
        }
        data[j] = f
      }
      out <- Vector{data}
    }
  }, numShards).Map(func(vec Vector, out chan types.Centroid) {
    out <- rphashStream.AddVectorOnlineStep(vec.Data, nil)
  }).Map(func(cent types.Centroid) {
    rphashStream.CentroidCounter.Add(cent)
  }).Run()

  centroids := rphashStream.GetCentroids()
  log.Println(centroids)
}
