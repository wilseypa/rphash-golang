package utils;

import (
    "github.com/wenkesj/rphash/types"
);

type Hash32Set struct {
    set map[int64]bool;
};

func NewHash32Set() *Hash32Set {
    return &Hash32Set{make(map[int64]bool)};
};

func (set *Hash32Set) AddAll(other types.HashSet) {
    for k, v := range other.GetS() {
        set.set[k] = v;
    }
};

func (set *Hash32Set) Add(i int64) bool {
    _, found := set.set[i]
    set.set[i] = true
    return !found;
};

func (set *Hash32Set) Contains(i int64) (a bool) {
    return a;
};

func (set *Hash32Set) GetS() map[int64]bool {
    return set.set;
};

func (set *Hash32Set) Get(i int64) bool {
    _, found := set.set[i]
    return found;
};

func (set *Hash32Set) Remove(i int64) {
    delete(set.set, i);
};

func (set *Hash32Set) Length() int {
    return len(set.set);
};
