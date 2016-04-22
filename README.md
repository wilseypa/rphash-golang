# RPHash
[![Build Status](https://travis-ci.org/wilseypa/rphash-golang.svg)](https://travis-ci.org/wilseypa/rphash-golang) [![Release Status](https://img.shields.io/badge/version-1.0.0-blue.svg)](https://github.com/wilseypa/rphash-golang/releases) ![RPHash](https://github.com/wilseypa/rphash-golang/blob/master/rphash.png)

RPHash takes clustering and unsupervised learning problems and solves them in an embarrassingly parallel manner.

**Clustering** is a core concept in data analysis. Issues arise with scalability and dimensionality, ever changing environments and compatibility, insecure communications and data movement.

**The solution** is secure, reliable, and fast data for large-scale distributed systems.

# Random Projection Hash (RPHash)
The algorithm was created for maximizing parallel computation while providing scalability for large scale deployment. It's suitable for high dimensional data sets and is scalable and streamline.

![Overview](https://github.com/wilseypa/rphash-golang/blob/master/overview.png)

# Installing
Ensure you have **Go**, **git**, and **mercurial** installed on your system. Additionally, ensure that you have your Go environment setup.

```sh
go get github.com/wilseypa/rphash-golang
# or, clone from source
git clone https://github.com/wilseypa/rphash-golang.git
```

# API

```sh
rphash-golang stream                  # Streaming command for clustering
  --num.clusters <#>                  # Number of clusters -> output centroids
  --num.shards <#>                    # Number of shards on the data
  --local.file <filename>             # Filename to cluster
  --cluster <rphash|streaming-kmeans> # Cluster algorithm
  --centroid.plots                    # Enable plots
  --centroid.plots.file <filename>    # Output dimension plot path
  --centroid.paint <filename>         # Output of a NxN matrix (experimental)
  --centroid.heat <filename>          # Output of a 3D heatmap (experimental)
  --hdfs.enable                       # Enable hdfs
  --hdfs.dir                          # hdfs directory
  [glow flags]                        # All other glow flags
```

# Test

```sh
go test ./tests -v -bench=.
```

# Developers
- Sam Wenke (**wenkesj**)
- Jacob Franklin (**frankljbe**)

# Documentation
- Sadiq Quasem (**quasemsm**)
