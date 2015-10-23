package utils;

import (
    "github.com/wenkesj/rphash/types"
);

type PQueue struct {
    plist map[int64]int64;
};

func NewPQueue(plist map[int64]int64) *PQueue {
    return &PQueue{plist:plist};
};

func (this *PQueue) Compare(x, y types.Centroid) int {
    cx := this.plist[x.GetID()];
    cy := this.plist[y.GetID()];
    if cx > cy {
        return 1;
    } else if cx < cy {
        return -1;
    }
    return 0;
};

func (this *PQueue) Remove(rem types.Centroid) (a types.Centroid) {
    return a;
};

func (this *PQueue) Add(rem types.Centroid) {

};

func (this *PQueue) Size() (a int) {
    return a;
};

func (this *PQueue) IsEmpty() (a bool) {
    return a;
};

func (this *PQueue) Poll() (a types.Centroid) {
    return a;
};
