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
  counts := khh.GetCounts();

  for i, value := range khh.GetTop() {
    //Due to initilization the time each number is entered is one less than the number itself
    if counts[i] != (value - 1) {
      t.Errorf("The count for element %v was not correct. Expected %v, Actual %v.", i, (value - 1), counts[i]);
    }
  }

    var previousValue = khh.GetTop()[0] - 1; // initilize to expected value so first loop doesn't fail
    //since each value is entered onece more than the previous value the minqueue
    //should be the last k values entered, in acesnding order, ending with the max value, 99.
    for i, value := range khh.GetTop() {
      if value != (previousValue + 1) {
        t.Errorf("The element at %v was not the expected value. Expected %v, Actual %v.", i, (previousValue + 1), value);
      }
      previousValue = value;
  }
  if khh.GetTop()[len(khh.GetTop()) - 1] != 99 {
    t.Errorf("The top count was not the correct the correct value. Expected  %v, Actual %v", 99, khh.GetTop()[len(khh.GetTop()) - 1]);
  }
};

func BenchmarkCountMinSketchAdd(b *testing.B) {
  k := 100;
  khh := itemset.NewKHHCountMinSketch(k);
  for i := 0; i < b.N; i++ {
    khh.Add(rand.Int63());
  }
};
