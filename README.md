# Scalable Big Data Clustering by Random Projection Hashing #
[![Build Status](https://travis-ci.org/wenkesj/rphash.svg)](https://travis-ci.org/wenkesj/rphash)

# Table of contents #
+ **[Installing and Testing](https://github.com/wenkesj/rphash#installing-and-testing)**
+ **[API](https://github.com/wenkesj/rphash#api)**
+ **[Developers](https://github.com/wenkesj/rphash#developers)**
+ **[Pull requests welcome](https://github.com/wenkesj/rphash/blob/master/TODO.md)**
+ **[Learn more](https://github.com/wenkesj/rphash/blob/master/REFERENCES.md)**
+ **[Versioning and updates](https://github.com/wenkesj/rphash/blob/master/CHANGELOG.md)**

# Installing and Testing #
```sh
git clone https://github.com/wenkesj/rphash
```
or

```sh
go get github.com/wenkesj/rphash
```
```sh
sh rphash/install
```

## Test ##
```sh
cd rphash/tests
go test -v -bench=.
```

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
