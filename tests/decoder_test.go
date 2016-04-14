package tests

import (
  "github.com/wilseypa/rphash-golang/decoder"
  "github.com/wilseypa/rphash-golang/hash"
  "github.com/wilseypa/rphash-golang/utils"
  "math/rand"
  "testing"
  "time"
)

func TestSpherical(t *testing.T) {
  var dimension, k, l, iterations int = 64, 6, 4, 10000
  sphere := decoder.NewSpherical(dimension, k, l)
  var collisions int = 0
  var distavg float64 = 0.0
  for j := 0; j < iterations; j++ {
    p1, p2 := make([]float64, dimension), make([]float64, dimension)
    for k := 0; k < dimension; k++ {
      p1[k] = rand.Float64()*2 - 1
      p2[k] = rand.Float64()*2 - 1
    }
    /* Get the distance of each vector from eachother. */
    distavg += utils.Distance(p1, p2)
    mh := hash.NewMurmur(1<<63 - 1)
    /* Decode from 24-dimensions -> 1-dimensional integer */
    hp1, hp2 := sphere.Hash(utils.Normalize(p1)), sphere.Hash(utils.Normalize(p2))
    /* Blurring the integers into a smaller space. */
    hash1, hash2 := mh.Hash(hp1), mh.Hash(hp2)
    if hash1 == hash2 {
      collisions++
    }
  }
  if collisions > (iterations / 100) {
    t.Errorf("More than 1 percent of the iterations resulted in collisions. %v collisions in %v iterations.",
      collisions, iterations)
  }
  t.Log("Average Distance: ", distavg/float64(iterations))
  t.Log("Percent collisions : ", float64(collisions)/float64(iterations))
  t.Log("√ Spherical Decoder test complete")
}

func BenchmarkSpherical(b *testing.B) {
  b.StopTimer()
  randomSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
  var d, k, l int = 64, 6, 4
  sphere := decoder.NewSpherical(d, k, l)
  p1, p2 := make([]float64, d), make([]float64, d)
  for i := 0; i < b.N; i++ {
    for j := 0; j < d; j++ {
      p1[j], p2[j] = randomSeed.NormFloat64(), randomSeed.NormFloat64()
    }
    b.StartTimer()
    hp1, hp2 := sphere.Hash(utils.Normalize(p1)), sphere.Hash(utils.Normalize(p2))
    b.StopTimer()
    if hp1 == nil || hp2 == nil {
      b.Error("Spherical hashes are null")
    }
  }
}
