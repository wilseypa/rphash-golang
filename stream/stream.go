package stream;

import (
    "math/rand"
    "github.com/wenkesj/rphash/utils/constructor"
    "github.com/wenkesj/rphash/utils/vectors"
    "github.com/wenkesj/rphash/types"
);

type Stream struct {
    counts []int32;
    centroids []float64;
    variance float64;
    centroidCounter types.CentroidCounter;
    random *rand.Rand;
    rphashStream types.StreamObject;
    lshGroup []types.LSH;
    decoder types.Decoder;
    projector types.Projector;
    hash types.Hash;
    varTracker types.StatTest;
};

var construct *constructor.Constructor;

func New(data []float64,
            centroidCounter types.CentroidCounterConstructor,
            rphashStream types.StreamObjectConstructor,
            lsh types.LSHConstructor,
            kmeans types.KMeansConstructor,
            centroid types.CentroidConstructor,
            decoder types.DecoderConstructor,
            projector types.ProjectorConstructor,
            hash types.HashConstructor,
            statTest types.StatTestConstructor) *Stream {

    construct = constructor.New(decoder, projector, hash, kmeans, centroid, centroidCounter, lsh, statTest, rphashStream);
    _rphashStream := construct.NewStreamObject();
    _random := rand.New(rand.NewSource(_rphashStream.GetRandomSeed()));
    _hash := construct.NewHash(_rphashStream.GetHashModulus());
    _decoder := construct.NewDecoder();
    _statTest := construct.NewStatTest(0.01);
    projections := _rphashStream.GetNumberOfProjections();
    k := _rphashStream.GetK() * projections;
    _centroidCounter := construct.NewCentroidCounter(k);
    _lshGroup := make([]types.LSH, projections);
    var _projector types.Projector;
    for i := 0; i < projections; i++ {
        _projector = construct.NewProjector(_rphashStream.GetDimensions(), _decoder.GetDimensionality(), _random.Int63());
        _lshGroup[i] = construct.NewLSH(_hash, _decoder, _projector);
    }
    return &Stream{
        counts: nil,
        centroids: nil,
        variance: 0,
        centroidCounter: _centroidCounter,
        random: _random,
        rphashStream: _rphashStream,
        lshGroup: _lshGroup,
        hash: _hash,
        decoder: _decoder,
        projector: _projector,
        varTracker: _statTest,
    };
};

func (this *Stream) AddVectorOnlineStep(vec []float64) int32 {
    var hash []int32;
    c := construct.NewCentroid(vec);

    tmpvar := this.varTracker.UpdateVarianceSample(vec);
    if this.variance != tmpvar {
        for _, lsh := range this.lshGroup {
            lsh.UpdateDecoderVariance(tmpvar);
        }
        this.variance = tmpvar;
    }
    for _, lsh := range this.lshGroup {
        hash, _ = lsh.MinHash(vec, this.rphashStream.GetNumberOfBlurs());
        for _, h := range hash {
            c.AddID(h);
        }
    }
    this.centroidCounter.Add(c);
    return this.centroidCounter.GetCount();
};

func (this *Stream) GetCentroids() []float64 {
    if this.centroids == nil {
        this.Stream();
        var centroids []float64;
        for _, cent := range this.centroidCounter.GetTop() {
            centroids = append(centroids, cent.Centroid());
        }
        this.centroids = construct.NewKMeans(this.rphashStream.GetK(), centroids, this.centroidCounter.GetCounts()).GetCentroids();
    }
    return this.centroids;
};

func (this *Stream) GetCentroidsOfflineStep() []float64 {
    var centroids []float64;
    var counts []int32;
    for i := 0; i < len(this.centroidCounter.GetTop()); i++ {
        centroids = append(centroids, this.centroidCounter.GetTop()[i].Centroid());
        counts = append(counts, this.centroidCounter.GetCounts()[i]);
    }
    this.centroids = construct.NewKMeans(this.rphashStream.GetK(), centroids, counts).GetCentroids();
    count := int((vectors.Max(counts) + vectors.Min(counts)) / 2);
    counts = []int32{};
    for i := 0; i < this.rphashStream.GetK(); i++ {
        counts = append(counts, int32(count));
    }
    this.counts = counts;
    return this.centroids;
};

func (this *Stream) Stream() {
    vecs := this.rphashStream.GetVectorIterator();
    for i := 0; i < len(vecs); i++ {
        this.AddVectorOnlineStep(vecs[i]);
    }
};
