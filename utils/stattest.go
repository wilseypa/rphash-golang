package utils;

import (
    "math/rand"
);

type StatTest struct {
    sampRatio float64;
};

func NewStatTest(sampRatio float64) *StatTest{
    return &StatTest{
        sampRatio: sampRatio,
    };
};

func HashCode(num int64) int64 {
    return int64(uint64(num) ^ uint64(num) >> 32);
};

func (this *StatTest) UpdateVarianceSample(row []float64) float64 {
    var n float64 = 0;
    var mean float64 = 0;
    var M2 float64 = 0;
    if rand.Float64() > this.sampRatio {
        return M2 / (n - 1.0);
    }
    for _, x := range row {
        n++;
        delta := x - mean;
        mean = mean + delta / n;
        M2 = M2 + delta * (x - mean);
    }
    if n < 2 {
        return 0;
    }
    return  M2 / (n - 1.0);
};

func VarianceSample(data [][]float64, sampRatio float64) float64 {
    var n float64 = 0;
    var mean float64 = 0;
    var M2 float64 = 0;
    len := len(data);
    for i := 0; i < int(sampRatio) * len; i++ {
        row := data[rand.Intn(len)];
        for _, x := range row {
            n++;
            delta := x - mean;
            mean = mean + delta / n;
            M2 = M2 + delta * (x - mean);
        }
    }
    if n < 2 {
        return 0;
    }
    return  M2 / (n - 1.0);
};


func (this *StatTest) VarianceAll(data [][]float64) float64 {
    var n float64 = 0;
    var mean float64 = 0;
    var M2 float64 = 0;
    for _, row := range data {
        for _, x := range row {
            n++;
            delta := x - mean;
            mean = mean + delta / n;
            M2 = M2 + delta * (x - mean);
        }
    }
    if n < 2 {
        return 0;
    }
    return  M2 / (n - 1.0);
};

func (this *StatTest) AverageAll(data [][]float64) float64{
    var n float64 = 0;
    var mean float64 = 0;
    for _, row := range data {
        for _, x := range row {
            n++;
            mean += x;
        }
    }
    return mean / n;
};

func (this *StatTest) VarianceCol(data [][]float64) []float64 {
    leng := len(data);
    if leng < 1 {
        return nil;
    }
    vars := make([]float64, len(data[0]));
    var n float64 = 0;
    var mean float64 = 0;
    var M2 float64 = 0;
    for i := 0; i < leng; i++ {
        n = 0;
        mean = 0;
        M2 = 0;
        for _, x := range data[i] {
            n++;
            delta := x - mean;
            mean = mean + delta / n;
            M2 = M2 + delta * (x - mean);
        }
        if n < 2 {
            vars[i] = 0;
        } else {
            vars[i] = M2 / (n - 1.0);
        };
    }
    return vars;
};

func (this *StatTest) AverageCol(data [][]float64) []float64 {
    n := len(data);
    if n < 1 {
        return nil;
    }
    d := len(data[0]);
    avgs := make([]float64, d);
    for _, tmp := range data {
        for j :=0; j < d; j++ {
            avgs[j] += (tmp[j] / float64(n));
        }
    }
    return avgs;
};


func (this *StatTest) Variance(row []float64) float64 {
    var n float64 = 0;
    var mean float64 = 0;
    var M2 float64 = 0;
    for _, x := range row {
        n++;
        delta := x - mean;
        mean = mean + delta / n;
        M2 = M2 + delta * (x - mean);
    }
    if n < 2 {
        return 0;
    }
    return  M2 / (n - 1.0);
};
