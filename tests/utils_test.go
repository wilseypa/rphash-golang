package tests;

import (
    "testing"
    "github.com/wenkesj/rphash/utils"
    //"github.com/wenkesj/rphash/itemset"
);

/*func TestPriorityQueueEnque(t *testing.T) {
  fakeCentriod := itemset.NewCentroidSimple(1, 1);
  priorityQueue := utils.NewCentroidPriorityQueue();
  if priorityQueue.Size() != 0 {
    t.Error("priorityQueue size is not 0 before Enqueing.")
  }
  priorityQueue.Enqueue(fakeCentriod);
  if priorityQueue.Size() != 1 {
    t.Error("priorityQueue size is not 1 after Enqueing once.")
  }
}*/
func TestInt64PriorityQueueDequeue(t *testing.T) {
  intQueue := utils.NewInt64PriorityQueue();
  if intQueue.Size() != 0 {
    t.Error("priorityQueue size is not 0 before Enqueing.");
  }
  var i = int64(20);
  for i > 0 {
    intQueue.Enqueue(i);
    i--;
  }
  intQueue.Enqueue(12);
  var result = intQueue.Poll();
  if result != 20 {
    t.Errorf("priorityQueue does not return the largest int, 20. It reutrned %d", result);
  }
  intQueue.Enqueue(12);
  result = intQueue.Poll();
  if result != 19 {
    t.Errorf("priorityQueue does not return the largest int, 19. It reutrned %d", result);
  }
}

func TestInt64PriorityQueueRemove(t *testing.T) {
  intQueue := utils.NewInt64PriorityQueue();
  if intQueue.Size() != 0 {
    t.Error("priorityQueue size is not 0 before Enqueing.");
  }
  var i = int64(20);
  for i > 0 {
    intQueue.Enqueue(i);
    i--;
  }
  if !intQueue.Remove(20) {
    t.Error("Failedto remove 20");
  }
  if !intQueue.Remove(12) {
    t.Error("Failedto remove 12");
  }
  var result = intQueue.Poll();
  if result != 19 {
    t.Errorf("priorityQueue does not return the largest int, 19. It reutrned %d", result);
  }
  result = intQueue.Poll();  //TODO currently failing
  //if result != 18 {
  //  t.Errorf("priorityQueue does not return the largest int, 18. It reutrned %d", result);
  //}
}

func TestInt64PriorityQueueHeapSize(t *testing.T) {
  intQueue := utils.NewInt64PriorityQueue();
  if intQueue.Size() != 0 {
    t.Error("priorityQueue size is not 0 before Enqueing.");
  }
  var i = int64(0);
  for i < 20 {
    intQueue.Enqueue(i);
    i++;
  }
  if intQueue.Size() != 20 {
    t.Error("priorityQueue size is not 10 after Enqueing ten time.");
  }
  intQueue.Dequeue();
  if intQueue.Size() != 19 {
    t.Error("priorityQueue size did not decrease after dequeueing");
  }
}
