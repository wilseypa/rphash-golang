/**
* Copyright (c) 2015, Philip A. Wilsey
* All rights reserved.
*
* Redistribution and use in source and binary forms, with or without
* modification, are permitted provided that the following conditions are met:
*
* * Redistributions of source code must retain the above copyright notice, this
*   list of conditions and the following disclaimer.
*
* * Redistributions in binary form must reproduce the above copyright notice,
*   this list of conditions and the following disclaimer in the documentation
*   and/or other materials provided with the distribution.
*
* THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
* AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
* IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
* DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
* FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
* DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
* SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
* CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
* OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
* OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*
* Sam Wenke (@wenkesj)
* Jacob Franklin (@frankljbe)
 */

package main

import (
  "fmt"
  "github.com/jessevdk/go-flags"
  "github.com/wenkesj/rphash/api"
  "github.com/wenkesj/rphash/parse"
  "io/ioutil"
  "os"
)

const version = "v1.0.0"

type Options struct {
  Verbose []bool `short:"v" long:"verbose" description:"Verbose output"`
}

type VersionCommand struct {
  All bool `short:"v" long:"version" description:"Display current version"`
}

type ClusterCommand struct {
  Input    string `short:"i" long:"input" optional:"no" description:"Path to input file"`
  Output   string `short:"o" long:"output" optional:"no" description:"Path to output file"`
  Label    string `short:"l" long:"label" optional:"no" description:"Dataset label"`
  Clusters int    `short:"c" long:"clusters" optional:"no" description:"Provide the number of centroids to compute"`
  Stream   bool   `short:"s" long:"stream" optional:"yes" description:"Start a stream"`
}

func (this *VersionCommand) Execute(args []string) error {
  fmt.Printf(version)
  return nil
}

func (this *ClusterCommand) Execute(args []string) error {
  inputFileName := this.Input
  outputFileName := this.Output
  label := this.Label
  clusters := this.Clusters

  parser := parse.NewParser()
  bytes, _ := ioutil.ReadFile(inputFileName)
  jsonData := parser.BytesToJSON(bytes)
  data := parser.JSONToFloat64Matrix(label, jsonData)

  if this.Stream {
    cluster := api.NewStreamRPHash(len(data[0]), clusters)
    for _, vector := range data {
      cluster.AppendVector(vector)
    }
    topCentroids := cluster.GetCentroids()
    jsonCentroids := parser.Float64MatrixToJSON(label, topCentroids)
    jsonBytes := parser.JSONToBytes(jsonCentroids)
    err := ioutil.WriteFile(outputFileName, jsonBytes, 0644)
    if err != nil {
      panic(err)
    }
  } else {
    cluster := api.NewSimpleRPHash(data, clusters)
    topCentroids := cluster.GetCentroids()
    jsonCentroids := parser.Float64MatrixToJSON(label, topCentroids)
    jsonBytes := parser.JSONToBytes(jsonCentroids)
    err := ioutil.WriteFile(outputFileName, jsonBytes, 0644)
    if err != nil {
      panic(err)
    }
  }
  return nil
}

var options Options
var versionCommand VersionCommand
var clusterCommand ClusterCommand
var parser = flags.NewParser(&options, flags.Default)

func main() {
  parser.AddCommand("cluster",
    "Run a clusterer",
    "Run a specified clustering algorithm",
    &clusterCommand)
  parser.AddCommand("version",
    "Display current version",
    "Display the current version of RPHash installed.",
    &versionCommand)
  if _, err := parser.Parse(); err != nil {
    os.Exit(1)
  }
}
