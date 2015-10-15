package types;

type Iterator interface {
    Next() int;
    HasNext() bool;
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

type Hash interface {
    Hash(k []int32) int32;
};

type Centroid interface {
    AddID(id int32);
    Centroid() float64;
};

type ItemSet interface {
    Add(c Centroid);
    GetCounts() []int32;
    GetTop() []Centroid;
};

type LSH interface {
    LSHHashSimple(r []float64) int32;
    LSHHashStream(r []float64, n int) ([]int32, int);
    UpdateDecoderVariance(vari float64);
};

type StatTest interface {
    UpdateVarianceSample(vec []float64) float64;
    VarianceSample();
};

type RPHashObject interface {
    GetK() int;
    GetDimensions() int;
    GetRandomSeed() int64;
    GetNumberOfBlurs() int;
    GetVectorIterator() Iterator;
    GetCentroids() [][]float64;
    GetPreviousTopID() []int32;
    SetPreviousTopID(i int32);
    AddCentroid(v []float64);
    SetCentroids(l [][]float64);
    ResetDataStream();
    GetNumberOfProjections() int;
    SetNumberOfProjections(probes int);
    SetInnerDecoderMultiplier(multiDim int);
    GetInnerDecoderMultiplier() int;
    SetNumBlur(parseInt int);
    SetRandomSeed(parseLong int64);
    GetHashModulus() int32;
    SetHashModulus(parseLong int32);
    SetDecoderType(dec Decoder);
    GetDecoderType() Decoder;
    ToString() string;
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
