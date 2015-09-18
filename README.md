# Scalable Big Data Clustering by Random Projection Hashing #
+ Sam Wenke, Jacob Franklin, Sadiq Quasem
+ Advised by Dr. Phillip A. Wilsey

## Description ##
Clustering aims to solve a significant problem; finding structure in a collection of unlabeled and unconstrained streams of data. Scalable clusters of large sets of data are widely applied in today’s cumulonimbus of networking. The importance of cloud computing is tangible, and is growing at an exponential rate. The answer to clustering can only be approached by the most carefully optimized and secured algorithm. We are producing an open source streaming clustering algorithm with log-linear complexity growth and intrinsic data security geared to ease scalability and security of large scale data applications.

Parallel algorithms and modern methods will be used. The aforementioned includes developing the source in the Go<sup>1</sup> language, deploying on Hadoop<sup>2</sup> containers, and producing an unrivaled open source technology.

## Goal ##
The largest online storage and service companies (consider Google, Amazon, Microsoft and Facebook) are estimated to store at least 1,200 petabytes of data between them. Security and efficiency can be overlooked as growing pains in modern fast paced application architectures. Our team is focused on providing a Scalable Big Data Clustering open source project to accommodate large scale Medical applications that are HIPAA<sup>3</sup> compliant and transfer sensitive patient information to secure databases. As medical technologies grow and patient care becomes more and more dependent on secure, reliable, and fast data, this project will provide for all of these.

<sub><sup>1</sup>Go is a programming language developed by Google Inc., it aims to provide memory safety, garbage collection, and light weight threading for parallel computing.</sub>

<sub><sup>2</sup>Hadoop allows for the distributed processing of large data sets across clusters of computers.</sub>

<sub><sup>3</sup>HIPAA Health Insurance Portability and Accountability Act</sub>

## API ##

### FJLT Projection ###

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
Allocate a new instance of the FJLTProjection.
+ Fast Johnson Lindenstrauss Transform under the Walsh-Hadamard vector matrix transform. Random embedding &phi; ~ FJLT(n, d, &epsilon;, p) is the product of P * H * D. Were &epsilon; is calculated.

``` go
func New(d, k, n int64) *FJLTProjection
```

Multiplies a matrix by a vector (single precision).

```go
func SGEMV(t, n, startpoint, startoutput int64, M, v, result []float64, alpha float64)
```

Generate a k-by-d matrix whose elements are
independently distributed as follows. With probabilty
1 - q set P as 0, and otherwise (with the remaining probabilty q)
draw P from a normal distribution of expectation 0 and variance
q^-1.

```go
func GenerateP(n, k, d, p int64, e float64, random *rand.Rand) []float64
```

Generate a d-by-d diagonal matrix where D is
drawn independently from {-1,1} with probability 1/2

```go
func GenerateD(d int64, random *rand.Rand) []float64
```
Normal distribution.
+ Takes an input pointer to hold the distribution data
Size of Distribution (m,n).
+ Outputs a matrix filled with normal distribution.
+ Uses Moro's Inverse CND distribution to
generate an arbitrary normal distribution with
mean mu and variance vari.

```go
func InvRandN(data []float64, m, n int64, mu, vari float64, random *rand.Rand)
```
Takes an input pointer to hold the distribution data
with the size of Distribution (m,n).
Outputs a matrix filled with uniform distribution.
```go
func RandU(data []float64, m, n int64, random *rand.Rand)
```
Moro's inverse Cumulative Normal Distribution
function approximation.
```go
func MoroInvCND(P float64) float64
```
Performs the FJLT on a matrix.
```go
func (_fjlt *FJLTProjection) FJLT(input []float64) []float64
```
Project a matrix.
```go
func (_fjlt *FJLTProjection) Project(input []float64) []float64
```
