package tests;

import (
  "github.com/wenkesj/rphash/utils"
  "github.com/wenkesj/rphash/itemset"
  "github.com/wenkesj/rphash/types"
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

  func TestCentriodQueue(t *testing.T) {
    testQueue := utils.NewCentroidPriorityQueue();
    var dimensionality = 20;
    fakeData := make([]float64, dimensionality, dimensionality);
    input := make([]types.Centroid, 5, 5);
    for i := range input {
      nextCentriod := itemset.NewCentroidSimple(dimensionality, int64(i));
      input[i] = nextCentriod;
      for j := 0; j <= i; j++ {
        input[i].UpdateCentroidVector(fakeData);
      }
      testQueue.Enqueue(input[i]);
    }
    //Remove the third centriod from the testQueue and the input
    testQueue.Remove(3);
    input = append(input[:3], input[4:]...);

    //Add another centoid to both the input and the testQueue
    nextCentriod := itemset.NewCentroidSimple(dimensionality, int64(9));
    input = append([]types.Centroid{nextCentriod}, input...);

    testQueue.Enqueue(nextCentriod);
    if testQueue.Size() != len(input) {
      t.Errorf("priorityQueue is not the correct size expected length: %v, actual length: %v", len(input), testQueue.Size());
    }
    for i, expectedValue := range input {
      actualValue := testQueue.Poll();
      if actualValue.GetID() != expectedValue.GetID() {
        t.Errorf("priorityQueue did not output the correct value at index: %v, expected: %v, actual %v.", i, expectedValue.GetID(), actualValue.GetID());
      }
    }
}
