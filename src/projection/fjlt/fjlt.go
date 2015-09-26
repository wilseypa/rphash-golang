/**
 * Fast Johnson Lindenstrauss Transform
 * 1st Step
 * @author Sam Wenke
 * @author Jacob Franklin
 * @reference Lee Carraher
 */

package fjlt;

import (
    "math"
    "math/rand"
);

type FJLTProjection struct {
    n int64;
    k int64;
    d int64;
    D []float64;
    P []float64;
    random *rand.Rand;
};

/**
 * Allocate a new instance of the FJLTProjection.
 * @return {*FJLTProjection}.
 */
func New(d, k, n int64) *FJLTProjection {
    random := rand.New(rand.NewSource(n));
    epsilon := float64(math.Sqrt(math.Log(float64(n)) / float64(k)));
    P := GenerateP(n, k, d, 2, epsilon, random);
    D := GenerateD(d, random);
    return &FJLTProjection{
        n: n,
        k: k,
        d: d,
        D: D,
        P: P,
    };
};

func SGEMV(t, n, startpoint, startoutput int64, M, v, result []float64, alpha float64) {
    var sum float64;
    var i int64;
    var j int64;
    for i = 0; i < t; i++ {
        sum = 0.0;
        for j = 0; j < n; j++ {
            sum += v[j + startpoint] * M[i * n + j];
            result[startoutput + i] = sum * alpha;
        }
    }
};

func GenerateP(n, k, d, p int64, e float64, random *rand.Rand) []float64 {
    var i int64;
    var j int64;
    data := make([]float64, k * d);
    q := float64((math.Pow(e, float64(p - 2)) * math.Pow(math.Log(float64(n)), float64(p))) / float64(d));
    if !(q < 1) {
        q = 1;
    }
    rdata := make([]float64, k * d);
    InvRandN(data, k, d, 0, 1 / float64(q), random);
    RandU(rdata, k, d, random);
    for i = 0; i < k; i++ {
        for j = 0; j < d; j++ {
            if rdata[i * d + j] < q {
                data[i * d + j] *= 0;
            } else {
                data[i * d + j] *= 1;
            }
        }
    }
    return data;
};

func GenerateD(d int64, random *rand.Rand) []float64 {
    var l int64;
    var i int64;
    var j int64;
    data := make([]float64, d);
    for i = 0; i < d; {
        l = random.Int63();
        for j = 0; j < 32 && i < d; j++ {
            if (l & 1) == 1 {
                data[i] = 1;
            } else {
                data[i] = -1;
            }
            l = l >> 1;
            i++;
        }
    }
    return data;
};

func InvRandN(data []float64, m, n int64, mu, vari float64, random *rand.Rand) {
    var i int64;
    var j int64;
    sd := float64(math.Sqrt(vari));
    for i = 0; i < m; i++ {
        for j = 0; j < n; j++ {
            data[i * n + j] = mu + sd * float64(MoroInvCND(random.Float64()));
        }
    }
};

func RandU(data []float64, m, n int64, random *rand.Rand) {
    var i int64;
    var j int64;
    for i = 0; i < m; i++ {
        for j = 0; j < n; j++ {
            data[i * n + j] = random.Float64();
        }
    }
};

func MoroInvCND(P float64) float64 {
    var z float64;
    a1 := 2.50662823884;
    a2 := -18.61500062529;
    a3 := 41.39119773534;
    a4 := -25.44106049637;
    b1 := -8.4735109309;
    b2 := 23.08336743743;
    b3 := -21.06224101826;
    b4 := 3.13082909833;
    c1 := 0.337475482272615;
    c2 := 0.976169019091719;
    c3 := 0.160797971491821;
    c4 := 2.76438810333863E-02;
    c5 := 3.8405729373609E-03;
    c6 := 3.951896511919E-04;
    c7 := 3.21767881768E-05;
    c8 := 2.888167364E-07;
    c9 := 3.960315187E-07;

    if P <= 0 || P >= 1.0 {
        /* Caused by numerical instability of random */
        P = 0.9999;
    }
    y := P - 0.5;
    if math.Abs(y) < 0.42 {
        z = y * y;
        z = y * (((a4 * z + a3) * z + a2) * z + a1) / ((((b4 * z + b3) * z + b2) * z + b1) * z + 1);
    } else {
        if y > 0 {
            z = float64(math.Log(-math.Log(1.0 - P)));
        } else {
            z = float64(math.Log(-math.Log(P)));
            z = c1 + z * (c2 + z * (c3 + z * (c4 + z * (c5 + z * (c6 + z * (c7 + z * (c8 + z * c9)))))));
            if y < 0 {
                z = -z;
            }
        }
    }
    return z;
};

/**
 * Performs the FJLT on a matrix.
 * @class {FJLTProjection} _fjlt.
 * @param {[]float64} input, Matrix.
 * @return {[]float64} new Matrix.
 */
func (_fjlt *FJLTProjection) FJLT(input []float64) []float64 {
    var curr int64;
    var a uint64;
    var b uint64;
    var c uint64;
    result := make([]float64, _fjlt.n * _fjlt.k);
    for curr = 0; curr < _fjlt.n; curr++ {
        startpoint := curr * _fjlt.d;
        startoutput := _fjlt.k * curr;
        for a = 0; a < uint64(_fjlt.d); a++ {
            input[int64(a) + startpoint] *= _fjlt.D[a];
        }
        l2 := uint64(math.Log(float64(_fjlt.d))/math.Log(2));
        for a = 0; a < l2; a++ {
            for b = 0; b < (1 << l2); b += (1 << (a + 1)) {
                for c = 0; c < (1 << a); c++ {
                    temp := input[startpoint + int64(b + c)];
                    input[startpoint + int64(b + c)] += input[startpoint + int64(b + c + (1 << a))];
                    input[startpoint + int64(b + c + (1 << a))] = temp - input[startpoint + int64(b + c + (1 << a))];
                }
            }
        }
        SGEMV(_fjlt.k, _fjlt.d, startpoint, startoutput, _fjlt.P, input, result, 1.0/float64(_fjlt.d));
    }
    return result;
};

/**
 * Project a matrix.
 * @class {FJLTProjection} _fjlt.
 * @param {[]float64} input, Matrix.
 */
func (_fjlt *FJLTProjection) Project(input []float64) []float64 {
    var a uint64;
    var b uint64;
    var c uint64;
    result := make([]float64, _fjlt.k);
    for a = 0; a < uint64(_fjlt.d); a++ {
        input[a] *= _fjlt.D[a];
    }
    l2 := uint64(math.Log(float64(_fjlt.d))/math.Log(2));
    for a = 0; a < l2; a++ {
        for b = 0; b < (1 << l2); b += (1 << (a + 1)) {
            for c = 0; c < (1 << a); c++ {
                temp := input[b + c];
                input[b + c] += input[b + c + (1 << a)];
                input[b + c + (1 << a)] = temp - input[b + c + (1 << a)];
            }
        }
    }
    SGEMV(_fjlt.k, _fjlt.d, 0, 0, _fjlt.P, input, result, 1.0/float64(_fjlt.d));
    return result;
};
