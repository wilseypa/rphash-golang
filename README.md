# Scalable Big Data Clustering by Random Projection Hashing #
[![Build Status](https://travis-ci.org/wenkesj/rphash.svg)](https://travis-ci.org/wenkesj/rphash)
[![Release Status](https://img.shields.io/badge/version-1.0.0-blue.svg)](https://github.com/wenkesj/rphash/releases)

**Clustering** is a core concept in data analysis. Issues arise with scalability and dimensionality, ever changing environments and compatibility, insecure communications and data movement

**The solution** is secure, reliable, and fast data for large- scale distributed systems.


**The Algorithm** provides more accurate clusters and an inherently distributed system.

![Clusters](https://github.com/wenkesj/rphash/blob/master/clusters.png)

**Random Projection Hash (RPHash)** has been created for maximizing parallel computation
while providing scalability for large scale deployment. It's suitable for high dimensional data sets and is
scalable and streamline.

![Overview](https://github.com/wenkesj/rphash/blob/master/overview.png)

# Table of contents #
+ **[Installing and Testing](https://github.com/wenkesj/rphash#installing-and-testing)**
+ **[API](https://github.com/wenkesj/rphash#api)**
+ **[Developers](https://github.com/wenkesj/rphash#developers)**
+ **[Pull requests welcome](https://github.com/wenkesj/rphash/blob/master/TODO.md)**
+ **[Learn more](https://github.com/wenkesj/rphash/blob/master/REFERENCES.md)**
+ **[Versioning and updates](https://github.com/wenkesj/rphash/blob/master/CHANGELOG.md)**

# Installing, Testing, and Plotting #
```sh
git clone --depth=50 --branch=master https://github.com/wenkesj/rphash.git wenkesj/rphash
cd wenkesj/rphash
export GOPATH=$HOME/<your-gopath>
export PATH=$HOME/<your-gopath>/bin:$PATH
go get -t -v ./...
sh install
```

## Test ##
```sh
go test ./tests -v -bench=.
```

## Plot ##
If you wish to have this functionality you must run 
```sh
go get github.com/gonum/plot
```
Plot tests. **[option]** is the name of the file/test plot.
```sh
sh rphash/plot [option]
```

For example, `sh rphash/plot kmeans`, will run rphash/plots/plot_kmeans.go.

# API #
The official documentation for the high performance big data clustering algorithm **RPHash**.

+ **[type RPHashObject](https://github.com/wenkesj/rphash#type-rphashobject)**
  + **[func NewSimpleArray](https://github.com/wenkesj/rphash#func-newsimplearray)**
+ **[type Simple](https://github.com/wenkesj/rphash#simple)**
  + **[func NewSimple](https://github.com/wenkesj/rphash#func-newsimple)**
  + **[func (\*Simple) Map](https://github.com/wenkesj/rphash)**
  + **[func (\*Simple) Reduce](https://github.com/wenkesj/rphash)**
  + **[func (\*Simple) GetCentroids](https://github.com/wenkesj/rphash)**
  + **[func (\*Simple) Run](https://github.com/wenkesj/rphash)**
  + **[func (\*Simple) GetParam](https://github.com/wenkesj/rphash)**
+ **[type Stream](https://github.com/wenkesj/rphash#stream)**
  + **[func NewStream](https://github.com/wenkesj/rphash#newstream)**
  + **[func (\*Stream) AddVectorOnlineStep](https://github.com/wenkesj/rphash)**
  + **[func (\*Stream) GetCentroids](https://github.com/wenkesj/rphash)**
  + **[func (\*Stream) GetCentroidsOfflineStep](https://github.com/wenkesj/rphash)**
  + **[func (\*Stream) Run](https://github.com/wenkesj/rphash)**

## type RPHashObject ##
An instance of the RPHashObject is the SimpleArray struct.

```go
import "github.com/wenkesj/rphash/reader/simplearray"
```

```go
type SimpleArray struct {
  data types.Iterator;
  dimension int;
  numberOfProjections int;
  decoderMultiplier int;
  randomSeed int64;
  hashModulus int64;
  k int;
  numberOfBlurs int;
  decoder types.Decoder;
  centroids [][]float64;
  topIDs []int64;
};
```

### func NewSimpleArray ###
```go
func NewSimpleArray(X [][]float64, k int) *SimpleArray
```

Returns a new RPHashObject.

## type Simple ##
```go
import "github.com/wenkesj/rphash/simple"
```

```go
type Simple struct {
  centroids [][]float64
  variance float64
  rphashObject RPHashObject
}
```

### func NewSimple ###
```go
func NewSimple(_rphashObject RPHashObject) *Simple
```

NewSimple returns an instance of the Simple struct.

### func (\*Simple) Map ###
```go
func (this *Simple) Map() RPHashObject
```

Maps all the default tasks to the RPHashObject. This will update and return the new RPHashObject.

### func (\*Simple) Reduce ###
```go
func (this *Simple) Reduce() RPHashObject
```

Performs all the default tasks on the RPHashObject. Updates and returns new RPHashObject.

### func (\*Simple) GetCentroids ###
```go
func (this *Simple) GetCentroids() [][]float64
```

Performs a **KMeans** operation on the **Simple's** centroids with the RPHashObject **K** value. Returns the calculated centroids.

### func (\*Simple) Run ###
```go
func (this *Simple) Run()
```

Performs the **Map** and **Reduce** functions and updates the centroids.

### func (\*Simple) GetParam ###
```go
func (this *Simple) GetParam() RPHashObject
```

Returns the RPHashObject of the Simple struct.

## type Stream ##
```go
import "github.com/wenkesj/rphash/stream"
```

```go
type Stream struct {
  counts []int64;
  centroids [][]float64;
  variance float64;
  centroidCounter types.CentroidItemSet;
  random *rand.Rand;
  rphashObject types.RPHashObject;
  lshGroup []types.LSH;
  decoder types.Decoder;
  projector types.Projector;
  hash types.Hash;
  varTracker types.StatTest;
};

```
### func NewStream ###
```go
func NewStream(_rphashObject types.RPHashObject) *Stream
```

### func (\*Stream) AddVectorOnlineStep ###
```go
func (this *Stream) AddVectorOnlineStep(vec []float64) int64
```

### func (\*Stream) GetCentroids ###
```go
func (this *Stream) GetCentroids() [][]float64
```

### func (\*Stream) GetCentroidsOfflineStep ###
```go
func (this *Stream) GetCentroidsOfflineStep() [][]float64
```

### func (\*Stream) Run ###
```go
func (this *Stream) Run()
```

# Developers #
+ Sam Wenke (**wenkesj**)
+ Jacob Franklin (**frankljbe**)
+ Sadiq Quasem (**quasemsm**)
