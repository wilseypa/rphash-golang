package reader

import (
  "github.com/wilseypa/rphash-golang/decoder"
  "github.com/wilseypa/rphash-golang/types"
  "github.com/wilseypa/rphash-golang/utils"
  "math"
  "math/rand"
)

type SimpleArray struct {
  data                types.Iterator
  numDataPoints       int
  dimension           int
  numberOfProjections int
  randomSeed          int64
  hashModulus         int64
  k                   int
  numberOfBlurs       int
  decoder             types.Decoder
  centroids           [][]float64
  topIDs              []int64
}

func NewSimpleArray(inData [][]float64, k int) *SimpleArray {
  randomSeed := rand.Int63()
  data := utils.NewIterator(inData)
  numDataPoints := len(inData)
  dimension := 2
  // As the number of rotations increases, the distance increases.
  // Increases the noise.
  numberOfRotations := 6
  numberOfSearches := 1
  numberOfProjections := 2
  numberOfBlurs := 2
  if data != nil {
    // Get the first vector in the data set's length.
    dimension = len(data.GetS()[0])
  }
  hashModulus := int64(math.MaxInt64)
  // Set the target dimension much lower.
  targetDimension := int(math.Floor(float64(dimension / 2)))
  decoder := decoder.NewSpherical(targetDimension, numberOfRotations, numberOfSearches)
  centroids := [][]float64{}
  topIDs := []int64{}
  return &SimpleArray{
    numDataPoints:       numDataPoints,
    data:                data,
    dimension:           dimension,
    numberOfProjections: numberOfProjections,
    randomSeed:          randomSeed,
    hashModulus:         hashModulus,
    k:                   k,
    numberOfBlurs:       numberOfBlurs,
    decoder:             decoder,
    centroids:           centroids,
    topIDs:              topIDs,
  }
}

func (this *SimpleArray) GetVectorIterator() types.Iterator {
  return this.data
}

func (this *SimpleArray) NumDataPoints() int {
  return this.numDataPoints
}

func (this *SimpleArray) GetK() int {
  return this.k
}

func (this *SimpleArray) GetDimensions() int {
  if this.dimension == 0 {
    this.dimension = len(this.data.GetS()[0])
  }
  return this.dimension
}

func (this *SimpleArray) GetHashModulus() int64 {
  return this.hashModulus
}

func (this *SimpleArray) GetRandomSeed() int64 {
  return this.randomSeed
}

func (this *SimpleArray) AddCentroid(v []float64) {
  this.centroids = append(this.centroids, v)
}

func (this *SimpleArray) SetCentroids(l [][]float64) {
  this.centroids = l
}

func (this *SimpleArray) GetCentroids() [][]float64 {
  return this.centroids
}

func (this *SimpleArray) GetNumberOfBlurs() int {
  return this.numberOfBlurs
}

func (this *SimpleArray) GetPreviousTopID() []int64 {
  return this.topIDs
}

func (this *SimpleArray) AppendVector(vector []float64) {
  this.data.Append(vector)
}

func (this *SimpleArray) SetPreviousTopID(top []int64) {
  this.topIDs = top
}

func (this *SimpleArray) SetRandomSeed(parseLong int64) {
  this.randomSeed = parseLong
}

func (this *SimpleArray) SetNumberOfProjections(probes int) {
  this.numberOfProjections = probes
}

func (this *SimpleArray) GetNumberOfProjections() int {
  return this.numberOfProjections
}

func (this *SimpleArray) SetHashModulus(parseLong int64) {
  this.hashModulus = parseLong
}

func (this *SimpleArray) SetDecoderType(decoder types.Decoder) {
  this.decoder = decoder
}

func (this *SimpleArray) GetDecoderType() types.Decoder {
  return this.decoder
}

func (this *SimpleArray) SetVariance(data [][]float64) {
  this.decoder.SetVariance(utils.VarianceSample(data, 0.01))
}

func (this *SimpleArray) GetVariance() float64 {
  return this.decoder.GetVariance()
}
