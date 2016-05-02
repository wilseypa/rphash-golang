package tests

import (
  "github.com/wilseypa/rphash-golang/clusterer"
  "github.com/wilseypa/rphash-golang/reader"
  "github.com/wilseypa/rphash-golang/stream"
  "github.com/wilseypa/rphash-golang/utils"
  "testing"
  "time"
)

func TestStreamingKMeansOnRandomData(t *testing.T) {
  numClusters := 10
  lines, err := utils.ReadLines("data/fake_data.txt");
  if err != nil {
    panic(err);
  }
  
  data := utils.StringArrayToFloatArray(lines);
  dimensionality := len(data[0]);
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

func TestStreamingRPHashOnRandomData(t *testing.T) {
  numClusters := 10
  lines, err := utils.ReadLines("data/fake_data.txt");
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
