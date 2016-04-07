# RPHash – Clustering the MNIST Dataset
This demo clusters the MNIST dataset into 10 clusters, one centroid for each digit classifier.


## How It Works

The dataset is clustered with the RPHash algorithm. The input vectors representing digits are mapped to the clusterer and then centroids are found.
The centroids are then plotted in the `demo/plots` directory where `centroid-drawing-#` are 28x28 plots of the centroids and `centroid-dimensions-#` is a plot of the strength of each dimension of the centroid vector.

To run the demo,

```sh
go run main.go
```

The output will look like this,

```sh
# Benchmark output after clustering...
# Time it takes to cluster 5,000 vectors with 784 dimensions.
2016/04/07 13:31:09 Time: 2.358581704s
```

```sh
# locaiton of the plots
demo/plots
# location of the data
demo/data
```
