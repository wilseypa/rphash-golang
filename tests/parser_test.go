package tests;

import (
  "testing"
  "io/ioutil"
  "github.com/wenkesj/rphash/parse"
);

const dataPath = "./data/";
const dataFileName = "data.json";
const dataLabel = "people";

func TestParser(t *testing.T) {
  bytesContents, _ := ioutil.ReadFile(dataPath + dataFileName);
  parser := parse.NewParser();
  data := parser.BytesToJSON(bytesContents);
  floats := parser.JSONToFloat64Matrix(dataLabel, data);
  t.Log(parser.Float64MatrixToJSON(dataLabel, floats));
};
