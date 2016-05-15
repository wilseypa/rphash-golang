package utils

import (
  "github.com/wilseypa/rphash-golang/types"
)

type Hash64Set struct {
  Set map[int64]bool
}

func NewHash64Set() *Hash64Set {
  return &Hash64Set{make(map[int64]bool)}
}

func (Set *Hash64Set) AddAll(other types.HashSet) {
  for k, v := range other.GetS() {
    Set.Set[k] = v
  }
}

func (Set *Hash64Set) Add(i int64) bool {
  _, found := Set.Set[i]
  Set.Set[i] = true
  return !found
}

func (Set *Hash64Set) Contains(i int64) bool {
  _, found := Set.Set[i]
  return found
}

func (Set *Hash64Set) GetS() map[int64]bool {
  return Set.Set
}

func (Set *Hash64Set) Get(i int64) bool {
  _, found := Set.Set[i]
  return found
}

func (Set *Hash64Set) Remove(i int64) {
  delete(Set.Set, i)
}

func (Set *Hash64Set) Length() int {
  return len(Set.Set)
}
