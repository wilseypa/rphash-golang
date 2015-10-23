package decoder;

import (
    "math"
    "github.com/wenkesj/rphash/utils"
);

var HashBits int = 64;

type Spherical struct {
    vAll [][][]float64;
    hbits int;
    d int;
    k int;
    l int;
    distance float64;
    variance float64;
};

func NewSpherical(d, k, L int) *Spherical {
    nvertex := 2.0 * d;
    hbits := int(math.Ceil(math.Log(float64(nvertex)) / math.Log(2)));
    kmax := int(HashBits / hbits);
    if (k > kmax) {
        k = kmax;
    }
    vAll := make([][][]float64, k * L);
    // r := make([]*rand.Rand, d);
    // for i := 0; i < d; i++ {
    //     r[i] = rand.New();
    // }
    // rotationMatrices := vAll;
    // for i := 0; i < k * L; i++ {
    //     rotationMatrices[i] = utils.RandomRotation(d, r);
    // }
    return &Spherical{
        vAll: vAll,
        hbits: hbits,
        d: d,
        k: k,
        l: L,
        distance: 0.0,
        variance: 1.0,
    };
};

func (this *Spherical) GetDimensionality() int {
    return this.d;
};

func (this *Spherical) GetErrorRadius() float64 {
    return float64(this.d);
};

func (this *Spherical) GetDistance() float64 {
    return this.distance;
};

func (this *Spherical) Hash(p []float64) []int64 {
    ri := 0;
    var h int64;
    g := make([]int64, this.l);
    for i := 0; i < this.l; i++ {
        g[i] = 0;
        for j := 0; j < this.k; j++ {
            vs := this.vAll[ri];
            h = utils.Argmaxi(p, vs, this.d);
            g[i] |= (h << (uint(this.hbits * j)));
            ri++;
        }
    }
    return g;
};

func (this *Spherical) GetVariance() float64 {
    return this.variance;
};

func (this *Spherical) SetVariance(parameterObject float64) {
    this.variance = parameterObject;
};

func (this *Spherical) Decode(f []float64) []int64 {
    return this.Hash(utils.Normalize(f));
};

func InnerDecoder() *Spherical {
    return NewSpherical(32, 3, 1);
};
