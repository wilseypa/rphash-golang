package utils

import (
    "errors"
    // "fmt"
)
//Since we are going to remove the smallest count we need this to be a min priorityQueue

type int64WithPriority struct {
    actualInt int64
    priority int64
}
type Int64PriorityQueue struct {
    heap     []int64WithPriority
    heapSize int
}

func NewInt64PriorityQueue() *Int64PriorityQueue {
  heap := make([]int64WithPriority, 16)
  return &Int64PriorityQueue{
    heapSize: 0,
    heap: heap,
  }
}

func (this *Int64PriorityQueue) Dequeue() (int64, error) {
    if this.heapSize < 1 {
      err := errors.New("Queue contains no int64s")
      return 0, err;
    }
    var result = this.heap[1];
    this.heap[1] = this.heap[this.heapSize];
    this.heapSize--;
    this.percolateDown(1);
    return result.actualInt, nil;
}

func (this *Int64PriorityQueue) IsEmpty() bool {
  return this.heapSize == 0;
}

func (this *Int64PriorityQueue) Poll() int64 {
  var result, error = this.Dequeue();
  if error != nil {
    return 0;
  }
  // fmt.Println(this.heap);
  return result;
}

//JF there is a better way to do this. I think we might need a non heap structure
func (this *Int64PriorityQueue) Remove(toRemove int64)  bool{
  for i := 1; i <= this.heapSize; i++ {
    if this.heap[i].actualInt == toRemove {
      this.heap[i] = this.heap[this.heapSize];
      this.heapSize --;
      //We dont know if we need to percolate up or down so do both
      this.percolateUp(i);
      this.percolateDown(i);
      return true;
    }
  }
  return false;
}

func (this *Int64PriorityQueue) Enqueue(newInt int64, priority int64) {
  newIntObj := int64WithPriority{actualInt: newInt, priority: priority}
  this.heapSize++;
  if this.heapSize == len(this.heap) {
    var newHeap = make([]int64WithPriority, len(this.heap) * 2);
    copy(newHeap, this.heap);
    this.heap = newHeap;
  }
  this.heap[this.heapSize] = newIntObj;
  this.percolateUp(this.heapSize);
}

func (this *Int64PriorityQueue) Size() int {
  return this.heapSize;
}

func (this *Int64PriorityQueue) percolateUp(lowerIndex int) {
    if lowerIndex < 2 {
      return;
    }
    var upperIndex = lowerIndex / 2;
    if this.compare(lowerIndex, upperIndex) < 0 {
      this.swap(lowerIndex, upperIndex);
      this.percolateUp(upperIndex);
    }
    //Else we have fixed the priorityQueue;
}


func (this *Int64PriorityQueue) percolateDown(upperIndex int) {
    var lowerIndex = 2 * upperIndex;
    if lowerIndex > this.heapSize {
        return // If this node has no children we are done.
    }
    if this.compare(lowerIndex, upperIndex) < 0 {
        this.swap(lowerIndex, upperIndex);
        this.percolateDown(lowerIndex);
        this.percolateDown(upperIndex);
    }else if lowerIndex + 1 <= this.heapSize && this.compare(lowerIndex + 1, upperIndex) < 0 {
        this.swap(lowerIndex + 1, upperIndex);
        this.percolateDown(lowerIndex + 1);
    }
    //Else we have fixed the priorityQueue;
}

func (this *Int64PriorityQueue) swap(index1 int, index2 int) {
  var temp = this.heap[index1];
    this.heap[index1] = this.heap[index2];
  this.heap[index2] = temp;
}

func (this *Int64PriorityQueue) compare(index1 int,index2 int) int {
    if this.heap[index1].priority > this.heap[index2].priority {
        return 1;
    } else if this.heap[index1].priority < this.heap[index2].priority {
        return -1;
    }
    return 0;
}
