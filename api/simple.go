package api;

import (
  "github.com/wenkesj/rphash/reader"
  "github.com/wenkesj/rphash/simple"
);

// The RPHash call takes in an array of arrays,
// Each array has the dimension of a schema of a data set.
func NewRPHash(data [][]float64, numClusters int) *simple.Simple {
  RPHashObject := reader.NewSimpleArray(data, numClusters);
  return simple.NewSimple(RPHashObject);
};
