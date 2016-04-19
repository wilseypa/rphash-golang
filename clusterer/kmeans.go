package clusterer

import (
  "fmt"
  "github.com/wilseypa/rphash-golang/projector"
  "github.com/wilseypa/rphash-golang/types"
  "github.com/wilseypa/rphash-golang/utils"
  "log"
  "math/rand"
)

type KMeans struct {
  k                   int
  n                   int
  data                [][]float64
  projectionDimension int
  means               [][]float64
  clusters            [][]int //Each row of clusters contatins all vectors in the data currently assigned to it.
  weights             []int64
}

func NewKMeansSimple(k int, data [][]float64) *KMeans {
  weights := make([]int64, len(data), len(data))
  for i := 0; i < len(data); i++ {
    weights[i] = int64(1)
  }
  return NewKMeansWeighted(k, data, weights)
}

func NewKMeansWeighted(k int, data [][]float64, weights []int64) *KMeans {
  if len(data) == 0 {
    log.Panic(data)
  }
  return &KMeans{
    k:                   k,
    data:                data,
    projectionDimension: 0,
    clusters:            nil,
    weights:             weights,
  }
}

//Vectors is a list of all assignedVectors currently assigned to the centriod we are computing
func (this *KMeans) ComputeCentroid(assignedVectors []int, data [][]float64) []float64 {
  d := len(data[0])
  centroid := make([]float64, d, d)
  for i := 0; i < d; i++ {
    centroid[i] = 0.0
  }
  var w_total int64 = 0
  for _, v := range assignedVectors {
    w_total += this.weights[v]
  }
  for _, v := range assignedVectors {
    vec := data[v]
    weight := float64(this.weights[v]) / float64(w_total)
    for i := 0; i < d; i++ {
      centroid[i] += (vec[i] * weight)
    }
  }
  return centroid
}

func (this *KMeans) UpdateMeans(data [][]float64) {
  for i := 0; i < this.k; i++ {
    this.means[i] = this.ComputeCentroid(this.clusters[i], data)
  }
}

func (this *KMeans) AssignClusters(data [][]float64) int {
  swaps := 0
  newClusters := [][]int{}
  for j := 0; j < this.k; j++ {
    newClusterList := []int{}
    newClusters = append(newClusters, newClusterList)
  }
  for clusterid := 0; clusterid < this.k; clusterid++ {
    for _, member := range this.clusters[clusterid] {
      nearest, _ := utils.FindNearestDistance(data[member], this.means)
      newClusters[nearest] = append(newClusters[nearest], member)
      if nearest != clusterid {
        swaps++
      }
    }
  }
  this.clusters = newClusters
  return swaps
}

func (this *KMeans) Run() {
  //This is a condition to avoid infinite Run..
  maxiters := 10000
  swaps := 3
  fulldata := this.data
  data := make([][]float64, 0)
  var p types.Projector = nil
  if this.projectionDimension != 0 {
    p = projector.NewDBFriendly(len(fulldata[0]), this.projectionDimension, rand.Int63())
  }
  for _, v := range fulldata {
    if p != nil {
      data = append(data, p.Project(v))
    } else {
      data = append(data, v)
    }
  }
  this.n = len(data)
  this.means = make([][]float64, this.k)
  for i := 0; i < this.k; i++ {
    this.means[i] = data[i*(this.n/this.k)]
  }
  this.clusters = make([][]int, this.k)
  //initilize cluster lists to be evenly diveded sequentailly
  for i := 0; i < this.k; i++ {
    cluster := make([]int, this.n/this.k)
    clusterStart := i * (this.n / this.k)
    for j := 0; j < this.n/this.k; j++ {
      cluster[j] = j + clusterStart
    }
    this.clusters[i] = cluster
  }
  for swaps > 2 && maxiters > 0 {
    maxiters--
    this.UpdateMeans(data)
    swaps = this.AssignClusters(data)
  }
  if maxiters == 0 {
    fmt.Println("Warning: Max Iterations Reached")
  }
  data = fulldata
  this.UpdateMeans(data)
}

func (this *KMeans) GetCentroids() [][]float64 {
  if this.means == nil {
    this.Run()
  }
  return this.means
}
