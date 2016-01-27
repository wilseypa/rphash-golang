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
func TestPriorityQueueEnquque(t *testing.T) {
  input := make([]int64, 5, 5);
  input[0] = 3;
  input[1] = 4;
  input[2] = 1;
  input[3] = 20;
  input[4] = 13;
  expectedResult := make([]int64, 5, 5);
  expectedResult[0] = 1;
  expectedResult[1] = 3;
  expectedResult[2] = 4;
  expectedResult[3] = 13;
  expectedResult[4] = 20;
  testQueue := utils.NewInt64PriorityQueue();
  for _, value := range input {
    testQueue.Enqueue(value);
  }
  if testQueue.Size() != len(input) {
    t.Errorf("priorityQueue is not the correct size expected length: %v, actual length: %v", len(input), testQueue.Size());
  }
  for i, expectedValue := range expectedResult {
    actualValue := testQueue.Poll();
    if actualValue != expectedValue {
      t.Errorf("priorityQueue did not output the correct value at index: %v, expected: %v, actual %v.", i, expectedValue, actualValue);
    }
  }
}
