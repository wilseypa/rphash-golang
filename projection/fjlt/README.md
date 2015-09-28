# fast-johnson-lindenstrauss-transform #
## API ##

#### type _FJLTProjection_ ####
```go
type FJLTProjection struct {
    n int;
    k int;
    d int;
    D []float64;
    P []float64;
    random *rand.Rand;
}
```
Fast Johnson Lindenstrauss Transform random embedding &phi; ~ FJLT(n, d, &epsilon;, p) is the product of P x H x D
#### func _New_ ####
``` go
func New(d, k, n int64) *FJLTProjection
```
Allocate a new instance of the FJLTProjection.
+ `d int64` Number of rows of D, Number of columns of P
+ `k int64` Number of rows of P, Number of columns of D
+ `n int64` Target dimension

#### func _FJLT_ ####
```go
func (*FJLTProjection) FJLT(input []float64) []float64
```
Performs the FJLT on a matrix
#### func _Project_ ####
```go
func (*FJLTProjection) Project(input []float64) []float64
```
Project a matrix
