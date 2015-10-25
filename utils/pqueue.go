package utils;

import (
    "github.com/wenkesj/rphash/types"
    "math"
);

<<<<<<< HEAD
type PQueue struct {
    plist map[int32]int32; //What is the point of this?
    heap []types.Centroid;
    heapSize int;
};

func NewPQueue(plist map[int32]int32) *PQueue {
    heapSize = 0;
    heap := [16]types.Centroid;
    return &PQueue{plist:plist};
};

func (this *PQueue) Deque() (types.Centroid, errors) {
    if(heapSize < 1) {
      return nil, errors.New("Queue contains no centriods");
    }
    result = heap[0];
    heap[0] = heap[heapSize - 1];
    this.percolateDown(0);
    heapSize--;
    return result, nil;
};

func (this *PQueue) Enque(newCentriod types.Centroid) {
  heapSize++;
  heap[heapSize] = newCentriod;
  this.percolateUp(heapSize);
};

//The children of index X are (2^X) and (2^X) + 1
func (this *PQueue) percolateUp(lowerIndex int){
  if(lowerIndex == 0) {
    return;
  }
  if(lowerIndex % 2 == 0) {
    lowerIndex = lowerIndex - 1;
  }
  var upperIndex = lowerIndex / 2;
  if(this.compare(lowerIndex, upperIndex) < 0) {
    swap(lowerIndex, upperIndex);
    if(upperIndex != 0)
      this.percolateUp(upperIndex);
    }
  }
  //Else we have fixed the priorityQueue;
}

func (this *PQueue) percolateDown(upperIndex int){
  var lowerIndex = 2 * upperIndex;
  if(lowerIndex <= this.heapSize) {
    return; // If this node has no children we are done.
  }
  if(this.compare(lowerIndex, upperIndex) < 0) {
    swap(lowerIndex, upperIndex);
    this.percolateDown(lowerIndex);
  }
  //Else we have fixed the priorityQueue;
}

func (this *PQueue) swap(index1 int, index2 int){
    var temp = heap[index1];
    heap[index1] = heap[index2];
    heap[index2] = temp;
}

func (this *PQueue) compare(centriod1 types.Centroid, centriod2 types.Centroid) int {
    id1 := centriod1.GetID(); //this.plist[centriod1.GetID()];
    id2 := centriod2.GetID()//this.plist[centriod2.GetID()];
    if id1 > id2 {
        return 1;
    } else if id1 < id2 {
        return -1;
    }
    return 0;
};
