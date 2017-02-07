package stream

import (
	"math/rand"
	"time"

	"github.com/wilseypa/rphash-golang/defaults"
	"github.com/wilseypa/rphash-golang/itemset"
	"github.com/wilseypa/rphash-golang/types"
)

type SimplifiedStream struct {
	K         int
	Counts    []int64
	Centroids [][]float64
}

type Stream struct {
	processedCount      int
	vectorCount         int
	counts              []int64
	centroids           [][]float64
	variance            float64
	CentroidCounter     types.CentroidItemSet
	randomSeedGenerator *rand.Rand
	rphashObject        types.RPHashObject
	lshGroup            []types.LSH
	decoder             types.Decoder
	projector           types.Projector
	hash                types.Hash
	varianceTracker     types.StatTest
	lshChannel          chan *itemset.Centroid
	RunCount            int
}

func NewStream(rphashObject types.RPHashObject) *Stream {
	randomSeedGenerator := rand.New(rand.NewSource(rphashObject.GetRandomSeed()))
	hash := defaults.NewHash(rphashObject.GetHashModulus())
	decoder := rphashObject.GetDecoderType()
	varianceTracker := defaults.NewStatTest(0.01)
	projections := rphashObject.GetNumberOfProjections()
	k := rphashObject.GetK() * projections
	CentroidCounter := defaults.NewCentroidCounter(k)
	lshGroup := make([]types.LSH, projections)
	lshChannel := make(chan *itemset.Centroid, 10000)
	var projector types.Projector
	for i := 0; i < projections; i++ {
		seed := time.Now().UnixNano()
		projector = defaults.NewProjector(rphashObject.GetDimensions(), decoder.GetDimensionality(), seed)
		lshGroup[i] = defaults.NewLSH(hash, decoder, projector)
	}
	return &Stream{
		counts:              nil,
		centroids:           nil,
		variance:            0,
		processedCount:      0,
		vectorCount:         0,
		CentroidCounter:     CentroidCounter,
		randomSeedGenerator: randomSeedGenerator,
		rphashObject:        rphashObject,
		lshGroup:            lshGroup,
		hash:                hash,
		decoder:             decoder,
		projector:           projector,
		varianceTracker:     varianceTracker,
		lshChannel:          lshChannel,
	}
}

func (this *Stream) AddVectorOnlineStep(vec []float64) *itemset.Centroid {
	c := itemset.NewCentroidStream(vec)
	tmpvar := this.varianceTracker.UpdateVarianceSample(vec)

	if this.variance != tmpvar {
		for _, lsh := range this.lshGroup {
			lsh.UpdateDecoderVariance(tmpvar)
		}
		this.variance = tmpvar
	}

	for _, lsh := range this.lshGroup {
		hash := lsh.LSHHashStream(vec, this.rphashObject.GetNumberOfBlurs())

		for _, h := range hash {
			c.AddID(h)
		}
	}
	return c
}

func (this *Stream) ProcessCentroids() {
	if this.centroids == nil {
		this.Run()
	}
}

func (this *Stream) GetLshCentroids() [][]float64 {
	var centroids [][]float64
	for _, cent := range this.CentroidCounter.GetTop() {
		centroids = append(centroids, cent.Centroid())
	}
	return centroids
}

func (this *Stream) GetLshCounts() []int64 {
	return this.CentroidCounter.GetCounts()
}

func (this *Stream) GetCentroids() [][]float64 {
	if this.centroids == nil {
		this.Run()
		var centroids [][]float64
		for _, cent := range this.CentroidCounter.GetTop() {
			centroids = append(centroids, cent.Centroid())
		}
		this.centroids = defaults.NewKMeansWeighted(this.rphashObject.GetK(), centroids, this.CentroidCounter.GetCounts()).GetCentroids()
	}
	return this.centroids
}

func (this *Stream) GetWeightedKMeans(centroids [][]float64) [][]float64 {
	if this.centroids == nil {
		this.centroids = defaults.NewKMeansWeighted(this.rphashObject.GetK(), centroids, this.CentroidCounter.GetCounts()).GetCentroids()
	}
	return this.centroids
}

func (this *Stream) AppendVector(vector []float64) {
	//JF this check is required to stop from overflowing memory in the lshChannel with very large data sets.
	if (this.vectorCount - this.processedCount) > 10000 {
		this.Run()
	}
	this.vectorCount++
	go func(vector []float64) {
		this.lshChannel <- this.AddVectorOnlineStep(vector)
		return
	}(vector)
}

func (this *Stream) Run() {
	for this.processedCount < this.vectorCount {
		cent := <-this.lshChannel
		this.CentroidCounter.Add(cent)
		this.processedCount++
	}
}

func (this *Stream) GetSimplifiedStream() SimplifiedStream {
	this.ProcessCentroids()
	return SimplifiedStream{this.rphashObject.GetK(), this.GetLshCounts(), this.GetLshCentroids()}
}

func doVectsMatch(vect1 []float64, vect2 []float64) bool {
	for i, val1 := range vect1 {
		if val1 != vect2[i] {
			return false
		}
	}
	return true
}

func (stream1 *SimplifiedStream) CombineSimpleStreams(stream2 SimplifiedStream) SimplifiedStream {
	newCounts := make([]int64, len(stream1.Counts), len(stream1.Counts)+len(stream2.Counts))
	newCents := make([][]float64, len(stream1.Counts), len(stream1.Centroids)+len(stream2.Centroids))

	for i, count := range stream1.Counts {
		newCounts[i] = count
	}

	for i, cent := range stream1.Centroids {
		newCents[i] = cent
	}

	for _, cent2 := range stream2.Centroids {
		added := false
		for i, cent1 := range stream1.Centroids {
			if doVectsMatch(cent1, cent2) {
				newCounts[i] = newCounts[i] + 1
				added = true
			}
		}
		if !added {
			newCounts = newCounts[:len(newCounts)+1]
			newCents = newCents[:len(newCents)+1]
			newCounts[len(newCounts)-1] = 1
			newCents[len(newCents)-1] = cent2
		}
	}

	return SimplifiedStream{stream1.K, newCounts, newCents}
}
