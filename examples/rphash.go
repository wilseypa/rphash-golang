package main;

import (
  "fmt"
  "io/ioutil"
  "github.com/wenkesj/rphash/api"
  "github.com/wenkesj/rphash/parse"
);

var numberOfClusters = 2;

const (
  exampleDataFileName = "example.json";
  exampleDataLabel = "people";
);

func main() {
  // Create a new parser for JSON -> []Float64 and []Float64 -> JSON.
  parser := parse.NewParser();

  // Read in the JSON file as bytes.
  bytes, _ := ioutil.ReadFile(exampleDataFileName);

  // Parse the bytes to JSON.
  json := parser.BytesToJSON(bytes);

  // Pass in the label and JSON.
  // Returns an array of []Float64's that correspond to normalized weights.
  data := parser.JSONToFloat64Matrix(exampleDataLabel, json);

  // Create an RPHash Cluster.
  // Pass in the data, and the number of clusters you want to create from the data.
  cluster := api.NewRPHash(data, numberOfClusters);

  // In reality, one will use our MapReduce enviornment to distribute
  // and harness the full power of the clustering framework.
  // For simplicities sake, perform Map
  cluster.Map();

  // Then Reduce
  cluster.Reduce();

  // Get the centroids from the data.
  centroids := cluster.GetCentroids();

  // Parse the centroids back to JSON and read the information obtained.
  jsonCentroidInformation := parser.Float64MatrixToJSON(exampleDataLabel, centroids);
  fmt.Println(jsonCentroidInformation);
};
