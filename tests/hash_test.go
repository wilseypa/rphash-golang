package tests;

import (
    "testing"
    "github.com/wenkesj/rphash/hash"
);

func TestMurmur(t *testing.T) {
  var hashSize = 100;

  testHash := hash.NewMurmur(int64(hashSize));
  resultCount := make([]int64, hashSize);
  for i := 0; i < hashSize * 100; i++ {
    fakeArray :=  make([]int64, 1);
    fakeArray[0] = int64(i);
    resultCount[testHash.Hash(fakeArray)]++;
  }
  var maxCount = hashSize * 13 / 10;
  var minCount = hashSize * 7 / 10;
  for index , indexCount := range resultCount {
    if(indexCount > int64(maxCount) || indexCount < int64(minCount)) {
      t.Errorf("X - Expected  between %d - %d results, for index %d got %d", minCount, maxCount, index, indexCount);
    }
  }
  t.Log("√ Murmur Hash test complete");
};

func BenchmarkLargeArray(b *testing.B) {
  var hashSize = 1000;
  testHash := hash.NewMurmur(int64(hashSize));
  testArray := make([]int64, 1000);
  for i := 0; i < b.N; i++ {
      b.StopTimer();
      for j := 0; j < len(testArray); j++ {
        testArray[j] = int64(i * j);
      }
      b.StartTimer();
      testHash.Hash(testArray);
  }
};
