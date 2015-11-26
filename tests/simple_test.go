package tests;

import (
  "testing"
  "github.com/wenkesj/rphash/reader"
  "github.com/wenkesj/rphash/simple"
);

func TestSimple(t *testing.T) {
  k := 100;
  data := make([][]float64, k);
  RPHashObject := reader.NewSimpleArray(data, k);
  t.Log(RPHashObject);
  simpleObject := simple.NewSimple(RPHashObject);
  simpleObject.Run();
  t.Log("√ Simple test complete");
};

/*func BenchmarkSimple(b *testing.B) {
  for i := 0; i < b.N; i++ {}
};*/
