package tests

import (
  "fmt"
  "github.com/wilseypa/rphash-golang/clusterer"
  "github.com/wilseypa/rphash-golang/reader"
  "github.com/wilseypa/rphash-golang/simple"
  "github.com/wilseypa/rphash-golang/utils"
  "math/rand"
  "testing"
  "time"
)

func TestSimpleLeastDistanceVsKmeans(t *testing.T) {

  //Create fake data
  var numClusters = 5
  var numRows = 400
  var dimensionality = 1000
  data := make([][]float64, numRows, numRows)
  for i := 0; i < numRows; i++ {
    row := make([]float64, dimensionality, dimensionality)
    for j := 0; j < dimensionality; j++ {
      row[j] = rand.Float64()
    }
    data[i] = row
  }

  start := time.Now()
  //Test RPHash with Fake Object
  RPHashObject := reader.NewSimpleArray(data, numClusters)
  simpleObject := simple.NewSimple(RPHashObject)
  simpleObject.Run()

  if len(RPHashObject.GetCentroids()) != numClusters {
    t.Errorf("Requested %v centriods. But RPHashSimple returned %v.", numClusters, len(RPHashObject.GetCentroids()))
  }
  rpHashResult := RPHashObject.GetCentroids()
  fmt.Println("RPHash: ", time.Since(start))
  //Find clusters using KMeans
  start = time.Now()
  clusterer := clusterer.NewKMeansSimple(numClusters, data)
  clusterer.Run()

  kMeansResult := clusterer.GetCentroids()
  fmt.Println("kMeans: ", time.Since(start))

  var kMeansAssignment = 0
  var rpHashAssignment = 0
  var matchingAssignmentCount = 0
  var kMeansTotalDist = float64(0)
  var rpHashTotalDist = float64(0)
  for _, vector := range data {
    rpHashAssignment, _ = utils.FindNearestDistance(vector, rpHashResult)
    kMeansAssignment, _ = utils.FindNearestDistance(vector, kMeansResult)
    kMeansTotalDist += utils.Distance(vector, kMeansResult[kMeansAssignment])
    rpHashTotalDist += utils.Distance(vector, rpHashResult[rpHashAssignment])
    //t.Log(rpHashAssignments[i], kMeansAssignments[i]);
    if rpHashAssignment == kMeansAssignment {
      matchingAssignmentCount += 1
    }
  }
  t.Log("RPHash:", rpHashTotalDist)
  t.Log("KMeans:", kMeansTotalDist)
  t.Log("Ratio: ", kMeansTotalDist/rpHashTotalDist)
}

func BenchmarkKMeans(b *testing.B) {
  var numClusters = 5
  var numRows = 4000
  var dimensionality = 1000
  data := make([][]float64, numRows, numRows)
  for i := 0; i < numRows; i++ {
    row := make([]float64, dimensionality, dimensionality)
    for j := 0; j < dimensionality; j++ {
      row[j] = rand.Float64()
    }
    data[i] = row
  }
  for i := 0; i < b.N; i++ {
    clusterer := clusterer.NewKMeansSimple(numClusters, data)
    clusterer.Run()

    clusterer.GetCentroids()
  }
}

func BenchmarkSimple(b *testing.B) {
  var numClusters = 5
  var numRows = 4000
  var dimensionality = 1000
  data := make([][]float64, numRows, numRows)
  for i := 0; i < numRows; i++ {
    row := make([]float64, dimensionality, dimensionality)
    for j := 0; j < dimensionality; j++ {
      row[j] = rand.Float64()
    }
    data[i] = row
  }
  for i := 0; i < b.N; i++ {
    RPHashObject := reader.NewSimpleArray(data, numClusters)
    simpleObject := simple.NewSimple(RPHashObject)
    simpleObject.Run()
    RPHashObject.GetCentroids()
  }
}
