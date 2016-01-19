# Scalable Big Data Clustering by Random Projection Hashing #
[![Build Status](https://travis-ci.org/wilseypa/rphash-golang.svg)](https://travis-ci.org/wilseypa/rphash-golang)
[![Release Status](https://img.shields.io/badge/version-1.0.0-blue.svg)](https://github.com/wilseypa/rphash-golang/releases)

**Clustering** is a core concept in data analysis. Issues arise with scalability and dimensionality, ever changing environments and compatibility, insecure communications and data movement

**The solution** is secure, reliable, and fast data for large- scale distributed systems.


**The Algorithm** provides more accurate clusters and an inherently distributed system.

![Clusters](https://github.com/wilseypa/rphash-golang/blob/master/clusters.png)

**Random Projection Hash (RPHash)** has been created for maximizing parallel computation
while providing scalability for large scale deployment. It's suitable for high dimensional data sets and is
scalable and streamline.

![Overview](https://github.com/wilseypa/rphash-golang/blob/master/overview.png)

# Table of contents #
+ **[Installing and Testing](https://github.com/wilseypa/rphash-golang#installing-and-testing)**
+ **[API](https://github.com/wilseypa/rphash-golang#api)**
+ **[Developers](https://github.com/wilseypa/rphash-golang#developers)**
+ **[Pull requests welcome](https://github.com/wilseypa/rphash-golang/blob/master/TODO.md)**
+ **[Learn more](https://github.com/wilseypa/rphash-golang/blob/master/REFERENCES.md)**
+ **[Versioning and updates](https://github.com/wilseypa/rphash-golang/blob/master/CHANGELOG.md)**

# Installing, Testing, and Plotting #
```sh
git clone --depth=50 --branch=master https://github.com/wilseypa/rphash-golang.git wilseypa/rphash-golang
cd wilseypa/rphash-golang
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
sh rphash-golang/plot [option]
```

For example, `sh rphash-golang/plot kmeans`, will run rphash-golang/plots/plot_kmeans.go.

# API #
The official documentation for the high performance big data clustering algorithm **RPHash**.

+ **[type RPHashObject](https://github.com/wilseypa/rphash-golang#type-rphashobject)**
  + **[func NewSimpleArray](https://github.com/wilseypa/rphash-golang#func-newsimplearray)**
+ **[type Simple](https://github.com/wilseypa/rphash-golang#simple)**
  + **[func NewSimple](https://github.com/wilseypa/rphash-golang#func-newsimple)**
  + **[func (\*Simple) Map](https://github.com/wilseypa/rphash-golang)**
  + **[func (\*Simple) Reduce](https://github.com/wilseypa/rphash-golang)**
  + **[func (\*Simple) GetCentroids](https://github.com/wilseypa/rphash-golang)**
  + **[func (\*Simple) Run](https://github.com/wilseypa/rphash-golang)**
  + **[func (\*Simple) GetParam](https://github.com/wilseypa/rphash-golang)**
+ **[type Stream](https://github.com/wilseypa/rphash-golang#stream)**
  + **[func NewStream](https://github.com/wilseypa/rphash-golang#newstream)**
  + **[func (\*Stream) AddVectorOnlineStep](https://github.com/wilseypa/rphash-golang)**
  + **[func (\*Stream) GetCentroids](https://github.com/wilseypa/rphash-golang)**
  + **[func (\*Stream) GetCentroidsOfflineStep](https://github.com/wilseypa/rphash-golang)**
  + **[func (\*Stream) Run](https://github.com/wilseypa/rphash-golang)**

## type RPHashObject ##
An instance of the RPHashObject is the SimpleArray struct.

```go
import "github.com/wilseypa/rphash-golang/reader/simplearray"
```

```go
type SimpleArray struct {
  data types.Iterator;
  dimension int;
  numberOfProjections int;
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
import "github.com/wilseypa/rphash-golang/simple"
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
import "github.com/wilseypa/rphash-golang/stream"
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
