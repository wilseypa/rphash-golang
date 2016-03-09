package utils

import (
  "github.com/wilseypa/rphash-golang/types"
)

type Hash64Set struct {
  set map[int64]bool
}

func NewHash64Set() *Hash64Set {
  return &Hash64Set{make(map[int64]bool)}
}

func (set *Hash64Set) AddAll(other types.HashSet) {
  for k, v := range other.GetS() {
    set.set[k] = v
  }
}

func (set *Hash64Set) Add(i int64) bool {
  _, found := set.set[i]
  set.set[i] = true
  return !found
}

func (set *Hash64Set) Contains(i int64) bool {
  _, found := set.set[i]
  return found
}

func (set *Hash64Set) GetS() map[int64]bool {
  return set.set
}

func (set *Hash64Set) Get(i int64) bool {
  _, found := set.set[i]
  return found
}

func (set *Hash64Set) Remove(i int64) {
  delete(set.set, i)
}

func (set *Hash64Set) Length() int {
  return len(set.set)
}
