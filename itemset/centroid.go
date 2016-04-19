package itemset

import (
  "github.com/wilseypa/rphash-golang/types"
  "github.com/wilseypa/rphash-golang/utils"
)

type Centroid struct {
  vec   []float64
  count int64
  ids   types.HashSet
  id    int64
}

func NewCentroidStream(data []float64) *Centroid {
  return NewCentroidWeighted(data, 1);
}

func NewCentroidWeighted(data []float64, weight int64) *Centroid {
  return &Centroid{
    vec:   data,
    ids:   utils.NewHash64Set(),
    count: weight,
    id:    0,
  }
}

func NewCentroidSimple(dim int, lsh int64) *Centroid {
  data := make([]float64, dim)
  ids := utils.NewHash64Set()
  ids.Add(lsh)
  return &Centroid{
    vec:   data,
    ids:   ids,
    count: 0,
    id:    lsh,
  }
}

func (this *Centroid) UpdateVector(data []float64) {
  var delta, x float64
  this.count++
  for i := 0; i < len(data); i++ {
    x = data[i]
    delta = x - this.vec[i]
    this.vec[i] = this.vec[i] + delta/float64(this.count)
  }
}

func (this *Centroid) Centroid() []float64 {
  return this.vec
}

func (this *Centroid) GetCount() int64 {
  return this.count
}

func (this *Centroid) GetID() int64 {
  return this.id
}

func (this *Centroid) GetIDs() types.HashSet {
  return this.ids
}

func (this *Centroid) AddID(h int64) {
  if this.ids.Length() == 0 {
    this.id = h
  }
  this.ids.Add(h)
}
