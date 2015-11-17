package tests;

import (
    "testing"
    "github.com/wenkesj/rphash/itemset"
);

func TestCountMinSketch(t *testing.T) {
  k := 64;
  itemset.NewKHHCountMinSketch(k);
  t.Log("√ Itemset test complete");
};

func Benchmark(b *testing.B) {
  k := 64;
  itemset.NewKHHCountMinSketch(k);
  for i := 0; i < b.N; i++ {}
};
