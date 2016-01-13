package tests;

import (
  "testing"
  "github.com/wenkesj/rphash/reader"
  "github.com/wenkesj/rphash/simple"
);

func TestSimple(t *testing.T) {
  k := 100;
  d := 5
  data := make([][]float64, k, k);
  row := make([]float64, d, d);
  for j := 0; j < d; j++ {
    row[j] = float64(0.0);
  }
  for i := 0; i < k; i++ {
      data[i] = row;
  }
  RPHashObject := reader.NewSimpleArray(data, k);
  simpleObject := simple.NewSimple(RPHashObject);
  simpleObject.Run();
  t.Log(simpleObject.GetCentroids());
  t.Log("√ Simple test complete");
};

/*func BenchmarkSimple(b *testing.B) {
  for i := 0; i < b.N; i++ {}
};*/
