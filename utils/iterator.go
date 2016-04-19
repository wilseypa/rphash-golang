package utils

type IterableSlice struct {
  position int
  data     [][]float64
  lshVals  []int64
}

func (this *IterableSlice) Next() (value []float64) {
  this.position++
  return this.data[this.position]
}

func (this *IterableSlice) Size() (count int) {
  return len(this.data);
}

func (this *IterableSlice) PeakLSH() (lshValue int64) {
  if this.lshVals == nil {
    panic("Cannot call PeakLSH until after StoreLSHValues")
  }
  return this.lshVals[this.position]
}

func (this *IterableSlice) StoreLSHValues(lshVals []int64) {
  this.lshVals = lshVals
}

func (this *IterableSlice) Append(data []float64) {
  this.data = append(this.data, data)
}

func (this *IterableSlice) HasNext() (ok bool) {
  this.position++
  if this.position >= len(this.data) {
    this.position--
    return false
  }
  this.position--
  return true
}

func (this *IterableSlice) GetS() [][]float64 {
  return this.data
}

func (this *IterableSlice) Reset() {
  this.position = -1
}

func NewIterator(data [][]float64) *IterableSlice {
  return &IterableSlice{-1, data, nil}
}
