package tests;

import (
  "testing"
  "github.com/wenkesj/rphash/reader"
  "github.com/wenkesj/rphash/simple"
  "math/rand"
  "github.com/wenkesj/rphash/clusterer"
);

func TestSimple(t *testing.T) {
  //Create fake data
  var numClusters = 4;
  var numRows = 100;
  var dimensionality = 10;
  data := make([][]float64, numRows, numRows);
  for i := 0; i < numRows; i++ {
    row := make([]float64, dimensionality, dimensionality);
    for j := 0; j < dimensionality; j++ {
      row[j] = rand.Float64();
    }
    data[i] = row;
  }

  //Test RPHash with Fake Object
  RPHashObject := reader.NewSimpleArray(data, numClusters);
  simpleObject := simple.NewSimple(RPHashObject);
  simpleObject.Run();

  if len(RPHashObject.GetCentroids()) != numClusters {
    t.Errorf("Requested %v centriods. But RPHashSimple returned %v.", numClusters, len(RPHashObject.GetCentroids()));
  }
  t.Log(RPHashObject.GetCentroids());

  //Find clusters using KMeans
  clusterer := clusterer.NewKMeansSimple(numClusters, data);
  clusterer.Run();
  var kMeansResult = clusterer.GetCentroids();
  t.Log(kMeansResult);
};

/*func BenchmarkSimple(b *testing.B) {
  for i := 0; i < b.N; i++ {}
};*/
