package defaults;

import (
    "github.com/wilseypa/rphash-golang/clusterer"
    "github.com/wilseypa/rphash-golang/types"
    "github.com/wilseypa/rphash-golang/decoder"
    "github.com/wilseypa/rphash-golang/hash"
    "github.com/wilseypa/rphash-golang/utils"
    "github.com/wilseypa/rphash-golang/itemset"
    "github.com/wilseypa/rphash-golang/lsh"
    "github.com/wilseypa/rphash-golang/projector"
    "github.com/wilseypa/rphash-golang/reader"
);

func NewMultiDecoder(dimension int, innerDec types.Decoder) types.Decoder {
    return decoder.NewMultiDecoder(dimension, innerDec);
};

func NewProjector(n, t int, randomseed int64) types.Projector {
    return projector.NewDBFriendly(n, t, randomseed);
};

func NewHash(hashMod int64) types.Hash {
    return hash.NewMurmur(hashMod);
};

func NewKMeansStream(k int, centroids [][]float64, counts []int64) types.Clusterer {
    return clusterer.NewKMeansStream(k, centroids, counts);
};

func NewKMeansSimple(k int, centroids [][]float64) types.Clusterer {
    return clusterer.NewKMeansSimple(k, centroids);
};

func NewCentroidStream(vec []float64) types.Centroid {
    return itemset.NewCentroidStream(vec);
};

func NewCentroidSimple(dim int, id int64) types.Centroid {
    return itemset.NewCentroidSimple(dim, id);
};

func NewCountMinSketch(k int) types.CountItemSet {
    return itemset.NewKHHCountMinSketch(k);
};

func NewCentroidCounter(k int) types.CentroidItemSet {
    return itemset.NewKHHCentroidCounter(k);
};

func NewLSH(hash types.Hash, decoder types.Decoder, projector types.Projector) types.LSH {
    return lsh.NewLSH(hash, decoder, projector);
};

func NewStatTest(vari float64) types.StatTest {
    return utils.NewStatTest(vari);
};

func NewSimpleArray(k int, data [][]float64) types.RPHashObject {
    return reader.NewSimpleArray(data, k);
};

func NewRPHashObject(dimension, k int) types.RPHashObject {
    return reader.NewStreamObject(dimension, k);
};
