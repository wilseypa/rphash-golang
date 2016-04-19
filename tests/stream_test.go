package tests

import (
  "github.com/wilseypa/rphash-golang/clusterer"
  "github.com/wilseypa/rphash-golang/reader"
  "github.com/wilseypa/rphash-golang/stream"
  "github.com/wilseypa/rphash-golang/utils"
  "math/rand"
  "testing"
)

func TestStreamingRPHash(t *testing.T) {
  // Create fake data
  var numClusters = 5
  var numRows = 500
  var dimensionality = 500
  var streamCount = 5

  rphashObject := reader.NewStreamObject(dimensionality, numClusters)
  rphashStream := stream.NewStream(rphashObject)
  kmeansStream := clusterer.NewKMeansStream(numClusters, 10, dimensionality)

  for n := 0; n < streamCount; n++ {
    for i := 0; i < numRows; i++ {
      row := make([]float64, dimensionality, dimensionality)
      for j := 0; j < dimensionality; j++ {
        row[j] = rand.Float64()
      }
      rphashStream.AppendVector(row)
      kmeansStream.AddDataPoint(row);
    }
    rphashStream.Run()
    if len(rphashStream.GetCentroids()) != numClusters {
      t.Errorf("RPHash Stream did not present the correct number of clusters.")
    }

    rpHashResult := rphashStream.GetCentroids()
    data := rphashStream.GetVectors()
    kMeansResult := kmeansStream.GetCentroids()

    kMeansAssignment := 0
    rpHashAssignment := 0
    matchingAssignmentCount := 0
    kMeansTotalDist := float64(0)
    rpHashTotalDist := float64(0)
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
}
