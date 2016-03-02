package api;

import (
  "github.com/wenkesj/rphash/reader"
  "github.com/wenkesj/rphash/simple"
  "github.com/wenkesj/rphash/stream"
);

func NewSimpleRPHash(data [][]float64, numClusters int) *simple.Simple {
  RPHashObject := reader.NewSimpleArray(data, numClusters);
  return simple.NewSimple(RPHashObject);
};

func NewStreamRPHash(dimensionality, numClusters int) *stream.Stream {
  rphashObject := reader.NewStreamObject(dimensionality, numClusters);
  return stream.NewStream(rphashObject);
};
