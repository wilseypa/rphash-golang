package stream;

import (
    "math/rand"
    "github.com/wenkesj/rphash/utils/constructor"
    "github.com/wenkesj/rphash/itemset/centroid"
    "github.com/wenkesj/rphash/utils/vectors"
    "github.com/wenkesj/rphash/types"
);

type Stream struct {
    counts []int32;
    centroids [][]float64;
    centroidCounter types.CentroidCounter;
    random *rand.Rand;
    rphashStream types.StreamObject;
    lshGroup []types.LSH;
    decoder types.Decoder;
    projector types.Projector;
    hash types.Hash;
    varTracker types.StatTest;
};

func New(centroidCounter types.CentroidCounterConstructor,
            rphashStream types.StreamObjectConstructor,
            lsh types.LSHConstructor,
            kmeans types.KMeansConstructor,
            centroid types.CentroidConstructor,
            decoder types.DecoderConstructor,
            projector types.ProjectorConstructor,
            hash types.HashConstructor,
            varTracker types.StatTestConstructor) *Stream {

    construct := constructor.New(decoder, projector, hash, kmeans, centroid, centroidCounter, lsh, varTracker, rphashStream);
    rphashStream := construct.NewStreamObject();
    random := rand.New(rand.NewSource(rphashStream.GetRandomSeed()));
    hash := construct.NewHash(rphashStream.GetHashModulus());
    decoder := construct.NewDecoder();
    projections := rphashStream.GetNumberOfProjections();
    varTracker := construct.NewStatTest(0.01);
    k = k * projections;
    centroidCounter := construct.NewCentroidCounter(k);
    for i := 0; i < projections; i++ {
        projector := construct.NewProjector(rphashStream.GetDimensions(), decoder.GetDimensionality(), random.Int63());
        lshGroup[i] = construct.NewLSH(hash, decoder, projector);
    }
    return &Stream{
        counts: nil,
        centroids: nil,
        centroidCounter: centroidCounter,
        random: random,
        rphashStream: rphashStream,
        lshGroup: lshGroup,
        hash: hash,
        decoder: decoder,
        projector: projector,
        varTracker: varTracker,
    };
};

func (this *Stream) AddVectorOnlineStep(vec []float64) int32 {
    var hash []int32;
    c := construct.NewCentroid(vec);

    tmpvar := this.varTracker.UpdateVarianceSample(vec);
    if variance != tmpvar {
        for _, lsh := range this.lshGroup {
            lsh.UpdateDecoderVariance(tmpvar);
        }
        variance = tmpvar;
    }
    for _, lsh := range this.lshGroup {
        hash = lsh.MinHash(vec, this.rphashStream.GetNumberOfBlurs());
        for _, h := range hash {
            c.AddID(h);
        }
    }
    this.centroidCounter.Add(c);
    return this.centroidCounter.count;
};

func (this *StreamObject) GetCentroids() [][]float64 {
    if this.centroids == nil {
        this.Stream();
        var centroids [][]float64;
        for _, cent := range this.centroidCounter.GetTop() {
            centroids.append(centroids, cent.Centroid());
        }
        this.centroids := construct.NewKMeans(this.rphashStream.GetK(), centroids, this.centroidCounter.GetCounts()).GetCentroids();
    }
    return this.centroids;
};

func (this *Stream) GetCentroidsOfflineStep() [][]float64 {
    var centroids [][]float64;
    var counts []int32;
    for i = 0; i < len(this.centroidCounter.GetTop()); i++ {
        centroids.append(centroids, this.centroidCounter.GetTop()[i].Centroid());
        counts.append(counts, this.centroidCounter.GetCounts()[i]);
    }
    this.centroids = construct.NewKMeans(this.rphashStream.GetK(), centroids, counts).GetCentroids();
    count := int((vectors.Max(counts) + vectors.Min(counts)) / 2);
    counts = []int32{};
    for i := 0; i < this.rphashStream.GetK(); i++ {
        counts.append(counts, int32(count));
    }
    this.counts := counts;
    return this.centroids;
};

func (this *Stream) Stream() {
    vecs := this.rphashStream.GetVectorIterator();
    for i := 0; i < len(vecs); i++ {
        this.AddVectorOnlineStep(vecs[i]);
    }
};
