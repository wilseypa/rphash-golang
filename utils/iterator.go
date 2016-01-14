package utils;

import (
  "fmt"
);

type IterableSlice struct {
    x int;
    s [][]float64;
};

func (s *IterableSlice) Next() (value []float64) {
    s.x++;
    return s.s[s.x];
};

func (s *IterableSlice) HasNext() (ok bool) {
    fmt.Println(s.x);
    s.x++;
    if s.x >= len(s.s) {
        s.x--;
        return false;
    }
    s.x--;
    return true;
};

func (s *IterableSlice) GetS() [][]float64 {
    return s.s;
};

func (s *IterableSlice) Reset() {
  s.x = -1;
}

func NewIterator(s [][]float64) *IterableSlice {
    return &IterableSlice{-1, s};
};
