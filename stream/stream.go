package stream;

import (
    "math/rand";
    "github.com/wenkesj/rphash/types"
);

type Stream struct {
    k int;
    data []float64;
    cc types.CentroidCounter;
    random *rand.Rand;
    rphash types.RPHashObject;
    lsh types.LSH;
    _lsh []types.LSH;
    decoder types.Decoder;
    projector types.Projector;
    hash types.Hash;
};

func New(k int, data []float64, randomseed int64,
            rphash types.RPHashObject,
            cc types.CentroidCounter,
            decoder types.Decoder,
            lsh types.LSH,
            projector types.Projector,
            hash types.Hash) *Stream {

    rp := rphash.New();
    return &Stream{
        k: k,
        data: data,
        cc: cc.New(k),
        random: rand.New(rand.NewSource(randomseed)),
        rphash: rp,
        _lsh: make([]types.LSH, rp.GetNumberOfProjections()),
        lsh: lsh,
        hash: hash.New(rp.GetHashModulus()),
        decoder: decoder.New(),
        projector: projector,
    };
};

func (this *Stream) Initialize() {
    projections := this.rphash.GetNumberOfProjections();
    this.k = this.k * projections;
    for i := 0; i < projections; i++ {
        projector := this.projector.New(this.rphash.GetDimensions(),
                                            this.decoder.GetDimensionality(),
                                            this.random.Int63());
        this._lsh[i] = this.lsh.New(this.hash,
                                        this.decoder,
                                        projector);
    }
};
