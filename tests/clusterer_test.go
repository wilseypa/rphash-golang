package tests;

import (
    "testing"
    "github.com/wilseypa/rphash-golang/clusterer"
);

func TestClustererUniformVectors(t *testing.T) {
  //initilize data
  var numClusters = 2;
  var numDataPoints = 8;
  var dimensionality = 4;
  data := make([][]float64, numDataPoints);
  for i := 0; i < numDataPoints; i++ {
    data[i] = make([]float64, dimensionality)
    for j := 0; j < dimensionality; j++ {
      data[i][j] = float64(i);
    }
  }

  //run test
  clusterer := clusterer.NewKMeansSimple(numClusters, data);
  clusterer.Run();
  var result = clusterer.GetCentroids();

  //Test Results
  if(len(result) != numClusters) {
    t.Errorf("Clusterer created %v clusters. When %v was input for k.", len(result), numClusters)
  }
  if(len(result[0]) != dimensionality) {
    t.Errorf("Cluster dimensionalioty of %v does not match the dimensionality of the input data, %v.", len(result[0]), dimensionality)
  }
  expectedResults := make([]float64, numClusters);
  expectedResults[0] = 1.5; // (0+1+2+3)/4 = 1.5
  expectedResults[1] = 5.5;//  (4+5+6+7)/4 = 5.5
  for i := 0; i < numClusters; i++ {
    for j := 0; j < dimensionality; j++ {
      if result[i][j] != expectedResults[i] {
        t.Errorf("Data did not cluster as expected. Data: %v, Clusters: %v. Failure at %v, %v.", data, result, i, j)
      }
    }
  }
};
