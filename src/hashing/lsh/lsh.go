/**
 * Locality Sensitive Hashing
 * 2nd Step
 * @author Sam Wenke
 * @author Jacob Franklin
 */
package lsh;

import (
    "math/rand"
);

var noise [][]float64;

/**
 * Choose a hash algorithm.
 */
type HashAlgorithm interface {
    Hash(s []int32) int32;
};

/**
 * Choose a decoder.
 */
type Decoder interface {
    SetVariance(parameterObject float64);
    GetDimensionality() int;
    Decode(f []float64) []int32;
    GetErrorRadius() float64;
    GetDistance() float64;
};

/**
 * Choose a projection.
 */
type Projector interface {
    Project(t []float64) []float64;
};

/**
 * There are 3 steps to LocalitySensitiveHashing,
 * Projecting,
 * Hashing,
 * Decoding.
 */
type LocalitySensitiveHashing struct {
    p []Projector;
    hal HashAlgorithm;
    dec Decoder;
    times int;
    radius float64;
    distance float64;
    random *rand.Rand;
};

func New(dec Decoder, p []Projector, hal HashAlgorithm, times int, randomseed int64) *LocalitySensitiveHashing {
    random := rand.New(rand.NewSource(randomseed));
    return &LocalitySensitiveHashing{
        p: p,
        hal: hal,
        dec: dec,
        times: times,
        random: random,
        radius: dec.GetErrorRadius() / float64(dec.GetDimensionality()),
    };
};

/**
 * Project -> Decode -> Hash!
 */
func (_lsh *LocalitySensitiveHashing) LSHHash(r []float64) int32{
    return _lsh.hal.Hash(_lsh.dec.Decode(_lsh.p[0].Project(r)));
};

func (_lsh *LocalitySensitiveHashing) UpdateDecoderVariance(variance float64) {
    _lsh.dec.SetVariance(variance);
};

func (_lsh *LocalitySensitiveHashing) GenNoiseTable(len, times int) {
    for j := 1; j < times; j++ {
        tmp := make([]float64, len);
        for k := 0; k < len; k++ {
            tmp[k] = _lsh.random.NormFloat64() * _lsh.radius;
        }
        noise = append(noise, tmp);
    }
};

func (_lsh *LocalitySensitiveHashing) LSHHashRadiusNo2Hash(r []float64, times int) []int32 {
    if noise == nil {
        _lsh.GenNoiseTable(len(r), times);
    }
    pr_r := _lsh.p[0].Project(r);
    nonoise := _lsh.dec.Decode(pr_r);
    ret := make([]int32, times * len(nonoise));
    copy(ret[0:len(nonoise)], nonoise[0:len(nonoise)])﻿;
    rtmp := make([]float64, len(pr_r));
    var tmp []float64;
    for j := 1; j < times; j++ {
        copy(rtmp[:len(pr_r)], pr_r[:len(pr_r)])﻿;
        tmp = noise[j - 1];
        for k := 0; k < len(pr_r); k++ {
            rtmp[k] = rtmp[k] + tmp[k];
        }
        nonoise = _lsh.dec.Decode(rtmp);
        copy(ret[j*len(nonoise):j*len(nonoise)+len(nonoise)], nonoise[0:len(nonoise)])﻿;
    }
    return ret;
};

func (_lsh *LocalitySensitiveHashing) LSHMinHashRadius(r []float64, radius float64, times int) int32 {
    pr_r := _lsh.p[0].Project(r);
    ret := _lsh.hal.Hash(_lsh.dec.Decode(pr_r));
    minret := ret;
    mindist := _lsh.dec.GetDistance();
    rtmp := make([]float64, len(pr_r));
    for j := 1; j < times; j++ {
        copy(rtmp[0:len(pr_r)], pr_r[0:len(pr_r)])﻿;
        for k := 0; k < len(pr_r); k++ {
            rtmp[k] = rtmp[k] + _lsh.random.NormFloat64() * _lsh.radius;
        }
        ret = _lsh.hal.Hash(_lsh.dec.Decode(rtmp));
        if _lsh.dec.GetDistance() < mindist {
            minret = _lsh.hal.Hash(_lsh.dec.Decode(rtmp));
            mindist = _lsh.dec.GetDistance();
        }
    }
    return ret;
};
