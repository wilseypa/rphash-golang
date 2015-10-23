package lsh;

import (
    "github.com/wenkesj/rphash/types"
);

type LSH struct {
    hash types.Hash;
    decoder types.Decoder;
    projector types.Projector;
    distance float64;
    noise [][]float64;
};

func NewLSH(hash types.Hash,
            decoder types.Decoder,
            projector types.Projector) *LSH {

    return &LSH{
        hash: hash,
        decoder: decoder,
        projector: projector,
        distance: 0.0,
        noise: nil,
    };
};

func (this *LSH) LSHHashStream(r []float64, times int) (a []int64) {
    return a;
};

func (this *LSH) LSHHashSimple(r []float64) int64 {
    return this.hash.Hash(this.decoder.Decode(this.projector.Project(r)));
};

func (this *LSH) Distance() float64 {
    return this.distance;
};

func (this *LSH) UpdateDecoderVariance(vari float64) {
    this.decoder.SetVariance(vari);
};
