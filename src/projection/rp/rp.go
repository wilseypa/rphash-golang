/**
 * Random Projection
 * 1st Step
 * @author Sam Wenke
 * @author Jacob Franklin
 */
package rp;

import (
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
 * @param {[]int} P - The size [t x n/6] set of vector indices that should be positive (+sqrt(3/t) => +1).
 * @param {[]int} M - The size [t x n/6] set of vector indices that should incur negative (-sqrt(3/t) => -1).
 * @param {int} n - Original dimension.
 * @param {int} t - Target/Projected dimension.
 */
func New(n, t int, randomseed int64) *RandomProjection {
    rando := rand.New(rand.NewSource(randomseed));
    M, P := make([][]int, t), make([][]int, t);
    rM, rP := 0, 0;
    probability := n / 6;
    for i := 0; i < t; i++ {
        orderedM, orderedP := make([]int, probability), make([]int, probability);
        for j := 0; j < n; j++ {
            rM, rP = rando.Intn(6), rando.Intn(6);
            if rM == 0 {
                orderedM = append(orderedM, j);
            } else if rP == 0 {
                orderedP = append(orderedP, j);
            }
        }
        tmpM, tmpP := make([]int, len(orderedM)), make([]int, len(orderedP));
        for k, val := range orderedM {
            tmpM[k] = val;
        }
        for k, val := range orderedP {
            tmpP[k] = val;
        }
        M[i], P[i] = tmpM, tmpP;
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
 * @param {[]float64} v - The input vector with the dimension t.
 * @return {[]float64} - Returns a reduced dimensional vector.
 */
func (rp *RandomProjection) Project(v []float64) []float64 {
    var sum float64;
    r := make([]float64, rp.t);
    scale := math.Sqrt(3.0 / float64(rp.t));
    for i := 0; i < rp.t; i++ {
        sum = 0.0;
        for _, val := range rp.M[i] {
            sum -= v[val] * scale;
        }
        for _, val := range rp.P[i] {
            sum += v[val] * scale;
        }
        r[i] = sum;
    }
    return r;
};
