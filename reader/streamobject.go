package reader

import (
  "github.com/wilseypa/rphash-golang/decoder"
  "github.com/wilseypa/rphash-golang/types"
  "github.com/wilseypa/rphash-golang/utils"
  // "math"
)

type StreamObject struct {
  data                types.Iterator
  numDataPoints       int
  numberOfProjections int
  randomSeed          int64
  numberOfBlurs       int
  k                   int
  dimension           int
  hashModulus         int64
  centroids           [][]float64
  topIDs              []int64
  decoder             types.Decoder
}

func NewStreamObject(dimension, k int) *StreamObject {
  var centroids [][]float64
  var topIDs []int64
  numberOfRotations := 2
  numberOfSearches := 2
  targetDimension := 10
  decoder := decoder.NewSpherical(targetDimension, numberOfRotations, numberOfSearches)
  data := utils.NewIterator([][]float64{})
  return &StreamObject{
    decoder:             decoder,
    data:                data,
    dimension:           dimension,
    randomSeed:          int64(0),
    hashModulus:         int(0 >> 1),
    numberOfProjections: 1,
    numberOfBlurs:       1,
    k:                   k,
    topIDs:              topIDs,
    centroids:           centroids,
    numDataPoints:       0,
  }
}

func (this *StreamObject) GetK() int {
  return this.k
}
func (this *StreamObject) NumDataPoints() int {
  return this.numDataPoints
}

func (this *StreamObject) GetDimensions() int {
  return this.dimension
}

func (this *StreamObject) GetRandomSeed() int64 {
  return this.randomSeed
}

func (this *StreamObject) GetNumberOfBlurs() int {
  return this.numberOfBlurs
}

func (this *StreamObject) GetVectorIterator() types.Iterator {
  return this.data
}

func (this *StreamObject) AppendVector(vector []float64) {
  this.numDataPoints++
  this.data.Append(vector)
}

func (this *StreamObject) GetCentroids() [][]float64 {
  return this.centroids
}

func (this *StreamObject) GetPreviousTopID() []int64 {
  return this.topIDs
}

func (this *StreamObject) SetPreviousTopID(top []int64) {
  this.topIDs = top
}

func (this *StreamObject) AddCentroid(v []float64) {
  this.centroids = append(this.centroids, v)
}

func (this *StreamObject) SetCentroids(l [][]float64) {
  this.centroids = l
}

func (this *StreamObject) GetNumberOfProjections() int {
  return this.numberOfProjections
}

func (this *StreamObject) SetNumberOfProjections(probes int) {
  this.numberOfProjections = probes
}

func (this *StreamObject) SetNumberOfBlurs(parseInt int) {
  this.numberOfBlurs = parseInt
}

func (this *StreamObject) SetRandomSeed(parseLong int64) {
  this.randomSeed = parseLong
}

func (this *StreamObject) GetHashModulus() int64 {
  return this.hashModulus
}

func (this *StreamObject) SetHashModulus(parseLong int64) {
  this.hashModulus = int64(parseLong)
}

func (this *StreamObject) SetDecoderType(dec types.Decoder) {
  this.decoder = dec
}

func (this *StreamObject) GetDecoderType() types.Decoder {
  return this.decoder
}

func (this *StreamObject) SetVariance(data [][]float64) {
  this.decoder.SetVariance(utils.VarianceSample(data, 0.01))
}

func (this *StreamObject) GetVariance() float64 {
  return this.decoder.GetVariance()
}
