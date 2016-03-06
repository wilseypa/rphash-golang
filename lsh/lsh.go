package lsh

import (
  "github.com/wenkesj/rphash/types"
  "math/rand"
)

type LSH struct {
  hash      types.Hash
  decoder   types.Decoder
  projector types.Projector
  distance  float64
  noise     [][]float64
  radius    float64
}

func NewLSH(hash types.Hash,
  decoder types.Decoder,
  projector types.Projector) *LSH {
  return &LSH{
    hash:      hash,
    decoder:   decoder,
    projector: projector,
    distance:  0.0,
    noise:     nil,
    radius:    decoder.GetErrorRadius() / float64(decoder.GetDimensionality()),
  }
}

func (this *LSH) GenerateNoiseTable(len, times int) {
  this.noise = [][]float64{}
  for j := 1; j < times; j++ {
    tmp := make([]float64, len)
    for k := 0; k < len; k++ {
      tmp[k] = rand.NormFloat64() * this.radius
    }
    this.noise = append(this.noise, tmp)
  }
}

func (this *LSH) LSHHashStream(r []float64, times int) []int64 {
  if this.noise == nil {
    this.GenerateNoiseTable(len(r), times)
  }
  pr_r := this.projector.Project(r)
  nonoise := this.decoder.Decode(pr_r)
  ret := make([]int64, times*len(nonoise))
  copy(ret[0:len(nonoise)], nonoise[0:])

  rtmp := make([]float64, len(pr_r))
  var tmp []float64
  for j := 1; j < times; j++ {
    copy(rtmp[0:len(pr_r)], pr_r[0:])
    tmp = this.noise[j-1]
    for k := 0; k < len(pr_r); k++ {
      rtmp[k] = rtmp[k] + tmp[k]
    }
    nonoise = this.decoder.Decode(rtmp)
    copy(ret[j*len(nonoise):j*len(nonoise)+len(nonoise)], nonoise[0:])
  }
  return ret
}

func (this *LSH) LSHHashSimple(r []float64) int64 {
  projectedSpace := this.projector.Project(r)
  decodedSpace := this.decoder.Decode(projectedSpace)
  hashedResult := this.hash.Hash(decodedSpace)
  return hashedResult
}

func (this *LSH) Distance() float64 {
  return this.distance
}

func (this *LSH) UpdateDecoderVariance(vari float64) {
  this.decoder.SetVariance(vari)
}
