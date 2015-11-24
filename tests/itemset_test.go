package tests;

import (
  "time"
  "testing"
  "math/rand"
  "github.com/wenkesj/rphash/itemset"
);

func TestCountMinSketch(t *testing.T) {
  k := 100;
  khh := itemset.NewKHHCountMinSketch(k);
  ts := time.Now().UnixNano() / int64(time.Millisecond);
  var i int64;
  for i = 1; i < 50000; i++ {
    khh.Add(rand.Int63n(i)/100);
  }
  t.Log(time.Now().UnixNano() / int64(time.Millisecond) - ts);
  t.Log(khh.GetTop());
  t.Log(khh.GetCounts());
  t.Log("√ Itemset test complete");
};

func BenchmarkCountMinSketch(b *testing.B) {
  k := 64;
  itemset.NewKHHCountMinSketch(k);
  for i := 0; i < b.N; i++ {}
};
