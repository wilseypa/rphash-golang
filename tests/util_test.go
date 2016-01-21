package tests;

import (
  "github.com/wenkesj/rphash/utils"
  "testing"
);

func TestNormalizeVector(t *testing.T) {
  input := make([]float64, 2, 2);
  input[0] = 3;
  input[1] = 4;
  expectedResult := make([]float64, 2, 2);
  expectedResult[0] = 3.0/5.0;
  expectedResult[1] = 4.0/5.0;
  result := utils.Normalize(input);
  if(len(result) != len(expectedResult)) {
    t.Errorf("Result vector has dimensionality %v, Expected Result has dimensionality %v.", len(result), len(expectedResult));
  }
  for i := 0; i < len(result); i++ {
    if(result[i] != expectedResult[i]) {
      t.Errorf("Dimension %v, equals %v in the result and %v in the expectedResult", i, result[i], expectedResult[i]);
    }
  }
}
