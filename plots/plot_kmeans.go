package main

import (
  "github.com/gonum/plot"
  "github.com/gonum/plot/plotter"
  "github.com/gonum/plot/plotutil"
  "github.com/gonum/plot/vg"
  "github.com/wenkesj/rphash/clusterer"
  "math/rand"
)

func main() {
  var numClusters = 5
  var numDataPoints = 100
  var dimensionality = 2

  // Randomly produce the data.
  data := make([][]float64, numDataPoints)
  dataPoints := make(plotter.XYs, numDataPoints)
  for i := 0; i < numDataPoints; i++ {
    data[i] = make([]float64, dimensionality)
    for j := 0; j < dimensionality; j++ {
      data[i][j], dataPoints[i].X, dataPoints[i].Y = rand.Float64(), rand.Float64(), rand.Float64()
    }
  }

  // Run the clusterer on the data points.
  clusterer := clusterer.NewKMeansSimple(numClusters, data)
  clusterer.Run()
  result := clusterer.GetCentroids()

  // group the points into the dimensional plot.
  resultPoints := make(plotter.XYs, numClusters)
  for i := 0; i < numClusters; i++ {
    resultPoints[i].X, resultPoints[i].Y = result[i][0], result[i][1]
  }

  p, err := plot.New()
  if err != nil {
    panic(err)
  }

  p.Title.Text = "KMeans"
  p.X.Label.Text = "X"
  p.Y.Label.Text = "Y"

  err = plotutil.AddScatters(p,
    "Data", dataPoints,
    "Clusters", resultPoints)

  if err != nil {
    panic(err)
  }

  if err := p.Save(4*vg.Inch, 4*vg.Inch, "kmeans.png"); err != nil {
    panic(err)
  }
}
