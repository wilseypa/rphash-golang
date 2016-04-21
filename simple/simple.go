package simple

import (
  "github.com/wilseypa/rphash-golang/defaults"
  "github.com/wilseypa/rphash-golang/types"
  "math"
  "runtime"
)

type Simple struct {
  centroids    [][]float64
  variance     float64
  rphashObject types.RPHashObject
}

func NewSimple(_rphashObject types.RPHashObject) *Simple {
  return &Simple{
    variance:     0,
    centroids:    nil,
    rphashObject: _rphashObject,
  }
}

// Map is doing the count.
func (this *Simple) Map() *Simple {
  runtime.GOMAXPROCS(runtime.NumCPU())
  vecs := this.rphashObject.GetVectorIterator()
  //var hashResult int64;
  targetDimension := int(math.Floor(float64(this.rphashObject.GetDimensions() / 2)))
  numberOfRotations := 6
  numberOfSearches := 1
  vec := vecs.Next()
  hash := defaults.NewHash(this.rphashObject.GetHashModulus())
  decoder := defaults.NewDecoder(targetDimension, numberOfRotations, numberOfSearches)
  projector := defaults.NewProjector(this.rphashObject.GetDimensions(), decoder.GetDimensionality(), this.rphashObject.GetRandomSeed())
  LSH := defaults.NewLSH(hash, decoder, projector)
  // k := int(float64(this.rphashObject.GetK()) * math.Log(float64(this.rphashObject.GetK())));
  CountMinSketch := defaults.NewCountMinSketch(this.rphashObject.GetK())
  var vecCount = 0
  //1000 is an arbitrary comprise between speed and size should be tweeked later.
  hashChannel := make(chan int64, this.rphashObject.NumDataPoints())
  hashValues := make([]int64, this.rphashObject.NumDataPoints(), this.rphashObject.NumDataPoints())
  for vecs.HasNext() {
    go func(vec []float64, index int) {
      // Project the Vector to lower dimension.
      // Decode the new vector for meaningful integers
      // Hash the new vector into a 64 bit int.
      value := LSH.LSHHashSimple(vec)
      hashValues[index] = value
      hashChannel <- value
      //hashResult = LSH.LSHHashSimple(vec);
      // Add it to the count min sketch to update frequencies.
    }(vec, vecCount)
    vecCount++
    vec = vecs.Next()
  }
  vecs.StoreLSHValues(hashValues)
  //TODO should we Paralelize this? slowest loop but also have to wait for LSH Loops
  for i := 0; i < vecCount; i++ {
    hashResult := <-hashChannel
    CountMinSketch.Add(hashResult)
  }
  this.rphashObject.SetPreviousTopID(CountMinSketch.GetTop())
  vecs.Reset()
  return this
}

// Reduce is finding out where the centroids are in respect to the real data.
func (this *Simple) Reduce() *Simple {
  vecs := this.rphashObject.GetVectorIterator()
  if !vecs.HasNext() {
    return this
  }

  var centroids []types.Centroid
  for i := 0; i < this.rphashObject.GetK(); i++ {
    // Get the top centroids.
    previousTop := this.rphashObject.GetPreviousTopID()
    centroid := defaults.NewCentroidSimple(this.rphashObject.GetDimensions(), previousTop[i])
    centroids = append(centroids, centroid)
  }

  // Iterate over the dataset and check CountMinSketch.
  //Paralelize loop
  var centriodChannels []chan []float64
  for i, _ := range centroids {
    centriodChannels = append(centriodChannels, make(chan []float64, 10000))
    go func(id int) {
      for true {
        newVec, ok := <-centriodChannels[id]
        if !ok {
          return
        }
        centroids[id].UpdateVector(newVec)
      }
    }(i)
  }
  vec := vecs.Next()
  var hashResult = int64(0)
  for vecs.HasNext() {
    hashResult = vecs.PeakLSH()
    // For each vector, check if it is a centroid.
    for i, cent := range centroids {
      // Get an idea where the LSH is in respect to the vector.
      if cent.GetIDs().Contains(hashResult) {
        //centriodChannels[i] <- vec;
        centriodChannels[i] <- vec
        break
      }
    }
    vec = vecs.Next()
  }
  for _, channel := range centriodChannels {
    close(channel)
  }

  for _, cent := range centroids {
    this.rphashObject.AddCentroid(cent.Centroid())
  }

  vecs.Reset()
  return this
}

func (this *Simple) GetCentroids() [][]float64 {
  if this.centroids == nil {
    this.Run()
  }
  // Perform the KMeans on the centroids.
  result := defaults.NewKMeansSimple(this.rphashObject.GetK(), this.centroids).GetCentroids()
  return result
}

func (this *Simple) Run() {
  this.Map().Reduce()
  this.centroids = this.rphashObject.GetCentroids()
}

func (this *Simple) GetRPHash() types.RPHashObject {
  return this.rphashObject
}
