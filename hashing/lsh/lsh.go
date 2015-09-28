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


/**
 * There are 3 steps to LocalitySensitiveHashing,
 * Projecting,
 * Hashing,
 * Decoding.
 */
var (
    noise [][]float64;
    decode Decoder;
    project Projector;
    hash Hash;
);

type LocalitySensitiveHashing struct {
    times int;
    radius float64;
    random *rand.Rand;
};

func New(times int, randomseed int64) *LocalitySensitiveHashing {
    random := rand.New(rand.NewSource(randomseed));
    return &LocalitySensitiveHashing{
        times: times,
        random: random,
        radius: decode.GetErrorRadius() / float64(decode.GetDimensionality()),
    };
};

/**
 * Project -> Decode -> Hash
 */
func LSHHash(r []float64) int32{
    return hash.Hash(decode.Decode(project.Project(r)));
};

func UpdateDecoderVariance(variance float64) {
    decode.SetVariance(variance);
};

func (this *LocalitySensitiveHashing) GenNoiseTable(len, times int) {
    for j := 1; j < times; j++ {
        tmp := make([]float64, len);
        for k := 0; k < len; k++ {
            tmp[k] = this.random.NormFloat64() * this.radius;
        }
        noise = append(noise, tmp);
    }
};

func (this *LocalitySensitiveHashing) LSHHashRadiusNo2Hash(r []float64, times int) []int32 {
    if noise == nil {
        this.GenNoiseTable(len(r), times);
    }
    pr_r := project.Project(r);
    nonoise := decode.Decode(pr_r);
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
        nonoise = decode.Decode(rtmp);
        copy(ret[j*len(nonoise):j*len(nonoise)+len(nonoise)], nonoise[0:len(nonoise)])﻿;
    }
    return ret;
};

func (this *LocalitySensitiveHashing) LSHMinHashRadius(r []float64, radius float64, times int) int32 {
    pr_r := project.Project(r);
    ret := hash.Hash(decode.Decode(pr_r));
    minret := ret;
    mindist := decode.GetDistance();
    rtmp := make([]float64, len(pr_r));
    for j := 1; j < times; j++ {
        copy(rtmp[0:len(pr_r)], pr_r[0:len(pr_r)])﻿;
        for k := 0; k < len(pr_r); k++ {
            rtmp[k] = rtmp[k] + this.random.NormFloat64() * this.radius;
        }
        ret = hash.Hash(decode.Decode(rtmp));
        if decode.GetDistance() < mindist {
            minret = hash.Hash(decode.Decode(rtmp));
            mindist = decode.GetDistance();
        }
    }
    return ret;
};
