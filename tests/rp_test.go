/** Test package for randomprojection */
package rphash;

import (
    "testing"
    "time"
    "fmt"
    "math/rand"
    "github.com/wenkesj/rphash/projection/rp"
);

func TestRandomProjection(t *testing.T) {
  //there is probably a better way to test this than hard coding.
  data := []float64{1.0,0.0,2.0,7.0,4.0,0.0,8.0,3.0,2.0,1.0};
  expectedResult := []float64{1.224744871391589, 13.472193585307478};
  var inDimensions, outDimentions int = 10, 2;
  //Use a uniform seed for testing
  var seed int64 = 0;
  RP := rp.New(inDimensions, outDimentions, seed);
  result := RP.Project(data);
  if(len(result) != len(expectedResult)){
    t.Error("The result and expected result are not the same length.");
  }
  for i := 0; i < len(result); i++ {
    if(result[i] != expectedResult[i]){
      t.Error(fmt.Sprintf("The result at index %d: %f did not match the expected result: %f", i, result[i], expectedResult[i]));
    }
  }
}

func BenchmarkRandomProjection(b *testing.B) {
    var inDimensions, outDimentions int = 10, 2;
    for i := 0; i < b.N; i++ {
        b.StopTimer();
        var randomGen = rand.New(rand.NewSource(int64(time.Now().Nanosecond())));
        data := make([]float64, inDimensions);
        for i := 0; i < inDimensions; i++ {
            data[i] = randomGen.Float64();
        }
        b.StartTimer();
        var seed int64 = int64(time.Now().Nanosecond());
        RP := rp.New(inDimensions, outDimentions, seed);
        RP.Project(data);
    }
}
