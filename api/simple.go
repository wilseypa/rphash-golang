package api;

import (
  "github.com/wenkesj/rphash/reader"
  "github.com/wenkesj/rphash/simple"
);

func NewRPHash(data [][]float64, numClusters int) *simple.Simple {
  RPHashObject := reader.NewSimpleArray(data, numClusters);
  return simple.NewSimple(RPHashObject);
};
