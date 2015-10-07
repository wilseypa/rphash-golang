package types;

type Decoder interface {
    New() Decoder;
    SetVariance(parameterObject float64);
    GetDimensionality() int;
    Decode(f []float64) []int32;
    GetErrorRadius() float64;
    GetDistance() float64;
};

type Projector interface {
    New(n, t int, randomseed int64) Projector;
    Project(v []float64) []float64;
};

type Hash interface {
    New(hashMod int32) Hash;
    Hash(k []int32) int32;
};

type CentroidCounter interface {
    New(k int) CentroidCounter;
};

type LSH interface {
    New(hash Hash,
        decoder Decoder,
        projector Projector) LSH;
    MinHash(r []float64, radius float64, randomseed int64, n int) ([]int32, int);
};

type RPHashObject interface {
    New() RPHashObject;
    GetDimensions() int;
    GetHashModulus() int32;
    GetNumberOfProjections() int;
};
