package tests;

import (
    "testing"
    "github.com/wenkesj/rphash/hash"
    "github.com/wenkesj/rphash/decoder"
    "github.com/wenkesj/rphash/projector"
    "github.com/wenkesj/rphash/lsh"
    "math"
);
// The datapoints are seeded in so that the first two data points are near eachother in euclidian geometery and the 3rd and 4th datapoint are
// near eachother in euclidian geometery. So the result1Cluster1 and result2Cluster1 should be closer together than the other two points.
//The same is true for the points in cluster two vs either point in cluster one.
func TestLSHSimple(t *testing.T) {
  var seed int64 = 0;
  var d, k, l int = 64, 6, 4;
  dataPoint1Cluster1 := []float64{1.0,0.0,2.0,7.0,4.0,0.0,8.0,3.0,2.0,1.0};
  dataPoint2Cluster1 := []float64{2.0,3.0,2.0,6.0,5.5,2.0,8.0,3.1,2.0,0.0};

  dataPoint1Cluster2 := []float64{10.0,-12.0,6.0,8.0,9.0,0.0,-2.0,9.0,15.0,4.0};
  dataPoint2Cluster2 := []float64{9.0,-10.0,2.0,8.0,9.5,0.0,-6.0,7.0,13.0,2.0};

  var inDimensions, outDimentions int = 10, 2;
  hash := hash.NewMurmur(1 << 63 - 1);
  decoder := decoder.NewSpherical(d, k, l);
  projector := projector.NewDBFriendly(inDimensions, outDimentions, seed);
  lsh := lsh.NewLSH(hash, decoder, projector);
  result1Cluster1 :=lsh.LSHHashSimple(dataPoint1Cluster1);
  result2Cluster1 :=lsh.LSHHashSimple(dataPoint2Cluster1);
  result1Cluster2 :=lsh.LSHHashSimple(dataPoint1Cluster2);
  result2Cluster2 :=lsh.LSHHashSimple(dataPoint2Cluster2);
  // Assert that results are still localy sensetive based on the original euclidian geometry
  if(math.Abs(float64(result1Cluster1 - result2Cluster1)) > math.Abs(float64(result1Cluster1 - result1Cluster2))) {
    t.Errorf("The first datapoint in cluster two is closer to the first data point in cluster one than the second data point in cluster one" +
            "datapoint cluster one datapoint one: %d, datapoint cluster one datapoint two: %d, datapoint cluster two datapoint one: %d",
            result1Cluster1, result2Cluster1, result1Cluster2);
  }
  if(math.Abs(float64(result1Cluster1 - result2Cluster1)) > math.Abs(float64(result1Cluster1 - result2Cluster2))) {
    t.Errorf("The second datapoint in cluster two is closer to the first data point in cluster one than the second data point in cluster one" +
      "\nCluster one datapoint one: %d, \nCluster one datapoint two: %d, \nCluster two datapoint two: %d",
      result1Cluster1, result2Cluster1, result2Cluster2);
  }
  if(math.Abs(float64(result1Cluster2 - result2Cluster2)) > math.Abs(float64(result1Cluster1 - result1Cluster2))) {
    t.Errorf("The first datapoint in cluster one is closer to the first data point in cluster two than the second data point in cluster two" +
      "\nCluster one datapoint one: %d, \nCluster two datapoint one: %d, \nCluster two datapoint two: %d",
      result1Cluster1, result1Cluster2, result2Cluster2);
  }

  t.Log("√ LSH Simple test complete");
};

func TestLSHStream(t *testing.T) {
  var seed int64 = 0;
  var d, k, l int = 64, 6, 4;
  data := []float64{1.0,0.0,2.0,7.0,4.0,0.0,8.0,3.0,2.0,1.0};
  var inDimensions, outDimentions int = 10, 2;
  hash := hash.NewMurmur(1 << 63 - 1);
  decoder := decoder.NewSpherical(d, k, l);
  projector := projector.NewDBFriendly(inDimensions, outDimentions, seed);
  lsh := lsh.NewLSH(hash, decoder, projector);
  lsh.LSHHashStream(data, 1);
  t.Log("√ LSH Stream test complete");
};

func BenchmarkSimple(b *testing.B) {
  var seed int64 = 0;
  var d, k, l int = 64, 6, 4;
  data := []float64{1.0,0.0,2.0,7.0,4.0,0.0,8.0,3.0,2.0,1.0};
  var inDimensions, outDimentions int = 10, 2;
  hash := hash.NewMurmur(1 << 63 - 1);
  decoder := decoder.NewSpherical(d, k, l);
  projector := projector.NewDBFriendly(inDimensions, outDimentions, seed);
  for i := 0; i < b.N; i++ {
      lsh := lsh.NewLSH(hash, decoder, projector);
      b.StopTimer();
      lsh.LSHHashSimple(data);
      b.StartTimer();
  }
};

func BenchmarkStream(b *testing.B) {
  var seed int64 = 0;
  var d, k, l int = 64, 6, 4;
  data := []float64{1.0,0.0,2.0,7.0,4.0,0.0,8.0,3.0,2.0,1.0};
  var inDimensions, outDimentions int = 10, 2;
  hash := hash.NewMurmur(1 << 63 - 1);
  decoder := decoder.NewSpherical(d, k, l);
  projector := projector.NewDBFriendly(inDimensions, outDimentions, seed);
  for i := 0; i < b.N; i++ {
      lsh := lsh.NewLSH(hash, decoder, projector);
      b.StopTimer();
      lsh.LSHHashStream(data, 1);
      b.StartTimer();
  }
};
