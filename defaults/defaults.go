package defaults;

import (
    "github.com/wenkesj/rphash/clusterer"
    "github.com/wenkesj/rphash/types"
    "github.com/wenkesj/rphash/decoder"
    "github.com/wenkesj/rphash/hash"
    "github.com/wenkesj/rphash/utils"
    "github.com/wenkesj/rphash/itemset"
    "github.com/wenkesj/rphash/lsh"
    "github.com/wenkesj/rphash/projector"
    "github.com/wenkesj/rphash/reader"
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
