package tests;

import (
    "testing"
    "github.com/wenkesj/rphash/clusterer"
);

func TestClusterer(t *testing.T) {
  var numberClusters = 2;
  var numDataPoints = 8;
  var dimensionality = 4;
  data := make([][]float64, numDataPoints);
  for i := 0; i < numDataPoints; i++ {
    data[i] = make([]float64, dimensionality)
    for j := 0; j < dimensionality; j++ {
      data[i][j] = float64(i);
    }
  }
  t.Log(data);
  clusterer := clusterer.NewKMeansSimple(numberClusters, data);
  clusterer.Run();
  var result = clusterer.GetCentroids();
  t.Log(result);
  t.Log("√ clusterer test complete");
};
