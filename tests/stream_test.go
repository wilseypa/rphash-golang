package tests

import (
  "github.com/wilseypa/rphash-golang/clusterer"
  "github.com/wilseypa/rphash-golang/reader"
  "github.com/wilseypa/rphash-golang/stream"
  "github.com/wilseypa/rphash-golang/utils"
  "testing"
  "time"
)
var filePath = "data/fake_data_500_1000.txt"
var numClusters = 10
var dimensionality = 500
var numDataPoints = float64(1000)
func TestStreamingKMeansOnRandomData(t *testing.T) {
  filereader := utils.NewDataFileReader(filePath);
  
  start := time.Now();
  kmeansStream := clusterer.NewKMeansStream(numClusters, 10, dimensionality)
  elapsedtime := time.Since(start);
  for {
    vector := filereader.Next()
    if vector == nil {
      break
    }
    start := time.Now();
    kmeansStream.AddDataPoint(vector);
    elapsedtime = elapsedtime + time.Since(start);
  }
  start = time.Now();
  result := kmeansStream.GetCentroids()
  elapsedtime = elapsedtime + time.Since(start);
  totalSqDist := float64(0)
  filereader = utils.NewDataFileReader(filePath);
  for {
    vector := filereader.Next()
    if vector == nil {
      break
    }
    _, dist := utils.FindNearestDistance(vector, result)
    totalSqDist += dist * dist
  }

  t.Log("Total Square Distance: ", totalSqDist);
  t.Log("Average Square Distance: ", totalSqDist/numDataPoints);
  t.Log("Runtime(seconds): ", elapsedtime.Seconds());

  if len(result) != numClusters {
    t.Errorf("RPHash Stream did not present the correct number of clusters.")
  }
}

func TestStreamingRPHashOnRandomData(t *testing.T) {
  filereader := utils.NewDataFileReader(filePath);
  
  start := time.Now();
  rphashObject := reader.NewStreamObject(dimensionality, numClusters)
  rphashStream := stream.NewStream(rphashObject)
  elapsedtime := time.Since(start);
  for {
    vector := filereader.Next()
    if vector == nil {
      break
    }
    start := time.Now();
    rphashStream.AppendVector(vector);
    elapsedtime = elapsedtime + time.Since(start);
  }
  start = time.Now();
  result := rphashStream.GetCentroids()
  elapsedtime = elapsedtime + time.Since(start);
  totalSqDist := float64(0)
  filereader = utils.NewDataFileReader(filePath);
  for {
    vector := filereader.Next()
    if vector == nil {
      break
    }
    _, dist := utils.FindNearestDistance(vector, result)
    totalSqDist += dist * dist
  }

  t.Log("Total Square Distance: ", totalSqDist);
  t.Log("Average Square Distance: ", totalSqDist/numDataPoints);
  t.Log("Runtime(seconds): ", elapsedtime.Seconds());

  if len(result) != numClusters {
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
