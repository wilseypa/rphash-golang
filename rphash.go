package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/chrislusf/glow/driver"
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

// go run rphash ./dataset.csv ./results.txt
func main() {

	// Check command-line arguemnts
	if len(os.Args) <= 1 {
		fmt.Print("Invalid Input Arguemnts\n")
		return
	}

	gob.Register(Centroid{})
	gob.Register(itemset.Centroid{})
	gob.Register(utils.Hash64Set{})
	flag.Parse()

	t1 := time.Now()
	records := utils.ReadCSV(os.Args[1])

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

	normalizedResults := Stream.GetCentroids()
	ts := time.Since(t1)

	file, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, result := range normalizedResults {
		for _, dimension := range result {
			file.WriteString(fmt.Sprintf("%f ", parse.DeNormalize(dimension)))
		}
		file.WriteString("\n")
	}
	file.WriteString("Time: " + ts.String())
}
