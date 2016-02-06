package generator;

import (
  "math/rand"
);

type Generator struct {
  random *rand.Rand;
};

func NewGenerator(seed int64) *Generator {
  random := rand.New(rand.NewSource(seed));
  return &Generator{
    random: random,
  };
};

func (this *Generator) GenerateData(numRows int, dimensionality int) [][]float64 {
  data := make([][]float64, numRows, numRows);
  for i := 0; i < numRows; i++ {
    row := make([]float64, dimensionality, dimensionality);
    for j := 0; j < dimensionality; j++ {
      row[j] = this.random.Float64();
    }
    data[i] = row;
  }
  return data;
};
