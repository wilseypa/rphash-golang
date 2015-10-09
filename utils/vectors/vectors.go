package vectors;

import (
    "math"
    "math/rand"
);

func Normalize(x []float64) []float64 {
    var length float64 = 0;
    for i := 0; i < len(x); i++ {
        length += (x[i] * x[i]);
    }
    length = math.Sqrt(length);
    ret := make([]float64, len(x));
    for i := 0; i < len(x); i++ {
        ret[i] = x[i] / length;
    }
    return ret;
};

func Random(d int, r []*rand.Rand) []float64 {
    v := make([]float64, d);
    for i := 0; i < d; i++ {
        v[i] =  r[i].NormFloat64();
    }
    return v;
};

func RandomRotation(d int, r2 []*rand.Rand) [][]float64 {
    R := make([][]float64, d);
    for i := 0; i < d; i++ {
        R[i] = Random(d, r2);
        u := R[i];
        for j := 0; j < i; j++ {
            v := R[j];
            vnorm := Norm(v);
            if (vnorm == 0) {
                return RandomRotation(d, r2);
            }
            vs := make([]float64, len(v));
            copy(vs, v);
            Scale(vs, Dot(v, u) / vnorm);
            u = Sub(u, vs);
        }
        u = Scale(u, 1.0 / Norm(u));
    }

    return R;
}

func Argmaxi(p []float64, vs [][]float64, d int) int32 {
    var maxi int32 = 0;
    var max float64 = 0;
    var abs float64;
    for i := 0; i < d; i++ {
        dot := Dot(p, vs[i]);
        if dot >= 0 {
            abs = dot;
        } else {
            abs = -dot;
        }
        if (abs < max) {
            continue;
        }
        max = abs;
        if dot >= 0 {
            maxi = int32(i);
        } else {
            maxi = int32(i + d);
        }
    }
    return maxi;
};

func Norm(t []float64) float64 {
    var n float64 = 0;
    for i := 0; i < len(t); i++ {
        n += t[i] * t[i];
    }
    return math.Sqrt(n);
};

func Scale(t []float64, s float64) []float64 {
    for i := 0; i < len(t); i++ {
        t[i] *= s;
    }
    return t;
};

func Dot(t, u []float64) float64 {
    var s float64 = 0;
    for i := 0; i < len(t); i++ {
        s += t[i] * u[i];
    }
    return s;
};

func Sub(t, u []float64) []float64 {
    for i := 0; i < len(t); i++ {
        t[i] -= u[i];
    }
    return t;
};

func Max(vec []int32) (a int32) {
    return a;
};

func Min(vec []int32) (a int32) {
    return a;
};
