package projector

import (
  "math"
  "math/rand"
)

type DBFriendly struct {
  negativeVectorIndices [][]int
  positiveVectorIndices [][]int
  inputDimensionality   int
  targetDimensionality  int
  random                *rand.Rand
}

func NewDBFriendly(inputDimensionality, targetDimensionality int, randomseed int64) *DBFriendly {
  const NONZEROINDICESCHANCE = 6
  rando := rand.New(rand.NewSource(randomseed))
  negativeVectorIndices, positiveVectorIndices := make([][]int, targetDimensionality), make([][]int, targetDimensionality)
  rM, rP := 0, 0
  probability := inputDimensionality / NONZEROINDICESCHANCE
  for i := 0; i < targetDimensionality; i++ {
    orderedNegativeIndices, orderedPositiveIndices := make([]int, probability), make([]int, probability)
    for j := 0; j < inputDimensionality; j++ {
      rM, rP = rando.Intn(NONZEROINDICESCHANCE), rando.Intn(NONZEROINDICESCHANCE)
      if rM == 0 {
        orderedNegativeIndices = append(orderedNegativeIndices, int(j))
      } else if rP == 0 {
        orderedPositiveIndices = append(orderedPositiveIndices, int(j))
      }
    }
    negativeRow, positiveRow := make([]int, len(orderedNegativeIndices)), make([]int, len(orderedPositiveIndices))
    for k, val := range orderedNegativeIndices {
      negativeRow[k] = val
    }
    for k, val := range orderedPositiveIndices {
      positiveRow[k] = val
    }
    negativeVectorIndices[i], positiveVectorIndices[i] = negativeRow, positiveRow
  }

  return &DBFriendly{
    negativeVectorIndices: negativeVectorIndices,
    positiveVectorIndices: positiveVectorIndices,
    inputDimensionality:   inputDimensionality,
    targetDimensionality:  targetDimensionality,
    random:                rando,
  }
}

func (this *DBFriendly) Project(inputVector []float64) []float64 {
  var sum float64
  reducedVector := make([]float64, this.targetDimensionality)
  scale := math.Sqrt(3 / float64(this.targetDimensionality))
  for i := 0; i < this.targetDimensionality; i++ {
    sum = 0
    for _, val := range this.negativeVectorIndices[i] {
      if val >= len(inputVector) || val < 0 {
        continue
      }
      sum -= inputVector[val] * scale
    }
    for _, val := range this.positiveVectorIndices[i] {
      if val >= len(inputVector) || val < 0 {
        continue
      }
      sum += inputVector[val] * scale
    }
    reducedVector[i] = sum
  }
  return reducedVector
}
