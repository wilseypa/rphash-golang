package api

import (
	"encoding/gob"
	"flag"
	"sync"

	"github.com/chrislusf/glow/flow"
	"github.com/wilseypa/rphash-golang/itemset"
	"github.com/wilseypa/rphash-golang/parse"
	"github.com/wilseypa/rphash-golang/reader"
	"github.com/wilseypa/rphash-golang/stream"
	"github.com/wilseypa/rphash-golang/utils"
)

var (
	f                  = flow.New()
	expectedDimensions = -1
	numClusters        = 6
)

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

func ClusterFile(filename string) [][]float64 {
	data := utils.ReadCSV(filename)
	return Cluster(data)
}

func Cluster(records [][]float64) [][]float64 {

	gob.Register(Centroid{})
	gob.Register(itemset.Centroid{})
	gob.Register(utils.Hash64Set{})
	flag.Parse()

	Object := reader.NewStreamObject(len(records[0]), numClusters)
	Stream := stream.NewStream(Object)

	outChannel := make(chan Centroid)

	ch := make(chan []float64)

	source := f.Channel(ch)

	f1 := source.Map(func(record []float64) Centroid {
		return Centroid{C: Stream.AddVectorOnlineStep(record)}
	}).AddOutput(outChannel)

	flow.Ready()

	var wg sync.WaitGroup

	goStart(&wg, func() {
		f1.Run()
	})

	goStart(&wg, func() {
		for out := range outChannel {
			Stream.CentroidCounter.Add(out.C)
		}
	})

	for _, record := range records {
		ch <- record
	}

	close(ch)
	wg.Wait()

	return Stream.GetCentroids()
}

func Denormalize(dimension float64) float64 {
	return parse.DeNormalize(dimension)
}
