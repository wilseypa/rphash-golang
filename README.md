# Scalable Big Data Clustering by Random Projection Hashing #
[![Build Status](https://travis-ci.org/wenkesj/rphash.svg)](https://travis-ci.org/wenkesj/rphash)
[![Release Status](https://img.shields.io/badge/version-1.0.0-blue.svg)](https://github.com/wenkesj/rphash/releases)

**Clustering** is a core concept in data analysis. Issues arise with scalability and dimensionality, ever changing environments and compatibility, insecure communications and data movement

**The solution** is secure, reliable, and fast data for large- scale distributed systems.


**The Algorithm** provides more accurate clusters and an inherently distributed system.

![Clusters](https://github.com/wenkesj/rphash/blob/master/clusters.png)

**Random Projection Hash (RPHash)** has been created for maximizing parallel computation
while providing scalability for large scale deployment. It's suitable for high dimensional data sets and is
scalable and streamline.

![Overview](https://github.com/wenkesj/rphash/blob/master/overview.png)

# Table of contents #
+ **[Installing and Testing](https://github.com/wenkesj/rphash#installing-and-testing)**
+ **[API](https://github.com/wenkesj/rphash#api)**
+ **[Developers](https://github.com/wenkesj/rphash#developers)**
+ **[Pull requests welcome](https://github.com/wenkesj/rphash/blob/master/TODO.md)**
+ **[Learn more](https://github.com/wenkesj/rphash/blob/master/REFERENCES.md)**
+ **[Versioning and updates](https://github.com/wenkesj/rphash/blob/master/CHANGELOG.md)**

# Installing, Testing, and Plotting #
```sh
git clone --depth=50 --branch=master https://github.com/wenkesj/rphash.git wenkesj/rphash
cd wenkesj/rphash
export GOPATH=$HOME/<your-gopath>
export PATH=$HOME/<your-gopath>/bin:$PATH
go get -t -v ./...
sh install
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
