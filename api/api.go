package api

import (
  "github.com/wilseypa/rphash-golang/reader"
  "github.com/wilseypa/rphash-golang/simple"
  "github.com/wilseypa/rphash-golang/stream"
)

func NewSimpleRPHash(data [][]float64, numClusters int) *simple.Simple {
  RPHashObject := reader.NewSimpleArray(data, numClusters)
  return simple.NewSimple(RPHashObject)
}

func NewStreamRPHash(dimensionality, numClusters int) *stream.Stream {
  rphashObject := reader.NewStreamObject(dimensionality, numClusters)
  return stream.NewStream(rphashObject)
}
