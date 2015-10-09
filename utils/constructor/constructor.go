package constructor;

import (
    "github.com/wenkesj/rphash/types"
);

type Constructor struct {
    decoder types.DecoderConstructor;
    projector types.ProjectorConstructor;
    hash types.HashConstructor;
    kmeans types.KMeansConstructor;
    centroid types.CentroidConstructor;
    centroidcounter types.CentroidCounterConstructor;
    lsh types.LSHConstructor;
    stattest types.StatTestConstructor;
    streamobject types.StreamObjectConstructor;
};

func New(decoder types.DecoderConstructor,
        projector types.ProjectorConstructor,
        hash types.HashConstructor,
        kmeans types.KMeansConstructor,
        centroid types.CentroidConstructor,
        centroidcounter types.CentroidCounterConstructor,
        lsh types.LSHConstructor,
        stattest types.StatTestConstructor,
        streamobject types.StreamObjectConstructor) *Constructor {

    return &Constructor{
        decoder: decoder,
        projector: projector,
        hash: hash,
        kmeans: kmeans,
        centroid: centroid,
        centroidcounter: centroidcounter,
        lsh: lsh,
        stattest: stattest,
        streamobject: streamobject,
    };
}

func (this *Constructor) NewDecoder() types.Decoder {
    return this.decoder.New();
};

func (this *Constructor) NewProjector(n, t int, randomseed int64) types.Projector {
    return this.projector.New(n, t, randomseed);
};

func (this *Constructor) NewHash(hashMod int32) types.Hash {
    return this.hash.New(hashMod);
};

func (this *Constructor) NewKMeans(k int, centroids []float64, counts []int32) types.KMeans {
    return this.kmeans.New(k, centroids, counts);
};

func (this *Constructor) NewCentroid(vec []float64) types.Centroid {
    return this.centroid.New(vec);
};

func (this *Constructor) NewCentroidCounter(k int) types.CentroidCounter {
    return this.centroidcounter.New(k);
};

func (this *Constructor) NewLSH(hash types.Hash, decoder types.Decoder, projector types.Projector) types.LSH {
    return this.lsh.New(hash, decoder, projector);
};

func (this *Constructor) NewStatTest(vari float64) types.StatTest {
    return this.stattest.New(vari);
};

func (this *Constructor) NewStreamObject() types.StreamObject {
    return this.streamobject.New();
};
