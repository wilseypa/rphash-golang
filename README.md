# Scalable Big Data Clustering by Random Projection Hashing #
[![Build Status](https://travis-ci.org/wilseypa/rphash-golang.svg)](https://travis-ci.org/wilseypa/rphash-golang)
[![Release Status](https://img.shields.io/badge/version-1.0.0-blue.svg)](https://github.com/wilseypa/rphash-golang/releases)

The goal is to create a simple, secure, distributed, scalable, and parallel clustering algorithm to be used on almost any system.

**Clustering** is a core concept in data analysis. Issues arise with scalability and dimensionality, ever changing environments and compatibility, insecure communications and data movement

**The solution** is secure, reliable, and fast data for large-scale distributed systems.

**The Algorithm** provides more accurate clusters and an inherently distributed system.

![Clusters](https://github.com/wilseypa/rphash-golang/blob/master/clusters.png)

**Random Projection Hash (RPHash)** has been created for maximizing parallel computation
while providing scalability for large scale deployment. It's suitable for high dimensional data sets and is
scalable and streamline.

![Overview](https://github.com/wilseypa/rphash-golang/blob/master/overview.png)

# Installing #
Ensure you have **Go**, **git**, and **mercurial** installed on your system. Additionally, ensure that you have your Go environment setup.
```sh
go get github.com/wilseypa/rphash-golang
```

# Demo #
Clustering the MNIST Dataset of 5,000 vectors representing digits that are 784 dimensions (28 x 28 grayscale image).
**[See Demo](https://github.com/wilseypa/rphash-golang/blob/master/demo)**

# Test #
```sh
go test ./tests -v -bench=.
```

# Developers #
+ Sam Wenke (**wenkesj**)
+ Jacob Franklin (**frankljbe**)
+ Sadiq Quasem (**quasemsm**)
