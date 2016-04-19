package types

import (
  "sync"
)

type Iterator interface {
  GetS() [][]float64
  StoreLSHValues([]int64)
  Append(vector []float64)
  PeakLSH() int64
  Next() (value []float64)
  HasNext() (ok bool)
  Reset()
  Size() (count int)
}

type PQueue interface {
  IsEmpty() bool
  Poll(i interface{})
  Push(i interface{})
  Pop() interface{}
  Length() int
  Less(i, j int) bool
  Swap(i, j int)
  String() string
}

type Decoder interface {
  SetVariance(parameterObject float64)
  GetDimensionality() int
  Decode(f []float64) []int64
  GetErrorRadius() float64
  GetDistance() float64
  GetVariance() float64
}

type Projector interface {
  Project(v []float64) []float64
}

type HashSet interface {
  Add(i int64) bool
  Get(i int64) bool
  AddAll(i HashSet)
  GetS() map[int64]bool
  Remove(i int64)
  Length() int
  Contains(i int64) bool
}

type Hash interface {
  Hash(k []int64) int64
}

type Centroid interface {
  Centroid() []float64
  UpdateVector(rp []float64)
  GetCount() int64
  GetID() int64
  GetIDs() HashSet
  AddID(h int64)
}

type CountItemSet interface {
  Add(c int64)
  GetCounts() []int64
  GetTop() []int64
  GetCount() int64
}

type CentroidItemSet interface {
  Add(c Centroid)
  GetCounts() []int64
  GetTop() []Centroid
  GetCount() int64
}

type LSH interface {
  LSHHashSimple(r []float64) int64
  LSHHashStream(r []float64, a int) []int64
  UpdateDecoderVariance(vari float64)
}

type StatTest interface {
  UpdateVarianceSample(vec []float64) float64
}

type RPHashObject interface {
  GetK() int
  NumDataPoints() int
  GetDimensions() int
  GetRandomSeed() int64
  GetNumberOfBlurs() int
  AppendVector(vector []float64)
  GetVectorIterator() Iterator
  GetCentroids() [][]float64
  GetPreviousTopID() []int64
  SetPreviousTopID(i []int64)
  AddCentroid(v []float64)
  SetCentroids(l [][]float64)
  GetNumberOfProjections() int
  SetNumberOfProjections(probes int)
  SetRandomSeed(parseLong int64)
  GetHashModulus() int64
  SetHashModulus(parseLong int64)
  SetDecoderType(dec Decoder)
  GetDecoderType() Decoder
  SetVariance(data [][]float64)
}

type Clusterer interface {
  GetCentroids() [][]float64
}

type StreamClusterer interface {
  AddVectorOnlineStep(x []float64, wg *sync.WaitGroup) Centroid
  GetCentroidsOfflineStep() [][]float64
}
