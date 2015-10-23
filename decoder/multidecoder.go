package decoder;

import (
    "math"
    "github.com/wenkesj/rphash/types"
);

type MultiDecoder struct {
    innerDec types.Decoder;
    dimension int;
    rounds int;
    distance float64;
};

func NewMultiDecoder(dimension int, innerDec types.Decoder) *MultiDecoder {
    rounds := int(math.Ceil(float64(dimension) / float64(innerDec.GetDimensionality())));
    return &MultiDecoder{
        dimension: dimension,
        rounds: rounds,
        innerDec: innerDec,
        distance: -1.0,
    };
};

func (this *MultiDecoder) GetDimensionality() int {
    return this.dimension;
};

func (this *MultiDecoder) Decode(f []float64) []int64 {
    if this.innerDec.GetDimensionality() == len(f) {
        return this.innerDec.Decode(f);
    }
    innerpartition := make([]float64, this.innerDec.GetDimensionality());
    copy(innerpartition[:int(math.Min(float64(len(f)), float64(len(innerpartition))))], f[:int(math.Min(float64(len(f)), float64(len(innerpartition))))]);
    tmp := this.innerDec.Decode(innerpartition);
    retLength := len(tmp);
    ret := make([]int64, retLength * this.rounds);
    copy(ret[:retLength], tmp[:retLength]);
    this.distance = this.innerDec.GetDistance();
    for i := 1; i < this.rounds; i++ {
		copy(innerpartition[0:int(math.Min(float64(len(f) - i * this.innerDec.GetDimensionality()), float64(len(innerpartition))))], f[i * this.innerDec.GetDimensionality():i * this.innerDec.GetDimensionality() + int(math.Min(float64(len(f) - i * this.innerDec.GetDimensionality()), float64(len(innerpartition))))]);
        tmp = this.innerDec.Decode(innerpartition);
        this.distance += this.innerDec.GetDistance();
        copy(ret[i * retLength:i * retLength + retLength], tmp[0:retLength]);
    }
    return ret;
};

func (this *MultiDecoder) GetErrorRadius() float64 {
    return this.innerDec.GetErrorRadius();
};

func (this *MultiDecoder) GetDistance() float64 {
    return this.distance;
};

func (this *MultiDecoder) GetVariance() float64 {
    return this.innerDec.GetVariance();
};

func (this *MultiDecoder) SetVariance(parameterObject float64) {
    this.innerDec.SetVariance(parameterObject);
};
