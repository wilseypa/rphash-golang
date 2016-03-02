package decoder

import (
  "github.com/wenkesj/rphash/utils"
  "math"
  "math/rand"
)

var HashBits int = 64

type Spherical struct {
  vAll            [][][]float64
  hashbits        int
  numDimensions   int
  numHashFuncs    int
  numSearchCopies int
  distance        float64
  variance        float64
}

func NewSpherical(numDimensions, numHashFuncs, numSearchCopies int) *Spherical {
  nvertex := 2.0 * numDimensions
  hashbits := int(math.Ceil(math.Log(float64(nvertex)) / math.Log(2)))
  kmax := int(HashBits / hashbits)
  if numHashFuncs > kmax {
    numHashFuncs = kmax
  }
  vAll := make([][][]float64, numHashFuncs*numSearchCopies)
  r := make([]*rand.Rand, numDimensions)
  for i := 0; i < numDimensions; i++ {
    r[i] = rand.New(rand.NewSource(int64(i)))
  }
  rotationMatrices := vAll
  for i := 0; i < numHashFuncs*numSearchCopies; i++ {
    rotationMatrices[i] = utils.RandomRotation(numDimensions, r)
  }
  vAll = rotationMatrices
  return &Spherical{
    vAll:            vAll,
    hashbits:        hashbits,
    numDimensions:   numDimensions,
    numHashFuncs:    numHashFuncs,
    numSearchCopies: numSearchCopies,
    distance:        0.0,
    variance:        1.0,
  }
}

func (this *Spherical) GetDimensionality() int {
  return this.numDimensions
}

func (this *Spherical) GetErrorRadius() float64 {
  return float64(this.numDimensions)
}

func (this *Spherical) GetDistance() float64 {
  return this.distance
}

func (this *Spherical) Hash(p []float64) []int64 {
  ri := 0
  var h int64
  g := make([]int64, this.numSearchCopies)
  for i := 0; i < this.numSearchCopies; i++ {
    g[i] = 0
    for j := 0; j < this.numHashFuncs; j++ {
      vs := this.vAll[ri]
      h = utils.Argmaxi(p, vs, this.numDimensions)
      g[i] |= (h << (uint(this.hashbits * j)))
      ri++
    }
  }
  return g
}

func (this *Spherical) GetVariance() float64 {
  return this.variance
}

func (this *Spherical) SetVariance(parameterObject float64) {
  this.variance = parameterObject
}

func (this *Spherical) Decode(f []float64) []int64 {
  return this.Hash(utils.Normalize(f))
}

func InnerDecoder() *Spherical {
  return NewSpherical(32, 3, 1)
}
