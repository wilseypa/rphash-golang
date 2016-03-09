package tests

import (
  "github.com/wilseypa/rphash-golang/parse"
  "io/ioutil"
  "testing"
)

const (
  // Data generated from https://www.mockaroo.com/.
<<<<<<< HEAD
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
  oldJSONData := oldJSON[dataLabel].([]interface{});
  newJSONData := newJSON[dataLabel].([]interface{});

  // Iterate over all the data and check consistency.
  for i, _ := range oldJSONData {
    oldJSONObject := oldJSONData[i].(map[string]interface{});
    newJSONObject := newJSONData[i].(map[string]interface{});
    for key, value := range oldJSONObject {
      newJSONValue, _ := parser.ConvertInterfaceToFloat64(newJSONObject[key]);
      oldJSONValue, _ := parser.ConvertInterfaceToFloat64(value);
      if oldJSONValue >= newJSONValue - float64(0.01) || oldJSONValue <= newJSONValue + float64(0.01) {
        t.Log("√ Matched key and normalized precision");
      } else {
        t.Log("Mismatched key or normalized precision off!");
      }
    }
  }
};
=======
  dataPath     = "./data/"
  dataFileName = "people.json"
  dataLabel    = "people"
)

func TestParser(t *testing.T) {
  parser := parse.NewParser()
  oldBytes, _ := ioutil.ReadFile(dataPath + dataFileName)
  oldJSON := parser.BytesToJSON(oldBytes)
  jsonFloats := parser.JSONToFloat64Matrix(dataLabel, oldJSON)
  newJSON := parser.Float64MatrixToJSON(dataLabel, jsonFloats)
  oldJSONData := oldJSON[dataLabel].([]interface{})
  newJSONData := newJSON[dataLabel].([]interface{})

  // Iterate over all the data and check consistency.
  for i, _ := range oldJSONData {
    oldJSONObject := oldJSONData[i].(map[string]interface{})
    newJSONObject := newJSONData[i].(map[string]interface{})
    for key, value := range oldJSONObject {
      newJSONValue, _ := parser.ConvertInterfaceToFloat64(newJSONObject[key])
      oldJSONValue, _ := parser.ConvertInterfaceToFloat64(value)
      if oldJSONValue >= newJSONValue-float64(0.01) || oldJSONValue <= newJSONValue+float64(0.01) {
        t.Log("√ Matched key and normalized precision")
      } else {
        t.Log("Mismatched key or normalized precision off!")
      }
    }
  }
}
>>>>>>> bbd90d3d48d2b3286575fccaab847e86681bbe39
