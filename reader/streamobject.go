package reader;

import (
    "github.com/wenkesj/rphash/decoder"
    "github.com/wenkesj/rphash/types"
    "github.com/wenkesj/rphash/utils"
);

type StreamObject struct {
    data [][]float64;
    numberOfProjections int;
    decoderMultiplier int;
    randomSeed int64;
    numberOfBlurs int;
    k int;
    dimension int;
    hashModulus int32;
    centroids [][]float64;
    topIDs []int32;
    decoder types.Decoder;
};

func NewStreamObject(dimension, k int) *StreamObject {
    innerDecoder := decoder.InnerDecoder();
    decoderMultiplier := 1;
    decoder := decoder.NewMultiDecoder(decoderMultiplier * innerDecoder.GetDimensionality(), innerDecoder);
    var centroids [][]float64;
    var topIDs []int32;
    return &StreamObject{
        decoder: decoder,
        dimension: dimension,
        randomSeed: int64(0),
        hashModulus: 2147483647,
        decoderMultiplier: decoderMultiplier,
        numberOfProjections: 2,
        numberOfBlurs: 2,
        k: k,
        data: nil,
        topIDs: topIDs,
        centroids: centroids,
    };
};

func (this *StreamObject) GetK() int {
    return this.k;
};

func (this *StreamObject) GetDimensions() int {
    return this.dimension;
};

func (this *StreamObject) GetRandomSeed() int64 {
    return this.randomSeed;
};

func (this *StreamObject) GetNumberOfBlurs() int {
    return this.numberOfBlurs;
};

func (this *StreamObject) GetVectorIterator() [][]float64 {
    return this.data;
};

func (this *StreamObject) GetCentroids() [][]float64 {
    return this.centroids;
};

func (this *StreamObject) GetPreviousTopID() []int32 {
    return this.topIDs;
};

func (this *StreamObject) SetPreviousTopID(top []int32) {
    this.topIDs = top;
};

func (this *StreamObject) AddCentroid(v []float64) {
    this.centroids = append(this.centroids, v);
};

func (this *StreamObject) SetCentroids(l [][]float64) {
    this.centroids = l;
};

func (this *StreamObject) GetNumberOfProjections() int {
    return this.numberOfProjections;
};

func (this *StreamObject) SetNumberOfProjections(probes int) {
    this.numberOfProjections = probes;
};

func (this *StreamObject) SetInnerDecoderMultiplier(multiDim int) {
    this.decoderMultiplier = multiDim;
};

func (this *StreamObject) GetInnerDecoderMultiplier() int {
    return this.decoderMultiplier;
};

func (this *StreamObject) SetNumberOfBlurs(parseInt int) {
    this.numberOfBlurs = parseInt;
};

func (this *StreamObject) SetRandomSeed(parseLong int64) {
    this.randomSeed = parseLong;
};

func (this *StreamObject) GetHashModulus() int32 {
    return this.hashModulus;
};

func (this *StreamObject) SetHashModulus(parseLong int32) {
    this.hashModulus = int32(parseLong);
};

func (this *StreamObject) SetDecoderType(dec types.Decoder) {
    this.decoder = dec;
};

func (this *StreamObject) GetDecoderType() types.Decoder {
    return this.decoder;
};

func (this *StreamObject) SetVariance(data [][]float64) {
    this.decoder.SetVariance(utils.VarianceSample(data, 0.01));
};
