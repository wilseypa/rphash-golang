<<<<<<< HEAD
# rphash-golang
Scalable Big Data Clustering by Random Projection Hashing: golang implementation
=======
# RPHash
[![Build Status](https://travis-ci.org/wilseypa/rphash-golang.svg)](https://travis-ci.org/wilseypa/rphash-golang) [![Release Status](https://img.shields.io/badge/version-1.0.0-blue.svg)](https://github.com/wilseypa/rphash-golang/releases) ![RPHash](https://github.com/wilseypa/rphash-golang/blob/master/rphash.png)

RPHash takes clustering and unsupervised learning problems and solves them in an embarrassingly parallel manner.

**Clustering** is a core concept in data analysis. Issues arise with scalability and dimensionality, ever changing environments and compatibility, insecure communications and data movement.

**The solution** is secure, reliable, and fast data for large-scale distributed systems.

# Random Projection Hash (RPHash)
The algorithm was created for maximizing parallel computation while providing scalability for large scale deployment. It's suitable for high dimensional data sets and is scalable and streamline.

![Overview](https://github.com/wilseypa/rphash-golang/blob/master/overview.png)

# Installing the API
Ensure you have **Go**, **git**, and **mercurial** installed on your system. Additionally, ensure that you have your Go environment setup.

```sh
go get github.com/wilseypa/rphash-golang
# or, clone from source
git clone https://github.com/wilseypa/rphash-golang.git
```

# Demo
A demo is presented that shows clustering the MNIST Dataset of 5,000 vectors representing digits that are 784 dimensions (28 x 28 grayscale image) in under 2 seconds using a MapReduce framework. **[See Demo](https://github.com/wilseypa/rphash-golang/blob/master/demo)**

# Test

```sh
go test ./tests -v -bench=.
```

# Developers
- Sam Wenke (**wenkesj**)
- Jacob Franklin (**frankljbe**)

# Documentation
- Sadiq Quasem (**quasemsm**)
>>>>>>> fc848e814b58c4eada1bae82a60fa994c2e262fb
