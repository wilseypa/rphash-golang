# Database Friendly Random Projection #
## API ##

#### type _RandomProjection_ ####
```go
type RandomProjection struct {
    M [][]int;
    P [][]int;
    n int;
    t int;
    random *rand.Rand;
}
```
#### func _New_ ####
Allocate a new instance of RandomProjection.
```go
func New(n, t int, randomseed int64) *RandomProjection
```
+ `P []int` - The size [t x n/6] set of vector indices that should be positive (+sqrt(3/t) => +1).
+ `M []int` - The size [t x n/6] set of vector indices that should incur negative (-sqrt(3/t) => -1).
+ `n int` - Original dimension.
+ `t int` - Target/Projected dimension.

#### func _Project_ ####
Project onto a random matrix of {-1, 1} to produce a reduced dimensional vector.
```go
func (*RandomProjection) Project(input []float64) []float64
```
+ `v []float64` - The input vector with the dimension t.
+ `[]float64` - Returns a reduced dimensional vector.
