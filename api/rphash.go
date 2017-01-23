package api

import (
	"encoding/gob"
	"flag"
	"sync"

	"github.com/chrislusf/glow/flow"
	"github.com/wilseypa/rphash-golang/itemset"
	"github.com/wilseypa/rphash-golang/reader"
	"github.com/wilseypa/rphash-golang/stream"
	"github.com/wilseypa/rphash-golang/utils"
)

type VectorStream struct {
	wg         *sync.WaitGroup
	ch         chan ([][]float64)
	Stream     *stream.Stream
	f1         *flow.Dataset
	outChannel chan []Centroid
}

type Centroid struct {
	C *itemset.Centroid
}

func goStart(wg *sync.WaitGroup, fn func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fn()
	}()
}

func ClusterFile(filename string, numClusters int, distributed bool, clusters int) [][]float64 {
	//data, _ := utils.ReadCsvWithClassif(filename)
	data := utils.ReadCsvStreaming(filename)
	vectStream := getVectStreamFlow(data.GetVectSize(), numClusters)

	// Driver code here

	startFlow(vectStream)
	feedCluster(data, vectStream)

	closeFlow(vectStream)
	return vectStream.Stream.GetCentroids()
}

func getVectStreamFlow(dimension int, numClusters int) VectorStream {
	ch := make(chan [][]float64)
	var wg sync.WaitGroup

	gob.Register(Centroid{})
	gob.Register(itemset.Centroid{})
	gob.Register(utils.Hash64Set{})
	flag.Parse()

	Object := reader.NewStreamObject(dimension, numClusters)
	Stream := stream.NewStream(Object)

	outChannel := make(chan []Centroid)

	f := flow.New()
	source := f.Channel(ch)

	f1 := source.Map(func(records [][]float64) []Centroid {
		centroidArr := make([]Centroid, len(records))
		for indx, record := range records {
			centroidArr[indx] = Centroid{C: Stream.AddVectorOnlineStep(record)}
		}
		return centroidArr
	}).AddOutput(outChannel)

	flow.Ready()

	return VectorStream{&wg, ch, Stream, f1, outChannel}
}

func startFlow(vectStream VectorStream) {
	goStart(vectStream.wg, func() {
		vectStream.f1.Run()
	})

	goStart(vectStream.wg, func() {
		for out := range vectStream.outChannel {
			for _, centroidVal := range out {
				vectStream.Stream.CentroidCounter.Add(centroidVal.C)
			}
		}
	})
}

func feedCluster(dataMatrix *utils.StreamMatrix, vectStream VectorStream) {
	/*
		for _, record := range data {
			rTmp := make([][]float64, 1)
			rTmp[0] = record
			vectStream.ch <- rTmp
		}*/
	// Assign the data chunks to channels.

	// Define metadata for dividing up the data
	clusters := 4
	vectChunks := 10
	size := dataMatrix.GetDataSetSize()
	divisions := clusters * vectChunks

	// Manage data dispatch to compute nodes
	for i := 0; i < divisions; i++ {

		// Determine the bounds for each chunk
		start := i * vectChunks
		end := (i + 1) * vectChunks
		if end >= size {
			end = size - 1
		}

		// Assign the chunk to the corresponding compute node
		// computeNode := start % clusters
		if end > start {
			//currSlice := data[start:end][:]
			tmpSize := end - start
			currSlice := make([][]float64, tmpSize)
			for indx := 0; indx < tmpSize; indx++ {
				currSlice[indx] = dataMatrix.GetNextVector()
			}
			vectStream.ch <- currSlice
		}
	}
}

func closeFlow(vectStream VectorStream) {
	close(vectStream.ch)
	vectStream.wg.Wait()
}

func Denormalize(dimension float64) float64 {
	//return parse.DeNormalize(dimension)
	return dimension
}
