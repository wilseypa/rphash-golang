package tests;

import (
    "testing"
    "github.com/wenkesj/rphash/utils"
    "github.com/wenkesj/rphash/itemset"
);

func TestPriorityQueueEnque(t *testing.T) {
  fakeCentriod := itemset.NewCentroidSimple(1, 1);
  var countlist map[int32]int32;
  priorityQueue := utils.NewPQueue(countlist);
  if priorityQueue.Size() != 0 {
    t.Error("priorityQueue size is not 0 before Enqueing.")
  }
  priorityQueue.Enque(fakeCentriod);
  if priorityQueue.Size() != 1 {
    t.Error("priorityQueue size is not 1 after Enqueing once.")
  }
}
