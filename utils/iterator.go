package utils;

// IterableSlice is a container data structure
// that supports iteration.
// That is, it satisfies IntIterator.
type IterableSlice struct {
    x int;
    s []interface{};
};

// IterableSlice.Next implements IntIterator.Next,
// satisfying the interface.
func (s *IterableSlice) Next() (value interface{}) {
    s.x++;
    return s.s[s.x];
};

func (s *IterableSlice) HasNext() (ok bool) {
    s.x++;
    if s.x >= len(s.s) {
        s.x--;
        return false;
    }
    s.x--;
    return true;
};

// NewSlice is a constructor that constructs an iterable
// container object from the native Go slice type.
func NewIterator(s []interface{}) *IterableSlice {
    return &IterableSlice{-1, s};
};
