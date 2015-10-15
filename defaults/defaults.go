package defaults;

import (
    "github.com/wenkesj/rphash/clusterer"
    "github.com/wenkesj/rphash/types"
    "github.com/wenkesj/rphash/decoder"
    "github.com/wenkesj/rphash/hash"
    "github.com/wenkesj/rphash/itemset"
    "github.com/wenkesj/rphash/lsh"
    "github.com/wenkesj/rphash/projector"
    "github.com/wenkesj/rphash/reader"
);

func NewMultiDecoder(dimension int, innerDec types.Decoder) types.MultiDecoder {
    return decoder.NewMultiDecoder(dimension, innerDec);
};

func NewProjector(n, t int, randomseed int64) types.Projector {
    return projector.NewDBFriendly(n, t, randomseed);
};

func NewHash(hashMod int32) types.Hash {
    return hash.NewMurmur(hashMod);
};

func NewKMeans(k int, centroids []float64, counts []int32) types.Clusterer {
    return clusterer.NewKMeans(k, centroids, counts);
};

func NewCountMinSketch() types.ItemSet {
    return itemset.NewKHHCountMinSketch();
};

func NewCentroid(vec []float64) types.Centroid {
    return itemset.NewCentroid(vec);
};

func NewCentroidCounter(k int) types.ItemSet {
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

func NewRPHashObject() types.RPHashObject {
    return reader.NewStreamObject();
};
