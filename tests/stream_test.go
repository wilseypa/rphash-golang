package tests

import (
  "github.com/wilseypa/rphash-golang/clusterer"
  "github.com/wilseypa/rphash-golang/reader"
  "github.com/wilseypa/rphash-golang/stream"
  "github.com/wilseypa/rphash-golang/utils"
  "math/rand"
  "testing"
  "time"
)

func TestStreamingKMeansOnNumImagesData(t *testing.T) {
  numClusters := 10
  lines, err := utils.ReadLines("../demo/data/MNISTnumImages5000.txt");
  if err != nil {
    panic(err);
  }
  dimensionality := len(lines[0]);
  data := utils.StringArrayToFloatArray(lines);
  
  start := time.Now();
  kmeansStream := clusterer.NewKMeansStream(numClusters, 10, dimensionality)
  for _, vector := range data {
    kmeansStream.AddDataPoint(vector);
  }
  result := kmeansStream.GetCentroids()
  time := time.Since(start);
  totalSqDist := float64(0)
  for _, vector := range data {
    _, dist := utils.FindNearestDistance(vector, result)
    totalSqDist += dist * dist
  }
  
  t.Log("Total Square Distance: ", totalSqDist);
  t.Log("Average Square Distance: ", totalSqDist/float64(len(data)));
  t.Log("Runtime(seconds): ", time.Seconds());
  
  if len(result) != numClusters {
    t.Errorf("RPHash Stream did not present the correct number of clusters.")
  }
}

func TestStreamingRPHashOnNumImagesData(t *testing.T) {
  numClusters := 10
  lines, err := utils.ReadLines("../demo/data/MNISTnumImages5000.txt");
  if err != nil {
    panic(err);
  }
  dimensionality := len(lines[0]);
  data := utils.StringArrayToFloatArray(lines);

  start := time.Now();
  rphashObject := reader.NewStreamObject(dimensionality, numClusters)
  rphashStream := stream.NewStream(rphashObject)
  for _, vector := range data {
    rphashStream.AppendVector(vector)
  }
  rpHashResult := rphashStream.GetCentroids()
  time := time.Since(start);
  totalSqDist := float64(0)
  for _, vector := range data {
    _, dist := utils.FindNearestDistance(vector, rpHashResult)
    totalSqDist += dist * dist
  }
  
  t.Log("Total Square Distance: ", totalSqDist);
  t.Log("Average Square Distance: ", totalSqDist/float64(len(data)));
  t.Log("Runtime(seconds): ", time.Seconds());
  
  if len(rpHashResult) != numClusters {
    t.Errorf("RPHash Stream did not present the correct number of clusters.")
  }
}
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
