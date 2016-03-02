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
  heap := make([]types.Centroid, 16, 16)
  return &CentriodPriorityQueue{
    heapSize: 0,
    heap:     heap,
  }
}

func (this *CentriodPriorityQueue) Dequeue() (types.Centroid, error) {
  if this.heapSize < 1 {
    err := errors.New("Queue contains no centroids")
    return nil, err
  }
  var result = this.heap[1]
  this.heap[1] = this.heap[this.heapSize]
  this.percolateDown(1)
  this.heapSize--
  return result, nil
}

func (this *CentriodPriorityQueue) Remove(idToRemove int64) bool {
  for i := 1; i < this.heapSize; i++ {
    if this.heap[i].GetID() == idToRemove {
      this.heap[i] = this.heap[this.heapSize]
      this.heap[this.heapSize] = nil
      this.heapSize--
      this.percolateDown(i)
      this.percolateUp(i)
      return true
    }
  }
  return false
}

func (this *CentriodPriorityQueue) IsEmpty() bool {
  return this.heapSize == 0
}

func (this *CentriodPriorityQueue) Poll() types.Centroid {
  var result, error = this.Dequeue()
  if error != nil {
    return nil
  }
  return result
}

func (this *CentriodPriorityQueue) Enqueue(newCentriod types.Centroid) {
  this.heapSize++
  if this.heapSize == len(this.heap) {
    var newHeap = make([]types.Centroid, len(this.heap)*2)
    copy(newHeap, this.heap)
    this.heap = newHeap
  }
  this.heap[this.heapSize] = newCentriod
  this.percolateUp(this.heapSize)
}

func (this *CentriodPriorityQueue) Size() int {
  return this.heapSize
}

//The children of index X are (2*X) and (2*X) + 1
func (this *CentriodPriorityQueue) percolateUp(lowerIndex int) {
  if lowerIndex < 2 {
    return
  }
  var upperIndex = lowerIndex / 2
  if this.compareAtPositions(lowerIndex, upperIndex) < 0 {
    this.swap(lowerIndex, upperIndex)
    this.percolateUp(upperIndex)
  }
  //Else we have fixed the priorityQueue;
  //Else we have fixed the priorityQueue;
}

func (this *CentriodPriorityQueue) percolateDown(upperIndex int) {
  var lowerIndex = 2 * upperIndex
  if lowerIndex > this.heapSize {
    return // If this node has no children we are done.
  }
  if this.compareAtPositions(lowerIndex, upperIndex) < 0 {
    this.swap(lowerIndex, upperIndex)
    this.percolateDown(lowerIndex)
    this.percolateDown(1)
  } else if lowerIndex+1 <= this.heapSize && this.compareAtPositions(lowerIndex+1, upperIndex) < 0 {
    this.swap(lowerIndex+1, upperIndex)
    this.percolateDown(lowerIndex + 1)
  }
  //Else we have fixed the priorityQueue;
}

func (this *CentriodPriorityQueue) swap(index1 int, index2 int) {
  var temp = this.heap[index1]
  this.heap[index1] = this.heap[index2]
  this.heap[index2] = temp
}

func (this *CentriodPriorityQueue) compare(centroid1 types.Centroid, centroid2 types.Centroid) int {
  count1 := centroid1.GetCount()
  count2 := centroid2.GetCount()
  if count1 > count2 {
    return 1
  } else if count1 < count2 {
    return -1
  }
  return 0
}

func (this *CentriodPriorityQueue) compareAtPositions(index1 int, index2 int) int {
  return this.compare(this.heap[index1], this.heap[index2])
}
