package itemset;

import (
    "github.com/wenkesj/rphash/utils"
    "github.com/wenkesj/rphash/types"
);

type Centroid struct {
    vec []float64;
    count int64;
    ids types.HashSet;
    id int64;
};

func NewCentroidStream(data []float64) *Centroid {
    return &Centroid{
        vec: data,
        ids: utils.NewHash64Set(),
        count: 1,
        id: 0,
    };
};

func NewCentroidSimple(dim int, id int64) *Centroid {
    data := make([]float64, dim);
    ids := utils.NewHash64Set();
    ids.Add(id);
    return &Centroid{
        vec: data,
        ids: ids,
        count: 0,
        id: id,
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

/* The number of vectors in the centroid.
 */
func (this *Centroid) GetCount() int64 {
    return this.count;
};

/* ID Represents the LSH of the center vector of the centroid.
 */
func (this *Centroid) GetID() int64 {
    return this.id;
};

func (this *Centroid) GetIDs() types.HashSet {
    return this.ids;
};

func (this *Centroid) AddID(h int64) {
    if this.ids.Length() == 0 {
        this.id = h;
    }
    this.ids.Add(h);
};
