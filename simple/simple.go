package simple;

import (
    "math"
    "github.com/wenkesj/rphash/types"
    "github.com/wenkesj/rphash/defaults"
);

type Simple struct {
    centroids [][]float64;
    variance float64;
    rphashObject types.RPHashObject;
    LSH types.LSH;
    CountMinSketch types.CountItemSet;
};

func NewSimple(_rphashObject types.RPHashObject) *Simple {
    return &Simple{
        variance: 0,
        centroids: nil,
        rphashObject: _rphashObject,
        LSH: nil,
        CountMinSketch: nil,
    };
};

func (this *Simple) Map() *Simple {
    vecs := this.rphashObject.GetVectorIterator();
    if !vecs.HasNext() {
        return this;
    }
    var hashResult int64;
    vec := vecs.Next();
    if this.LSH == nil {
        hash := defaults.NewHash(this.rphashObject.GetHashModulus());
        decoder := this.rphashObject.GetDecoderType();
        projector := defaults.NewProjector(this.rphashObject.GetDimensions(), decoder.GetDimensionality(), this.rphashObject.GetRandomSeed());
        this.LSH = defaults.NewLSH(hash, decoder, projector);
        k := int(float64(this.rphashObject.GetK()) * math.Log(float64(this.rphashObject.GetK())));
        this.CountMinSketch = defaults.NewCountMinSketch(k);
    }

    for vecs.HasNext() {
        hashResult = this.LSH.LSHHashSimple(vec);
        this.CountMinSketch.Add(hashResult);
        vec = vecs.Next();
    }

    this.rphashObject.SetPreviousTopID(this.CountMinSketch.GetTop());
    vecs.Reset();
    return this;
};

func (this *Simple) Reduce() *Simple {
    vecs := this.rphashObject.GetVectorIterator();
    if !vecs.HasNext() {
        return this;
    }
    var centroids []types.Centroid;
    vec := vecs.Next();
    for i := 0; i < this.rphashObject.GetK(); i++ {
        centroid := defaults.NewCentroidSimple(this.rphashObject.GetDimensions(), this.rphashObject.GetPreviousTopID()[i]);
        centroids = append(centroids, centroid);
    }
    // Iterate over the dataset and check if the CountMinSketch contains the count.
    for vecs.HasNext() {
        var hashResult = this.LSH.LSHHashSimple(vec);
        for _, cent := range centroids {
            if cent.GetIDs().Contains(hashResult) {
                cent.UpdateVector(vec);
                break;
            }
        }
        vec = vecs.Next();
    }
    for _, cent := range centroids {
        this.rphashObject.AddCentroid(cent.Centroid());
    }
    vecs.Reset();
    return this;
};

func (this *Simple) GetCentroids() [][]float64 {
    if this.centroids == nil {
        this.Run();
    }
    return defaults.NewKMeansSimple(this.rphashObject.GetK(), this.centroids).GetCentroids();
};

/**
 * Map->LSH Reduce->Centroids.
 */
func (this *Simple) Run() {
    this.Map().Reduce();
    this.centroids = this.rphashObject.GetCentroids();
}

func (this *Simple) GetRPHash() types.RPHashObject {
    return this.rphashObject;
};
