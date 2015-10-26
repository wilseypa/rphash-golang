package utils

import (
    "errors"
    "github.com/wenkesj/rphash/types"
)

type CentriodPriorityQueue struct {
    heap     []types.Centroid
    heapSize int
}

func NewCentroidPriorityQueue() *CentriodPriorityQueue {
    heap := make([]types.Centroid, 0, 16)
    return &CentriodPriorityQueue{
        heapSize: 0,
        heap:  heap,
    }
}

func (this *CentriodPriorityQueue) Dequeue() (types.Centroid, error) {
    if this.heapSize < 1 {
    err := errors.New("Queue contains no centroids")
        return nil, err;
    }
    var result = this.heap[0];
    this.heap[0] = this.heap[this.heapSize];
    this.percolateDown(0);
    this.heapSize--;
    return result, nil;
}

func (this *CentriodPriorityQueue) IsEmpty() bool {
  return this.heapSize == 0;
}

func (this *CentriodPriorityQueue) Poll() (types.Centroid) {
  var result, error = this.Dequeue();
  if error == nil {
    return nil;
  }
  return result;
}

func (this *CentriodPriorityQueue) Enqueue(newCentriod types.Centroid) {
    this.heapSize++;
    this.heap[this.heapSize] = newCentriod;
    this.percolateUp(this.heapSize);
}

func (this *CentriodPriorityQueue) Size() int {
  return this.heapSize;
}

//The children of index X are (2*X) and (2*X) + 1
func (this *CentriodPriorityQueue) percolateUp(lowerIndex int) {
    if lowerIndex == 0 {
        return;
    }
    if lowerIndex%2 == 0 {
        lowerIndex = lowerIndex - 1;
    }
    var upperIndex = lowerIndex / 2;
    if this.compareAtPositions(lowerIndex, upperIndex) < 0 {
        this.swap(lowerIndex, upperIndex);
        if upperIndex != 0 {
            this.percolateUp(upperIndex);
        }
    }
    //Else we have fixed the priorityQueue;
}

func (this *CentriodPriorityQueue) percolateDown(upperIndex int) {
    var lowerIndex = 2 * upperIndex;
    if lowerIndex <= this.heapSize {
        return // If this node has no children we are done.
    }
    if this.compareAtPositions(lowerIndex, upperIndex) < 0 {
        this.swap(lowerIndex, upperIndex);
        this.percolateDown(lowerIndex);
    }
    //Else we have fixed the priorityQueue;
}

func (this *CentriodPriorityQueue) swap(index1 int, index2 int) {
  var temp = this.heap[index1];
    this.heap[index1] = this.heap[index2];
  this.heap[index2] = temp;
}

func (this *CentriodPriorityQueue) compare(centroid1 types.Centroid, centroid2 types.Centroid) int {
  id1 := centroid1.GetID();
    id2 := centroid2.GetID();
    if id1 > id2 {
        return 1;
    } else if id1 < id2 {
        return -1;
    }
    return 0;
}

func (this *CentriodPriorityQueue) compareAtPositions(index1 int, index2 int) int {
  return this.compare(this.heap[index1], this.heap[index2]);
}
