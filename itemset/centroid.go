package itemset

import (
	"github.com/wilseypa/rphash-golang/types"
	"github.com/wilseypa/rphash-golang/utils"
)

type Centroid struct {
	Vec   []float64
	Count int64
	Ids   *utils.Hash64Set
	Id    int64
}

func NewCentroidStream(data []float64) *Centroid {
	return NewCentroidWeighted(data, 1)
}

func NewCentroidWeighted(data []float64, weight int64) *Centroid {
	return &Centroid{
		Vec:   data,
		Ids:   utils.NewHash64Set(),
		Count: weight,
		Id:    0,
	}
}

func NewCentroidSimple(dim int, lsh int64) *Centroid {
	data := make([]float64, dim)
	Ids := utils.NewHash64Set()
	Ids.Add(lsh)
	return &Centroid{
		Vec:   data,
		Ids:   Ids,
		Count: 0,
		Id:    lsh,
	}
}

func (this *Centroid) UpdateVector(data []float64) {
	var delta, x float64
	this.Count++
	for i := 0; i < len(data); i++ {
		x = data[i]
		delta = x - this.Vec[i]
		this.Vec[i] = this.Vec[i] + delta/float64(this.Count)
	}
}

func (this *Centroid) Centroid() []float64 {
	return this.Vec
}

func (this *Centroid) GetCount() int64 {
	return this.Count
}

func (this *Centroid) GetID() int64 {
	return this.Id
}

func (this *Centroid) GetIDs() types.HashSet {
	return this.Ids
}

func (this *Centroid) AddID(h int64) {
	if this.Ids.Length() == 0 {
		this.Id = h
	}
	this.Ids.Add(h)
}
