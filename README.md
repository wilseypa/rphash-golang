# Scalable Big Data Clustering by Random Projection Hashing #
[![Build Status](https://travis-ci.org/wenkesj/rphash.svg)](https://travis-ci.org/wenkesj/rphash)
[![Release Status](https://img.shields.io/badge/version-1.0.0-blue.svg)](https://github.com/wenkesj/rphash/releases)

The goal is to create a simple, secure, distributed, scalable, and parallel clustering algorithm to be used on almost any system.

**Clustering** is a core concept in data analysis. Issues arise with scalability and dimensionality, ever changing environments and compatibility, insecure communications and data movement

**The solution** is secure, reliable, and fast data for large-scale distributed systems.


**The Algorithm** provides more accurate clusters and an inherently distributed system.

![Clusters](https://github.com/wenkesj/rphash/blob/master/clusters.png)

**Random Projection Hash (RPHash)** has been created for maximizing parallel computation
while providing scalability for large scale deployment. It's suitable for high dimensional data sets and is
scalable and streamline.

![Overview](https://github.com/wenkesj/rphash/blob/master/overview.png)

# Table of contents #
+ **[Installing and Testing](https://github.com/wenkesj/rphash#installing-testing-and-plotting)**
+ **[API](https://github.com/wenkesj/rphash#api)**
+ **[Examples](https://github.com/wenkesj/rphash/blob/master/examples/rphash.go)**
+ **[Learn more](https://github.com/wenkesj/rphash/blob/master/REFERENCES.md)**
+ **[Versioning and updates](https://github.com/wenkesj/rphash/blob/master/CHANGELOG.md)**
+ **[Pull requests welcome](https://github.com/wenkesj/rphash/blob/master/TODO.md)**
+ **[Developers](https://github.com/wenkesj/rphash#developers)**

# Installing Testing and Plotting #
```sh
git clone --depth=50 --branch=master https://github.com/wenkesj/rphash.git wenkesj/rphash
cd wenkesj/rphash
export GOPATH=$HOME/<your-gopath>
export PATH=$HOME/<your-gopath>/bin:$PATH
go get -t -v ./...
sh install
```

## Example ##
Here is a simple example of RPHash clustering on a single node. The Algorithm Maps the functions and then Reduces in order to find the of the clusters centroids. It takes in a JSON file and assigns weights to field value, performs the RPHash clustering algorithm, and then outputs the results to a JSON file. The field value weights will be used for multi-dimensional clustering. From the centroids, you can obtain patterns and information from the dataset.

```sh
# cd examples
go run rphash.go
```

```go
package main;

import (
  "io/ioutil"
  "github.com/wenkesj/rphash/api"
  "github.com/wenkesj/rphash/parse" // A parser for JSON to Float64 arrays
);

// number of rows = 100
// number of dimensions = 10
var numberOfClusters = 3;

const (
  exampleInputFileName = "input.json";
  exampleOutputFileName = "output.json";
  exampleDataLabel = "people";
);

func main() {
  // Create a new parser for JSON -> []Float64 and []Float64 -> JSON.
  parser := parse.NewParser();

  // Read in the JSON file as bytes.
  bytes, _ := ioutil.ReadFile(exampleInputFileName);

  // Parse the bytes to JSON.
  jsonData := parser.BytesToJSON(bytes);

  // Pass in the label and JSON.
  // Returns an array of []Float64's that correspond to normalized weights.
  data := parser.JSONToFloat64Matrix(exampleDataLabel, jsonData);

  // Create an RPHash Cluster.
  // Pass in the data, and the number of clusters you want to create from the data.
  cluster := api.NewRPHash(data, numberOfClusters);

  // In reality, one will use our MapReduce enviornment to distribute
  // and harness the full power of the clustering framework.
  // For simplicities sake, perform Map
  cluster.Run();

  // Get the centroids from the data.
  centroids := cluster.GetCentroids();

  // Parse the centroids back to JSON and read the information obtained.
  // Output a new file to json for readability.
  jsonCentroids := parser.Float64MatrixToJSON(exampleDataLabel, centroids);
  jsonBytes := parser.JSONToBytes(jsonCentroids);
  err := ioutil.WriteFile(exampleOutputFileName, jsonBytes, 0644);
  if err != nil {
    panic(err);
  }
};
```

## Test ##
```sh
go test ./tests -v -bench=.
```

## Plot ##
If you wish to have this functionality you must run
```sh
go get github.com/gonum/plot
```
Plot tests. **[option]** is the name of the file/test plot.
```sh
sh rphash/plot [option]
```

For example, `sh rphash/plot kmeans`, will run rphash/plots/plot_kmeans.go.

# Developers #
+ Sam Wenke (**wenkesj**)
+ Jacob Franklin (**frankljbe**)
+ Sadiq Quasem (**quasemsm**)
