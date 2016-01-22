package simple;

import (
    "math"
    "fmt"
    "github.com/wenkesj/rphash/types"
    "github.com/wenkesj/rphash/defaults"
);

type Simple struct {
    centroids [][]float64;
    variance float64;
    rphashObject types.RPHashObject;
};

func NewSimple(_rphashObject types.RPHashObject) *Simple {
    return &Simple{
        variance: 0,
        centroids: nil,
        rphashObject: _rphashObject,
    };
};

func (this *Simple) Map() types.RPHashObject {
    hash := defaults.NewHash(this.rphashObject.GetHashModulus());
    vecs := this.rphashObject.GetVectorIterator();
    if !vecs.HasNext() {
        return this.rphashObject;
    }
    decoder := this.rphashObject.GetDecoderType();
    projector := defaults.NewProjector(this.rphashObject.GetDimensions(), decoder.GetDimensionality(), this.rphashObject.GetRandomSeed());
    lshfunc := defaults.NewLSH(hash, decoder, projector);
    var hashResult int64;
    k := int(float64(this.rphashObject.GetK()) * math.Log(float64(this.rphashObject.GetK())));
    countMin := defaults.NewCountMinSketch(k);
    for vecs.HasNext() {
        vec := vecs.Next();
        hashResult = lshfunc.LSHHashSimple(vec);
        countMin.Add(hashResult);
    }
    this.rphashObject.SetPreviousTopID(countMin.GetTop());
    vecs.Reset();
    return this.rphashObject;
};

func (this *Simple) Reduce() types.RPHashObject {
    vecs := this.rphashObject.GetVectorIterator();
    if !vecs.HasNext() {
        return this.rphashObject;
    }
    vec := vecs.Next();
    hash := defaults.NewHash(this.rphashObject.GetHashModulus());
    decoder := this.rphashObject.GetDecoderType();
    projector := defaults.NewProjector(this.rphashObject.GetDimensions(), decoder.GetDimensionality(), this.rphashObject.GetRandomSeed());
    lshfunc := defaults.NewLSH(hash, decoder, projector);
    var centroids []types.Centroid;
    for i := 0; i < this.rphashObject.GetK(); i++ {
        centroids = append(centroids, defaults.NewCentroidSimple(this.rphashObject.GetDimensions(), this.rphashObject.GetPreviousTopID()[i]));
    }
    // JF 1/22/16 not sure about this change
    //for _, id := range this.rphashObject.GetPreviousTopID() {
    //    centroids = append(centroids, defaults.NewCentroidSimple(this.rphashObject.GetDimensions(), id));
    //}
    for vecs.HasNext() {
        var hashResult = lshfunc.LSHHashSimple(vec);
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
    return this.rphashObject;
};

func (this *Simple) GetCentroids() [][]float64 {
    if this.centroids == nil {
        this.Run();
    }
    fmt.Println("Centriod: ", this.centroids);
    fmt.Println("K: ", this.rphashObject.GetK());
    return defaults.NewKMeansSimple(this.rphashObject.GetK(), this.centroids).GetCentroids();
};

func (this *Simple) Run() {
    this.Map();
    this.Reduce();
    this.centroids = this.rphashObject.GetCentroids();
}

func (this *Simple) GetParam() types.RPHashObject {
    return this.rphashObject;
};
