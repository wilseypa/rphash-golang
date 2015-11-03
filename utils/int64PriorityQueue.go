package utils

import (
    "errors"
)

type Int64PriorityQueue struct {
    heap     []int64
    heapSize int
}

func NewInt64PriorityQueue() *Int64PriorityQueue {
    heap := make([]int64, 0, 16)
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
    this.percolateDown(1);
    this.heapSize--;
    return result, nil;
}

func (this *Int64PriorityQueue) IsEmpty() bool {
  return this.heapSize == 0;
}

func (this *Int64PriorityQueue) Poll() int64 {
  var result, error = this.Dequeue();
  if error == nil {
    return 0;
  }
  return result;
}

//JF there is a better way to do this. I think we might need a non heap structure
func (this *Int64PriorityQueue) Remove(toRemove int64)  bool{
  for i := 1; i <= this.heapSize; i++ {
    if this.heap[i] == toRemove {
      this.heap[i] = this.heap[this.heapSize];
      this.heapSize --;
      //We wdont know if we need to percolate up or down so do both
      this.percolateUp(i);
      this.percolateDown(i);
      this.heapSize --;
      return true;
    }
  }
  return false;
}

func (this *Int64PriorityQueue) Enqueue(newInt int64) {
    this.heapSize++;
    this.heap[this.heapSize] = newInt;
    this.percolateUp(this.heapSize);
}

func (this *Int64PriorityQueue) Size() int {
  return this.heapSize;
}

//The children of index X are (2*X) and (2*X) + 1
func (this *Int64PriorityQueue) percolateUp(lowerIndex int) {
    if lowerIndex == 0 {
        return;
    }
    if lowerIndex%2 == 0 {
        lowerIndex = lowerIndex - 1;
    }
    var upperIndex = lowerIndex / 2;
    if this.compare(lowerIndex, upperIndex) < 0 {
        this.swap(lowerIndex, upperIndex);
        if upperIndex != 0 {
            this.percolateUp(upperIndex);
        }
    }
    //Else we have fixed the priorityQueue;
}

func (this *Int64PriorityQueue) percolateDown(upperIndex int) {
    var lowerIndex = 2 * upperIndex;
    if lowerIndex <= this.heapSize {
        return // If this node has no children we are done.
    }
    if this.compare(lowerIndex, upperIndex) < 0 {
        this.swap(lowerIndex, upperIndex);
        this.percolateDown(lowerIndex);
    }
    //Else we have fixed the priorityQueue;
}

func (this *Int64PriorityQueue) swap(index1 int, index2 int) {
  var temp = this.heap[index1];
    this.heap[index1] = this.heap[index2];
  this.heap[index2] = temp;
}

func (this *Int64PriorityQueue) compare(index1 int,index2 int) int {
    if this.heap[index1] > this.heap[index1] {
        return 1;
    } else if this.heap[index1] < this.heap[index2] {
        return -1;
    }
    return 0;
}
