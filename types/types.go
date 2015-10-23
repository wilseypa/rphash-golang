package types;

type Iterator interface {
    GetS() [][]float64;
    Next() (value []float64);
    HasNext() (ok bool);
};

type PQueue interface {
    IsEmpty() bool;
    Poll(i interface{});
    Push(i interface{});
    Pop() interface{};
    Length() int;
    Less(i, j int) bool;
    Swap(i, j int);
    String() string;
};

type Decoder interface {
    SetVariance(parameterObject float64);
    GetDimensionality() int;
    Decode(f []float64) []int32;
    GetErrorRadius() float64;
    GetDistance() float64;
    GetVariance() float64;
};

type Projector interface {
    Project(v []float64) []float64;
};

type HashSet interface {
    Add(i int32) bool;
    Get(i int32) bool;
    AddAll(i HashSet);
    GetS() map[int32]bool;
    Remove(i int32);
    Length() int;
    Contains(i int32) bool;
};

type Hash interface {
    Hash(k []int32) int32;
};

type Centroid interface {
    UpdateCentroidVector(data []float64);
    Centroid() []float64;
    UpdateVector(rp []float64);
    GetCount() int32;
    GetID() int32;
    GetIDs() HashSet;
    AddID(h int32);
};

type CountItemSet interface {
    Add(c int32);
    GetCounts() []int32;
    GetTop() []int32;
    GetCount() int32;
};

type CentroidItemSet interface {
    Add(c Centroid);
    GetCounts() []int32;
    GetTop() []Centroid;
    GetCount() int32;
};

type LSH interface {
    LSHHashSimple(r []float64) int32;
    LSHHashStream(r []float64, a int) []int32;
    UpdateDecoderVariance(vari float64);
};

type StatTest interface {
    UpdateVarianceSample(vec []float64) float64;
};

type RPHashObject interface {
    GetK() int;
    GetDimensions() int;
    GetRandomSeed() int64;
    GetNumberOfBlurs() int;
    GetVectorIterator() Iterator;
    GetCentroids() [][]float64;
    GetPreviousTopID() []int32;
    SetPreviousTopID(i []int32);
    AddCentroid(v []float64);
    SetCentroids(l [][]float64);
    GetNumberOfProjections() int;
    SetNumberOfProjections(probes int);
    SetInnerDecoderMultiplier(multiDim int);
    GetInnerDecoderMultiplier() int;
    SetRandomSeed(parseLong int64);
    GetHashModulus() int32;
    SetHashModulus(parseLong int32);
    SetDecoderType(dec Decoder);
    GetDecoderType() Decoder;
    SetVariance(data [][]float64);
};

type Clusterer interface {
    GetCentroids() [][]float64;
    GetParam() RPHashObject;
};

type StreamClusterer interface {
    AddVectorOnlineStep(x []float64) int32;
    GetCentroidsOfflineStep() [][]float64;
};
