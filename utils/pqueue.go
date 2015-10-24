package utils;

import (
    "github.com/wenkesj/rphash/types"
);

type PQueue64 struct { };

/* TODO Compare function */
// func compare(n1, n2 int64) int {
//     cn1 := countlist[HashCode(n1)];
//     cn2 := countlist[HashCode(n2)];
//     if cn1 > cn2 {
//         return +1;
//     }
//     else if cn1 < cn2 {
//         return -1;
//     }
//     return 0;
// };

func NewPQueue64() *PQueue64 {
    return &PQueue64{};
};

func (this *PQueue64) Compare(x, y int64) (a int) {
    return a;
};

func (this *PQueue64) Remove(rem int64) (a int64) {
    return a;
};

func (this *PQueue64) Add(rem int64) {

};

func (this *PQueue64) Size() (a int) {
    return a;
};

func (this *PQueue64) IsEmpty() (a bool) {
    return a;
};

func (this *PQueue64) Poll() (a int64) {
    return a;
};

type PQueueCentroid struct { };

func NewPQueueCentroid() *PQueueCentroid {
    return &PQueueCentroid{};
};

/* TODO Compare function */
// func compare(n1, n2 types.Centroid) int {
//     cn1 := countlist[n1.GetID()];
//     cn2 := countlist[n2.GetID()];
//     if cn1 > cn2 {
//         return +1;
//     }
//     else if cn1 < cn2 {
//         return -1;
//     }
//     return 0;
// };

func (this *PQueueCentroid) Compare(x, y types.Centroid) (a int) {
    return a;
};

func (this *PQueueCentroid) Remove(rem types.Centroid) (a types.Centroid) {
    return a;
};

func (this *PQueueCentroid) Add(rem types.Centroid) {

};

func (this *PQueueCentroid) Size() (a int) {
    return a;
};

func (this *PQueueCentroid) IsEmpty() (a bool) {
    return a;
};

func (this *PQueueCentroid) Poll() (a types.Centroid) {
    return a;
};
