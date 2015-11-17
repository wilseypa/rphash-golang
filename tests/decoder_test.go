package tests;

import (
  "testing"
  "math/rand"
  "github.com/wenkesj/rphash/decoder"
  "github.com/wenkesj/rphash/hash"
  "github.com/wenkesj/rphash/utils"
);

func TestSpherical(t *testing.T) {
  var d, k, l, iterations int = 64, 6, 4, 100;
  sphere := decoder.NewSpherical(d, k, l);
  for i := 0; i < 100; i++ {
    var ct int = 0;
    var distavg float64 = 0.0;
    for j := 0; j < iterations; j++ {
      p1, p2 := make([]float64, d), make([]float64, d);
      for j := 0; j < d; j++ {
        p1[j], p2[j] = rand.NormFloat64(), rand.NormFloat64();
      }
      distavg += utils.Distance(p1, p2);
      mh := hash.NewMurmur(1 << 63 - 1);
      hp1, hp2 := sphere.Hash(utils.Normalize(p1)), sphere.Hash(utils.Normalize(p2));
      hash1, hash2 := mh.Hash(hp1), mh.Hash(hp2);
      if hash1 == hash2 {
        ct++;
      }
    }
    if avg := distavg/float64(iterations); avg > 12 || avg < 11 {
      t.Error("X Average not met...");
    }
  }
  t.Log("√ Spherical Decoder test complete");
};

func BenchmarkSpherical(b *testing.B) {
  for i := 0; i < b.N; i++ {
    var d, k, l int = 64, 6, 4;
    sphere := decoder.NewSpherical(d, k, l);
    mh := hash.NewMurmur(1 << 63 - 1);
    p1, p2 := make([]float64, d), make([]float64, d);
    for j := 0; j < d; k++ {
      p1[j], p2[j] = rand.NormFloat64(), rand.NormFloat64();
    }
    b.StopTimer();
    hp1, hp2 := sphere.Hash(utils.Normalize(p1)), sphere.Hash(utils.Normalize(p2));
    mh.Hash(hp1); mh.Hash(hp2);
    b.StartTimer();
  }
};
