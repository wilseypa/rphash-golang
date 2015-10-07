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
    GetK() int;
    GetDimensions() int;
    getRandomSeed() int32;
    getHashmod() int32;
    GetNumberOfBlurs() int;
    // Iterator<float[]> GetVectorIterator();
    GetCentroids() [][]float64;
    GetPreviousTopID() int32;
    SetPreviousTopID(i int32);
    AddCentroid(v []float64);
    SetCentroids(l [][]float64);
    ResetDataStream();
    GetNumberOfProjections() int;
    SetNumberOfProjections(probes int);
    SetInnerDecoderMultiplier(multiDim int);
    GetInnerDecoderMultiplier() int;
    SetNumBlur(parseInt int);
    SetRandomSeed(parseLong int32);
    GetHashModulus() int32;
    SetHashModulus(parseLong int32);
    SetDecoderType(dec Decoder);
    GetDecoderType() Decoder;
    ToString() string;
    SetVariance(data [][]float64);
};
