package clusterer
import (
  "math"
  "github.com/wilseypa/rphash-golang/utils"
  "github.com/wilseypa/rphash-golang/itemset"
  "math/rand"
)
//Implements clusterer type
type KMeansStream struct {
  debug               int
  k                   int
  n                   int
  dataCount           int
  dimensionality      int
  maxCandidateClusters int
  frequency           float64
  frequencyChange     float64
  random              rand.Source
  candidateClusters   []itemset.Centroid
}

func NewKMeansStream(k int, n int, dimensionality int) *KMeansStream {
  frequency := 1.0 / (float64(k) * (1 + math.Log10(float64(n))))
  maxCandidateClusters := int(math.Log10(float64(n)) * float64(k))
  candidateClusters := []itemset.Centroid{}
  return &KMeansStream{
    debug:                0,
    k:                    k,
    n:                    n,
    dataCount:            0,
    dimensionality:       dimensionality,
    maxCandidateClusters: maxCandidateClusters,
    frequency:            frequency,
    frequencyChange:      1.1,
    candidateClusters:    candidateClusters,
  }
}

func (this *KMeansStream) AddDataPoint(data []float64) {
  this.addDataPointWeighted(data, 1);
}
//Add a new data point to the stream
func (this *KMeansStream) addDataPointWeighted(data []float64, weight int64) {
  if len(data) != this.dimensionality {
    return
    // panic("The input data does not have the correct dimenstionality")
  }
  minIndex := 0;
  minDist := 0.0;
  for i, centriod := range this.candidateClusters {
    currDist := utils.Distance(data,centriod.Centroid())
    if i == 0 || minDist > currDist {
      minDist = currDist;
      minIndex = i;
    }
  }
  minDistSquared := minDist * minDist
  if len(this.candidateClusters) < this.k || rand.Float64() < float64(weight) * (minDistSquared/this.frequency) {
    this.candidateClusters = append(this.candidateClusters, *itemset.NewCentroidWeighted(data, weight))
  }else{
    this.candidateClusters[minIndex].UpdateVector(data)
  }
  if len(this.candidateClusters) > this.maxCandidateClusters {
    this.reduceCandidateClusters()
  }
}

func (this *KMeansStream) reduceCandidateClusters() {
  this.frequency = this.frequency * this.frequencyChange;
  oldCandidateClusters := make([]itemset.Centroid, len(this.candidateClusters), len(this.candidateClusters))
  copy(oldCandidateClusters, this.candidateClusters)
  this.candidateClusters = []itemset.Centroid{};
  for _, centriod := range oldCandidateClusters {
    this.addDataPointWeighted(centriod.Centroid(), centriod.GetCount());
  }
}

func (this *KMeansStream) GetCentroids() [][]float64 {
  data := make([][]float64, len(this.candidateClusters), len(this.candidateClusters))
  for i := 0; i < len(this.candidateClusters); i++ {
    data[i] = this.candidateClusters[i].Centroid();
  }
  simple := NewKMeansSimple(this.k, data);
  return simple.GetCentroids();
}
