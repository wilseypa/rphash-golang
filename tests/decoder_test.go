package tests;

import (
  "testing"
  "time"
  "math/rand"
  "github.com/wenkesj/rphash/decoder"
  "github.com/wenkesj/rphash/hash"
  "github.com/wenkesj/rphash/utils"
);

func TestSpherical(t *testing.T) {
  var d, k, l, iterations int = 64, 6, 4, 1000;
  sphere := decoder.NewSpherical(d, k, l);
  for i := 0; i < 100; i++ {
    var ct int = 0;
    var distavg float64 = 0.0;
    for j := 0; j < iterations; j++ {
      p1, p2 := make([]float64, d), make([]float64, d);
      for k := 0; k < d; k++ {
        p1[k] = rand.Float64() * 2 - 1;
        p2[k] = p1[k] + rand.NormFloat64() * float64(i/100);
      }
      distavg += utils.Distance(p1, p2);
      t.Log(distavg);
      mh := hash.NewMurmur(1 << 63 - 1);
      hp1, hp2 := sphere.Hash(utils.Normalize(p1)), sphere.Hash(utils.Normalize(p2));
      hash1, hash2 := mh.Hash(hp1), mh.Hash(hp2);
      if hash1 == hash2 {
        ct++;
      }
    }
    t.Log(distavg / float64(iterations), "\t", ct / iterations);
    //TODO test actual output of spherical decoder. here rather than logging
  }
  t.Log("√ Spherical Decoder test complete");
};

func BenchmarkSpherical(b *testing.B) {
  b.StopTimer();
  randomSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
  var d, k, l int = 64, 6, 4;
  sphere := decoder.NewSpherical(d, k, l);
  p1, p2 := make([]float64, d), make([]float64, d);
  for i := 0; i < b.N; i++ {
    for j := 0; j < d; j++ {
      p1[j], p2[j] = randomSeed.NormFloat64(), randomSeed.NormFloat64();
    }
    b.StartTimer();
    hp1, hp2 := sphere.Hash(utils.Normalize(p1)), sphere.Hash(utils.Normalize(p2));
    b.StopTimer();
    if(hp1 == nil || hp2 == nil) {
      b.Error("Spherical hashes are null");
    }
  }
};
