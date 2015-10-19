package simple;

import (
    "github.com/wenkesj/rphash/types"
    "github.com/wenkesj/rphash/defaults"
);

type Simple struct {
    centroids [][]float64;
    variance float64;
    rphashObject types.RPHashObject;
};

func NewSimple(_rphashObject types.RPHash) *Simple {
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
    projector := defaults.NewProjector(this.rphashObject.GetDimension(), decoder.GetDimensionality(), this.rphashObject.GetRandomSeed());
    lshfunc := defaults.NewLSH(decoder, projector, hash);
    var hash int32;
    k := int(this.rphashObject.GetK() * math.Log(this.rphashObject.GetK()));
    countMin := defaults.NewCentroidCounter(k);
    for vecs.HasNext() {
        vec := vecs.Next();
        hash = lshfunc.LSHHashSimple(vec);
        countMin.Add(hash);
    }
    this.rphashObject.SetPreviousTopID(countMin.GetTop());
    return this.rphashObject;
};

func (this *Simple) Reduce() types.RPHashObject {
    vecs := this.rphashObject.GetVectorIterator();
    if !vecs.HasNext() {
        return this.rphashObject;
    }
    vec := vecs.Next();
    blurValue := this.rphashObject.GetNumberOfBlurs();
    hash := defaults.NewHash(this.rphashObject.GetHashModulus());
    decoder := this.rphashObject.GetDecoderType();
    projector := defaults.NewProjector(this.rphashObject.GetDimension(), decoder.GetDimensionality(), this.rphashObject.GetRandomSeed());
    lshfunc := defaults.NewLSH(decoder, projector, hash);
    var hash []int32;
    var centroids []types.Centroid;
    for _, id := range this.rphashObject.GetPreviousTopID() {
        centroids = append(centroids, defaults.NewCentroid(this.rphashObject.GetDimension(), id));
    }
    for vecs.HasNext() {
        hash = lshfunc.LSHHashStream(vec, blurValue);
        for _, cent := range centroids {
            for h := range hash {
                if cent.ids.Contains(h) {
                    cent.UpdateVector(vec);
                    break;
                }
            }
        }
        vec = vecs.Next();
    }
    for _, cent := range centroids {
        this.rphashObject.AddCentroid(cent.Centroid());
    }
    return this.rphashObject;
};

func (this *Simple) GetCentroids() [][]float64 {
    if this.centroids == nil {
        this.Run();
    }
    return defaults.NewKMeans(this.streamObject.GetK(), this.centroids).GetCentroids();
};

func (this *Simple) Run() {
    this.Map();
    this.Reduce();
    this.centroids = this.streamObject.GetCentroids();
}

func (this *Simple) GetParam() types.RPHashObject {
    return this.streamObject;
};
