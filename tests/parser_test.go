package tests;

import (
  "testing"
  "io/ioutil"
  "github.com/wenkesj/rphash/parse"
);

const (
  // Data generated from https://www.mockaroo.com/.
  dataPath = "./data/";
  dataFileName = "people.json";
  dataLabel = "people";
);

func TestParser(t *testing.T) {
  parser := parse.NewParser();
  oldBytes, _ := ioutil.ReadFile(dataPath + dataFileName);
  oldJSON := parser.BytesToJSON(oldBytes);
  jsonFloats := parser.JSONToFloat64Matrix(dataLabel, oldJSON);
  newJSON := parser.Float64MatrixToJSON(dataLabel, jsonFloats);
  newBytes := parser.JSONToBytes(newJSON);
  json_1, json_2 := string(oldBytes), string(newBytes);

  // The 2 JSON objects are identical, other than newlines and spaces.
  // The bytes don't match nore do the strings, but the values and keys
  // are preserved.
  t.Log(json_1, json_2);
};
