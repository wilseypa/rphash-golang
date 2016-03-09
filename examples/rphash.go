package main

import (
  "github.com/wilseypa/rphash-golang/api"
  "github.com/wilseypa/rphash-golang/parse"
  "io/ioutil"
)

var numberOfClusters = 4

const (
  exampleInputFileName  = "input.json"
  exampleOutputFileName = "output.json"
  exampleDataLabel      = "people"
)

func main() {
  parser := parse.NewParser()
  bytes, _ := ioutil.ReadFile(exampleInputFileName)
  jsonData := parser.BytesToJSON(bytes)
  data := parser.JSONToFloat64Matrix(exampleDataLabel, jsonData)
  cluster := api.NewRPHash(data, numberOfClusters)

  topCentroids := cluster.GetCentroids()

  jsonCentroids := parser.Float64MatrixToJSON(exampleDataLabel, topCentroids)

  jsonBytes := parser.JSONToBytes(jsonCentroids)
  err := ioutil.WriteFile(exampleOutputFileName, jsonBytes, 0644)
  if err != nil {
    panic(err)
  }
}
