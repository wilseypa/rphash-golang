package tests;

import (
  "testing"
  "github.com/wilseypa/rphash-golang/reader"
  // "github.com/wilseypa/rphash-golang/simple"
);

func TestSimple(t *testing.T) {
  var k = 4;
  var numRows = 100;
  var dimensionality = 100;
  data := make([][]float64, numRows, numRows);
  for i := 0; i < numRows; i++ {
    row := make([]float64, dimensionality, dimensionality);
    for j := 0; j < dimensionality; j++ {
      row[j] = float64(i);
    }
    data[i] = row;
  }
  RPHashObject := reader.NewSimpleArray(data, k);
  // simpleObject := simple.NewSimple(RPHashObject);
  // simpleObject.Run();
  t.Log(RPHashObject.GetCentroids());
};

/*func BenchmarkSimple(b *testing.B) {
  for i := 0; i < b.N; i++ {}
};*/
