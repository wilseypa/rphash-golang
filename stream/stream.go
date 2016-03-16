package stream

import (
  "github.com/wilseypa/rphash-golang/defaults"
  "github.com/wilseypa/rphash-golang/types"
  "github.com/wilseypa/rphash-golang/utils"
  "math/rand"
  "sync"
)

type Stream struct {
  RunCount            int
  counts              []int64
  centroids           [][]float64
  variance            float64
  CentroidCounter     types.CentroidItemSet
  randomSeedGenerator *rand.Rand
  rphashObject        types.RPHashObject
  lshGroup            []types.LSH
  decoder             types.Decoder
  projector           types.Projector
  hash                types.Hash
  varianceTracker     types.StatTest
}

func NewStream(rphashObject types.RPHashObject) *Stream {
  randomSeedGenerator := rand.New(rand.NewSource(rphashObject.GetRandomSeed()))
  hash := defaults.NewHash(rphashObject.GetHashModulus())
  decoder := rphashObject.GetDecoderType()
  varianceTracker := defaults.NewStatTest(0.01)
  projections := rphashObject.GetNumberOfProjections()
  k := rphashObject.GetK() * projections
  CentroidCounter := defaults.NewCentroidCounter(k)
  lshGroup := make([]types.LSH, projections)
  var projector types.Projector
  for i := 0; i < projections; i++ {
    projector = defaults.NewProjector(rphashObject.GetDimensions(), decoder.GetDimensionality(), randomSeedGenerator.Int63())
    lshGroup[i] = defaults.NewLSH(hash, decoder, projector)
  }
  return &Stream{
    counts:              nil,
    centroids:           nil,
    variance:            0,
    RunCount:            0,
    CentroidCounter:     CentroidCounter,
    randomSeedGenerator: randomSeedGenerator,
    rphashObject:        rphashObject,
    lshGroup:            lshGroup,
    hash:                hash,
    decoder:             decoder,
    projector:           projector,
    varianceTracker:     varianceTracker,
  }
}

func (this *Stream) AddVectorOnlineStep(vec []float64, wg *sync.WaitGroup) types.Centroid {
  if wg != nil {
    defer wg.Done()
  }
  c := defaults.NewCentroidStream(vec)
  tmpvar := this.varianceTracker.UpdateVarianceSample(vec)

  if this.variance != tmpvar {
    for _, lsh := range this.lshGroup {
      lsh.UpdateDecoderVariance(tmpvar)
    }
    this.variance = tmpvar
  }

  for _, lsh := range this.lshGroup {
    hash := lsh.LSHHashStream(vec, this.rphashObject.GetNumberOfBlurs())

    for _, h := range hash {
      c.AddID(h)
    }
  }
  return c
}

func (this *Stream) GetCentroids() [][]float64 {
  if this.centroids == nil {
    if this.RunCount == 0 {
      this.Run()
    }
    var centroids [][]float64
    for _, cent := range this.CentroidCounter.GetTop() {
      centroids = append(centroids, cent.Centroid())
    }
    this.centroids = defaults.NewKMeansStream(this.rphashObject.GetK(), centroids, this.CentroidCounter.GetCounts()).GetCentroids()
  }
  return this.centroids
}

func (this *Stream) GetVectors() [][]float64 {
  return this.rphashObject.GetVectorIterator().GetS()
}

func (this *Stream) AppendVector(vector []float64) {
  this.rphashObject.AppendVector(vector)
}

func (this *Stream) GetCentroidsOfflineStep() [][]float64 {
  var centroids [][]float64
  var counts []int64
  for i := 0; i < len(this.CentroidCounter.GetTop()); i++ {
    centroids = append(centroids, this.CentroidCounter.GetTop()[i].Centroid())
    counts = append(counts, this.CentroidCounter.GetCounts()[i])
  }
  this.centroids = defaults.NewKMeansStream(this.rphashObject.GetK(), centroids, counts).GetCentroids()
  count := int((utils.Max(counts) + utils.Min(counts)) / 2)
  counts = []int64{}
  for i := 0; i < this.rphashObject.GetK(); i++ {
    counts = append(counts, int64(count))
  }
  this.counts = counts
  return this.centroids
}

func Monitor(centChannel chan types.Centroid, wg *sync.WaitGroup) {
  wg.Wait()
  close(centChannel)
}

func (this *Stream) Run() {
  this.RunCount++
  count := 0
  vecs := this.rphashObject.GetVectorIterator()
  vecCount := len(vecs.GetS())
  centChannel := make(chan types.Centroid, vecCount)
  wg := new(sync.WaitGroup)

  vec := vecs.Next()
  for vecs.HasNext() {
    wg.Add(1)
    go func(vec []float64, centChannel chan types.Centroid, wg *sync.WaitGroup) {
      centChannel <- this.AddVectorOnlineStep(vec, wg)
    }(vec, centChannel, wg)
    count++
    vec = vecs.Next()
  }

  go Monitor(centChannel, wg)

  for i := 0; i < vecCount; i++ {
    c := <-centChannel
    if c == nil {
      return
    }
    this.CentroidCounter.Add(c)
  }
}
