/**
 * First Step,
 * @author Sam Wenke
 * @author Jacob Franklin
 */
package rphash;

import (
    "log"
    "math"
    "math/rand"
);

type RandomProjection struct {
    M [][]int;
    P [][]int;
    n int;
    t int;
    random *rand.Rand;
};

/**
 * Allocate a new instance of RandomProjection.
 * @param {[]int} P - The size [t x n/6] set of vector indices that should be positive (+sqrt(3/t) => +1)
 * @param {[]int} M - The size [t x n/6] set of vector indices that should incur negative (-sqrt(3/t) => -1)
 * @param {int} n - Original dimension
 * @param {int} t - Target/Projected dimension
 */
func NewRandomProjection(n, t int, randomseed int64) *RandomProjection {
    rando := rand.New(rand.NewSource(randomseed));
    M, P := make([][]int, t), make([][]int, t);
    r1, r2 := 0, 0;
    probability := n / 6;
    for i := 0; i < t; i++ {
        ordered1, ordered2 := make([]int, probability), make([]int, probability);
        for j := 0; j < n; j++ {
            r1, r2 = rando.Intn(6), rando.Intn(6);
            if r1 == 0 {
                ordered1 = append(ordered1, j);
            } else if r2 == 0 {
                ordered2 = append(ordered2, j);
            }
        }
        tmp1, tmp2 := make([]int, len(ordered1)), make([]int, len(ordered2));
        for k, val := range ordered1 {
            tmp1[k] = val;
        }
        for k, val := range ordered2 {
            tmp2[k] = val;
        }
        M[i], P[i] = tmp1, tmp2;
    }

    return &RandomProjection{
        M: M,
        P: P,
        n: n,
        t: t,
        random: rando,
    };
};

/**
 * Project onto a random matrix of {-1, 1} to produce a reduced dimensional vector.
 * @param {[]float64} v - The input vector
 * @return {[]float64} - Returns a reduced dimensional vector
 */
func (rp *RandomProjection) Project(v []float64) []float64 {
    if (len(v) < 2) {
        log.Fatal("Input Vector must have more than 1 element");
        return nil;
    }
    var sum float64;
    var tmp []int;
    r := make([]float64, rp.t);
    scale := math.Sqrt(3.0 / float64(rp.t));
    for i := 0; i < rp.t; i++ {
        sum = 0.0;
        tmp = rp.M[i];
        for _, val := range tmp {
            sum -= v[val] * scale;
        }
        tmp = rp.P[i];
        for _, val := range tmp {
            sum += v[val] * scale;
        }
        r[i] = sum;
    }
    return r;
};
