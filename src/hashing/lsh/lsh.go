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

type HashAlgorithm interface {
    hash(s []int32) int32;
};

type Decoder interface {
    setVariance(parameterObject float64);
    getDimensionality() int;
    decode(f []float64) []int32;
    getErrorRadius() float64;
    getDistance() float64;
};

type Projector interface {
    project(t []float64) []float64;
};

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
        radius: dec.getErrorRadius() / float64(dec.getDimensionality()),
    };
};

func (_lsh *LocalitySensitiveHashing) lshHash(r []float64) int32{
    return _lsh.hal.hash(_lsh.dec.decode(_lsh.p[0].project(r)));
};

func (_lsh *LocalitySensitiveHashing) updateDecoderVariance(variance float64) {
    _lsh.dec.setVariance(variance);
};

func (_lsh *LocalitySensitiveHashing) genNoiseTable(len, times int) {
    for j := 1; j < times; j++ {
        tmp := make([]float64, len);
        for k := 0; k < len; k++ {
            tmp[k] = _lsh.random.NormFloat64() * _lsh.radius;
        }
        noise = append(noise, tmp);
    }
};

func (_lsh *LocalitySensitiveHashing) lshHashRadiusNo2Hash(r []float64, times int) []int32 {
    if noise == nil {
        _lsh.genNoiseTable(len(r), times);
    }
    pr_r := _lsh.p[0].project(r);
    nonoise := _lsh.dec.decode(pr_r);
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
        nonoise = _lsh.dec.decode(rtmp);
        copy(ret[j*len(nonoise):j*len(nonoise)+len(nonoise)], nonoise[0:len(nonoise)])﻿;
    }
    return ret;
};

func (_lsh *LocalitySensitiveHashing) lshMinHashRadius(r []float64, radius float64, times int) int32 {
    pr_r := _lsh.p[0].project(r);
    ret := _lsh.hal.hash(_lsh.dec.decode(pr_r));
    minret := ret;
    mindist := _lsh.dec.getDistance();
    rtmp := make([]float64, len(pr_r));
    for j := 1; j < times; j++ {
        copy(rtmp[0:len(pr_r)], pr_r[0:])﻿;
        for k := 0; k < len(pr_r); k++ {
            rtmp[k] = rtmp[k] + _lsh.random.NormFloat64() * _lsh.radius;
        }
        ret = _lsh.hal.hash(_lsh.dec.decode(rtmp));
        if _lsh.dec.getDistance() < mindist {
            minret = _lsh.hal.hash(_lsh.dec.decode(rtmp));
            mindist = _lsh.dec.getDistance();
        }
    }
    return ret;
};
