package utils;

type Hash32Set struct {
    set map[int32]bool;
};

func NewHash32Set() *Hash32Set {
    return &Hash32Set{make(map[int32]bool)};
};

func (set *Hash32Set) AddAll(other *Hash32Set) {
    for k, v := range other.set {
        set.set[k] = v;
    }
};

func (set *Hash32Set) Add(i int32) bool {
    _, found := set.set[i]
    set.set[i] = true
    return !found;
};

func (set *Hash32Set) Get(i int32) bool {
    _, found := set.set[i]
    return found;
};

func (set *Hash32Set) Remove(i int32) {
    delete(set.set, i);
};

func (set *Hash32Set) Length() int {
    return len(set.set);
};
