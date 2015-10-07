/**
 * Random Projection
 * 1st Step
 * @author Sam Wenke
 * @author Jacob Franklin
 */
package rp

import (
	"math"
	"math/rand"
)

type RandomProjection struct {
	negativeVectorIndices [][]int
	positiveVectorIndices [][]int
	n                     int
	targetDimensionality  int
	random                *rand.Rand
}

/**
 * Allocate a new instance of RandomProjection.
 * @param {int} n - Original dimension.
 * @param {int} targetDimensionality - Target/Projected dimension.
 * @param {int} randomseed - Random seed.
 */
func New(n, targetDimensionality int, randomseed int64) *RandomProjection {
	rando := rand.New(rand.NewSource(randomseed))
	negativeVectorIndices, positiveVectorIndices := make([][]int, targetDimensionality), make([][]int, targetDimensionality)
	rM, rP := 0, 0
	probability := n / 6
	for i := 0; i < targetDimensionality; i++ {
		orderedM, orderedP := make([]int, probability), make([]int, probability)
		for j := 0; j < n; j++ {
			rM, rP = rando.Intn(6), rando.Intn(6)
			if rM == 0 {
				orderedM = append(orderedM, int(j))
			} else if rP == 0 {
				orderedP = append(orderedP, int(j))
			}
		}
		tmpM, tmpP := make([]int, len(orderedM)), make([]int, len(orderedP))
		for k, val := range orderedM {
			tmpM[k] = val
		}
		for k, val := range orderedP {
			tmpP[k] = val
		}
		negativeVectorIndices[i], positiveVectorIndices[i] = tmpM, tmpP
	}

	return &RandomProjection{
		negativeVectorIndices: negativeVectorIndices,
		positiveVectorIndices: positiveVectorIndices,
		n:                    n,
		targetDimensionality: targetDimensionality,
		random:               rando,
	}
}

/**
 * Project onto a random matrix of {-1, 1} to produce a reduced dimensional vector.
 * @return {[]float64} reducedVector - Returns a reduced dimensional vector with dimension t.
 */
func (this *RandomProjection) Project(inputVector []float64) []float64 {
	var sum float64
	reducedVector := make([]float64, this.targetDimensionality)
	scale := math.Sqrt(3 / float64(this.targetDimensionality))
	for i := 0; i < this.targetDimensionality; i++ {
		sum = 0
		for _, val := range this.negativeVectorIndices[i] {
			sum -= inputVector[val] * scale
		}
		for _, val := range this.positiveVectorIndices[i] {
			sum += inputVector[val] * scale
		}
		reducedVector[i] = sum
	}
	return reducedVector
}
