package stream;

import (
    "math/rand"
    "github.com/wenkesj/rphash/utils"
    "github.com/wenkesj/rphash/types"
    "github.com/wenkesj/rphash/defaults"
);

type Stream struct {
    counts []int32;
    centroids []float64;
    variance float64;
    centroidCounter types.CentroidCounter;
    random *rand.Rand;
    rphashStream types.RPHashObject;
    lshGroup []types.LSH;
    decoder types.Decoder;
    projector types.Projector;
    hash types.Hash;
    varTracker types.StatTest;
};

func NewStream(data []float64) *Stream {
    _rphashStream := defaults.NewRPHashObject();
    _random := rand.New(rand.NewSource(_rphashStream.GetRandomSeed()));
    _hash := defaults.NewHash(_rphashStream.GetHashModulus());
    _decoder := _rphashStream.GetDecoderType();
    _statTest := defaults.NewStatTest(0.01);
    projections := _rphashStream.GetNumberOfProjections();
    k := _rphashStream.GetK() * projections;
    _centroidCounter := defaults.NewCentroidCounter(k);
    _lshGroup := make([]types.LSH, projections);
    var _projector types.Projector;
    for i := 0; i < projections; i++ {
        _projector = defaults.NewProjector(_rphashStream.GetDimensions(), _decoder.GetDimensionality(), _random.Int63());
        _lshGroup[i] = defaults.NewLSH(_hash, _decoder, _projector);
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
    c := defaults.NewCentroid(vec);

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
        this.centroids = defaults.NewKMeans(this.rphashStream.GetK(), centroids, this.centroidCounter.GetCounts()).GetCentroids();
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
    this.centroids = defaults.NewKMeans(this.rphashStream.GetK(), centroids, counts).GetCentroids();
    count := int((utils.Max(counts) + utils.Min(counts)) / 2);
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
