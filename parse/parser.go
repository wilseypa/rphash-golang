package parse;

import (
  "reflect"
  "encoding/json"
);

type Parser struct {
  schema map[string]reflect.Type;
  label string;
};

func NewParser() *Parser {
  return &Parser{
    label: "",
    schema: nil,
  };
};

// Convert an array of bytes to a JSON struct.
func (this *Parser) BytesToJSON(bytesContents []byte) map[string]interface{} {
  var data map[string]interface{}
  if err := json.Unmarshal(bytesContents, &data); err != nil {
    panic(err);
  }
  return data;
};

// Convert a json object with a schema to an array of 64 bit floats.
func (this *Parser) JSONToFloat64(jsonMap map[string]interface{}) []float64 {
  // var i = 0;

  // Create an array of 64 bit floats of the same size.
  result := make([]float64, len(this.schema));

  // Iterate over the schema fields and assign floating point values to each field value.
  // for _key, _type := range this.schema {
  //   // value := jsonMap[_key];
  //   // result[i] = convert value to float value;
  //   // i++;
  // }
  return result;
};

// Convert an array of 64 bit floats to JSON according to a schema.
func (this *Parser) Float64ToJSON(floats []float64) map[string]interface{} {
  // var i = 0;

  // Create an JSON object.
  jsonMap := make(map[string]interface{});

  // Iterate over the schema fields and assign a value to each key from the floating point array.
  // for _key, _type := range this.schema {
  //   float := floats[i];
  //   jsonMap[_key] = convert the floating point number back to the real value;
  //   i++;
  // }
  return jsonMap;
};

// Convert a JSON table to a array of float64 arrays.
// The data should come in like this:
// {
//  "label": [{
//   "field-1": "value-1",
//   ...
//  }, {
//  "field-1": "value-1",
//  ...
//  }]
// }
func (this *Parser) JSONToFloat64Matrix(label string, dataSet map[string]interface{}) [][]float64 {
  // Assign a label to the specific schema.
  this.label = label;

  // Read the data in as an array of json objects.
  data := dataSet[label].([]interface{});
  count := len(data);

  // Allocate an array of arrays for the return.
  matrix := make([][]float64, count, count);

  // Create a schema based on an entry in the data.
  this.schema = this.CreateSchema(data[0].(map[string]interface{}));

  // Convert the json data to weighted float values.
  for i := 0; i < count; i++ {
    matrix[i] = this.JSONToFloat64(data[i].(map[string]interface{}));
  }
  return matrix;
};

// Convert a matrix of 64 bit floats to JSON according to a json schema.
// label - string associated with JSON data set schema.
// data - the array of arrays associated with the entries of data.
func (this *Parser) Float64MatrixToJSON(label string, dataSet [][]float64) map[string]interface{} {
  count := len(dataSet);

  // Create an array of JSON objects.
  data := make([]map[string]interface{}, count);

  // Create a json object to hold the array of JSON objects with the specific label.
  result := make(map[string]interface{});

  // Convert the weighted float values back to the JSON using the schema.
  for i := 0; i < count; i++ {
    data[i] = this.Float64ToJSON(dataSet[i]);
  }

  // Assign the label to the json data.
  result[this.label] = data;
  return result;
};

func (this *Parser) CreateSchema(jsonMap map[string]interface{}) map[string]reflect.Type {
  keys := len(jsonMap);

  // Create an empty schema
  schema := make(map[string]reflect.Type, keys);

  // Assign the key associated with the JSON field to its value type.
  for _key, _value := range jsonMap {
    schema[_key] = reflect.TypeOf(_value);
  }
  return schema;
};
