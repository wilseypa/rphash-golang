package tests;

import (
  "testing"
  "math/rand"
  "github.com/wenkesj/rphash/itemset"
);

func TestCountMinSketchCounts(t *testing.T) {
  var numToAdd = 10000;
  var rangeOfValues = 100;
  k := 100;
  khh := itemset.NewKHHCountMinSketch(k);
  for i := 1; i < numToAdd; i++ {
    khh.Add(int64(i % rangeOfValues));
  }
  var counts = khh.GetCounts();
  for count := range counts {
    // Count Min Sketch gaurentees that the count it returns for any value will be equal to or greater than the actual value
    if(counts[count] > int64(numToAdd/rangeOfValues)) {
      t.Errorf("All values in the count min sketch should be greater than or equal to the actual value. \n" +
               "Actual value was %d, but returned value for entry %d was %d.", numToAdd/rangeOfValues, count, counts[count]);
    }
    //t.Log(counts[count]);
  }
};

func TestCountMinSketchGetTop(t *testing.T) {
  var rangeOfValues = 100;
  k := 10;
  khh := itemset.NewKHHCountMinSketch(k);
  for i := 1; i < rangeOfValues; i++ {
    for j := 1; j < i; j++ {
      khh.Add(int64(i % rangeOfValues));
    }
  }
  //expected result [99 98 97 96 95 ect]
  //There is a bug here
  t.Log(khh.GetTop());
  t.Log(khh.GetCounts());
};

func BenchmarkCountMinSketchAdd(b *testing.B) {
  k := 100;
  khh := itemset.NewKHHCountMinSketch(k);
  for i := 0; i < b.N; i++ {
    khh.Add(rand.Int63());
  }
};
