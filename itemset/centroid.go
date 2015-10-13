package itemset;

import (
    "github.com/wenkesj/rphash/utils"
);

type Centroid struct {
    vec []float64;
    count int32;
    ids *utils.Hash32Set;
    id int32;
};

func NewCentroid(data []float64) *Centroid {
    return &Centroid{
        vec: data,
        ids: utils.NewHash32Set(),
        count: 1,
        id: 0,
    };
};

func (this *Centroid) UpdateCentroidVector(data []float64) {
	var delta, x float64;
	this.count++;
	for i := 0; i < len(data); i++ {
		x = data[i];
		delta = x - this.vec[i];
		this.vec[i] = this.vec[i] + delta / float64(this.count);
	}
};

func (this *Centroid) Centroid() []float64{
	return this.vec;
};

func (this *Centroid) UpdateVector(rp []float64) {
	this.UpdateCentroidVector(rp);
};

func (this *Centroid) GetCount() int32 {
	return this.count;
};

func (this *Centroid) AddID(h int32) {
	if this.ids.Length() == 0 {
        this.id = h;
    }
	this.ids.Add(h);
};
