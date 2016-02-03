package tests;

import (
  "testing"
  "github.com/wenkesj/rphash/reader"
  "github.com/wenkesj/rphash/simple"
  "math/rand"
  "github.com/wenkesj/rphash/clusterer"
  "github.com/wenkesj/rphash/utils"
);

func TestSimpleLeastDistanceVsKmeans(t *testing.T) {

  //Create fake data
  var numClusters = 16;
  var numRows = 500;
  var dimensionality = 100;
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

  //Sort centriods by LSH result
  rpHashResult := RPHashObject.GetCentroids();
  //Find clusters using KMeans and sort by LSH result
  clusterer := clusterer.NewKMeansSimple(numClusters, data);
  clusterer.Run();

  kMeansResult := clusterer.GetCentroids();

  var kMeansAssignment = 0;
  var rpHashAssignment = 0;
  var matchingAssignmentCount = 0;
  var kMeansTotalDist = float64(0);
  var rpHashTotalDist = float64(0);
  for _, vector := range data {
    rpHashAssignment = utils.FindNearestDistance(vector, rpHashResult);
    kMeansAssignment = utils.FindNearestDistance(vector, kMeansResult);
    kMeansTotalDist += utils.Distance(vector, kMeansResult[kMeansAssignment]);
    rpHashTotalDist += utils.Distance(vector, rpHashResult[rpHashAssignment]);
    //t.Log(rpHashAssignments[i], kMeansAssignments[i]);
    if rpHashAssignment == kMeansAssignment {
      matchingAssignmentCount += 1;
    }
  }
  t.Log("RPHash:", rpHashTotalDist);
  t.Log("KMeans:", kMeansTotalDist);
  t.Log("Ratio: ", kMeansTotalDist/rpHashTotalDist)
};

/*func BenchmarkSimple(b *testing.B) {
  for i := 0; i < b.N; i++ {}
};*/
