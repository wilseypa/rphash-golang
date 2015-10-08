package types;

type DecoderContructor interface {
    New() Decoder;
};

type ProjectorContructor interface {
    New(n, t int, randomseed int64) Projector;
};

type HashContructor interface {
    New(hashMod int32) Hash;
};

type KMeansContructor interface {
    New() KMeans;
}

type CentroidContructor interface {
    New(vec []float64) Centroid;
};

type CentroidCounterContructor interface {
    New(k int) CentroidCounter;
};

type LSHContructor interface {
    New(hash Hash, decoder Decoder, projector Projector) LSH;
};

type StatTestConstructor interface {
    New(vari float64) StatTest;
};

type StreamObjectConstructor interface {
    New() StreamObject;
};

type Decoder interface {
    SetVariance(parameterObject float64);
    GetDimensionality() int;
    Decode(f []float64) []int32;
    GetErrorRadius() float64;
    GetDistance() float64;
};

type Projector interface {
    Project(v []float64) []float64;
};

type Hash interface {
    Hash(k []int32) int32;
};

type KMeans interface {
    GetCentroids();
};

type Centroid interface {
    AddId(id int32);
};

type CentroidCounter interface {
    Add(c Centroid);
    GetCounts();
    GetTop();
};

type LSH interface {
    MinHash(r []float64, radius float64, randomseed int64, n int) ([]int32, int);
};

type StatTest interface {
    UpdateVarianceSample(vec []float64) float64;
    VarianceSample();
};

type StreamObject interface {
    GetK() int;
    GetDimensions() int;
    getRandomSeed() int32;
    getHashmod() int32;
    GetNumberOfBlurs() int;
    GetVectorIterator() []float64;
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
