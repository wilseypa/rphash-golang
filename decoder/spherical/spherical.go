package spherical;

import (
    "math"
    // "math/rand"
    "github.com/wenkesj/rphash/utils/vectors"
);

/* vAll - Random rotation matrices.
 * d - Dimension of the feature space.
 * k - Number of elementary hash functions (h) to be concatenated to
 * obtain a reliable enough hash function (g). LSH queries becomes
 * more selective with increasing k, due to the reduced the
 * probability of collision.
 * l - Number of "copies" of the bins (with a different random matrices).
 * Increasing L will increase the number of points the should be
 * scanned linearly during query. */
var HashBits int = 64;
var distance float64 = 0.0;
var variance float64 = 1.0;

type Spherical struct {
    vAll [][][]float64;
    hbits int;
    d int;
    k int;
    l int;
};

/* d - number of dimensions
 * k - number of elementary hash functions
 * L - number of copies to search */
func New(d, k, L int) *Spherical {
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
    //     rotationMatrices[i] = vectors.RandomRotation(d, r);
    // }
    return &Spherical{
        vAll: vAll,
        hbits: hbits,
        d: d,
        k: k,
        l: L,
    };
};

func (this *Spherical) GetDimensionality() int {
    return this.d;
};

func (this *Spherical) GetErrorRadius() float64 {
    return float64(this.d);
};

func (this *Spherical) GetDistance() float64 {
    return distance;
};

func (this *Spherical) Hash(p []float64) []int32 {
    ri := 0;
    var h int32;
    g := make([]int32, this.l);
    for i := 0; i < this.l; i++ {
        g[i] = 0;
        for j := 0; j < this.k; j++ {
            vs := this.vAll[ri];
            h = vectors.Argmaxi(p, vs, this.d);
            g[i] |= (h << (uint(this.hbits * j)));
            ri++;
        }
    }
    return g;
};

func SetVariance(parameterObject float64) {
    variance = parameterObject;
};

func (this *Spherical) Decode(f []float64) []int32 {
    return this.Hash(vectors.Normalize(f));
};
