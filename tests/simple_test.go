package tests;

import (
  "testing"
  "github.com/wenkesj/rphash/reader"
  "github.com/wenkesj/rphash/simple"
  "math/rand"
  "github.com/wenkesj/rphash/clusterer"
  "github.com/wenkesj/rphash/utils"
  "github.com/wenkesj/rphash/hash"
  "github.com/wenkesj/rphash/decoder"
  "github.com/wenkesj/rphash/projector"
  "github.com/wenkesj/rphash/lsh"
  "sort"
);

func TestSimple(t *testing.T) {
  //LSH function used for testing results only
  //Totally independent of the seed. Will fluctuate from +-%15 matching.
  var seed int64 = 0;
  var dimensions, clusters, searchCopies int = 100, 8, 4;

  var inDimensions, outDimentions int = 100, 2;
  hash := hash.NewMurmur(1 << 63 - 1);
  decoder := decoder.NewSpherical(dimensions, clusters, searchCopies);
  projector := projector.NewDBFriendly(inDimensions, outDimentions, seed);
  lsh := lsh.NewLSH(hash, decoder, projector);

  //Create fake data
  var numClusters = 8;
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

  t.Log(RPHashObject.GetCentroids());
  //Sort centriods by LSH result
  rpHashMap := make(map[int64][]float64);
  for _, result := range RPHashObject.GetCentroids() {
    t.Log("RPHASH: ", lsh.LSHHashSimple(result));
    rpHashMap[lsh.LSHHashSimple(result)] = result;
  }

  rpHashResult := make([][]float64, numClusters, numClusters)
  var keys []int
  for key := range rpHashMap {
    keys = append(keys, int(key))
  }
  sort.Ints(keys)
  for i, key := range keys {
    rpHashResult[i] = rpHashMap[int64(key)];
  }

  //Find clusters using KMeans and sort by LSH result
  clusterer := clusterer.NewKMeansSimple(numClusters, data);
  clusterer.Run();

  //Sort centriods by LSH result
  kMeansMap := make(map[int64][]float64);
  for _, result := range clusterer.GetCentroids() {
    t.Log("KMEANS: ",lsh.LSHHashSimple(result));
    kMeansMap[lsh.LSHHashSimple(result)] = result;
  }

  kMeansResult := make([][]float64, numClusters, numClusters)
  keys = nil;
  for key := range kMeansMap {
    keys = append(keys, int(key))
  }
  sort.Ints(keys)
  for i, key := range keys {
    kMeansResult[i] = kMeansMap[int64(key)];
  }

  kMeansAssignments := make([]int, numRows, numRows);
  rpHashAssignments := make([]int, numRows, numRows);
  var matchingAssignmentCount = 0;
  for i, vector := range data {
    rpHashAssignments[i] = utils.FindNearestDistance(vector, rpHashResult);
    kMeansAssignments[i] = utils.FindNearestDistance(vector, kMeansResult);
    t.Log(rpHashAssignments[i], kMeansAssignments[i]);
    if rpHashAssignments[i] == kMeansAssignments[i] {
      matchingAssignmentCount += 1;
    }
  }
  t.Log("Percent Matching: ", float64(matchingAssignmentCount)/float64(numRows))
};

/*func BenchmarkSimple(b *testing.B) {
  for i := 0; i < b.N; i++ {}
};*/
