package main

import (
	"container/list"
	"encoding/gob"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/chrislusf/glow/driver"
	"github.com/chrislusf/glow/flow"
	"github.com/wilseypa/rphash-golang/defaults"
	"github.com/wilseypa/rphash-golang/itemset"
	"github.com/wilseypa/rphash-golang/reader"
	"github.com/wilseypa/rphash-golang/stream"
	"github.com/wilseypa/rphash-golang/utils"
)

var (
	expectedDimensions = -1
	numClusters        = 8
	dataDimen          = 22268
	dataPath           = "Data/DMMPLCN.csv"
	dataResults        = "Test/output.rphash"
	paritions          = 1
)

type Centroid struct {
	C *itemset.Centroid
}

type CentroidData struct {
	dimensions int
	matrixList *list.List
}

// Function used to combine the data in the list (of the CentroidData structure)
// Into a single matrix (type: [][]float64).
func (this *CentroidData) getUnderlyingMatrix() [][]float64 {

	// Determine the final dimensions of the matrixList
	totalSize := 0
	for e := this.matrixList.Front(); e != nil; e = e.Next() {
		totalSize += len(e.Value.([][]float64))
	}

	// Allocate the new matrix
	newMatrix := make([][]float64, totalSize)
	//for i := range newMatrix {
	//	newMatrix[i] = make([]float64, this.dimensions)
	//}

	// Transfer over the list items to the matrix
	indx := 0
	for e := this.matrixList.Front(); e != nil; e = e.Next() {
		tmpMatrix := e.Value.([][]float64)
		for i := 0; i < len(tmpMatrix); i++ {
			newMatrix[indx] = tmpMatrix[i]
			indx++
		}
	}

	// Return the new marix
	return newMatrix
}

func goStart(wg *sync.WaitGroup, fn func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fn()
	}()
}

func main() {
	gob.Register(Centroid{})
	gob.Register(itemset.Centroid{})
	gob.Register(utils.Hash64Set{})
	flag.Parse()

	t1 := time.Now()

	Object := reader.NewStreamObject(dataDimen, numClusters)
	Stream := stream.NewStream(Object)
	var mutex = &sync.Mutex{}

	outChannel := make(chan [][]float64)

	ch := make(chan [2]string)

	source := flow.New().Channel(ch)

	f1 := source.Map(func(strs [2]string) (string, string) {
		return strs[0], strs[1]
	}).Partition(paritions).Map(func(key string, line string) (string, []float64) {
		result := strings.Split(line, ",")
		floatLine := make([]float64, len(result)-1)
		for i, val := range result {
			if i < len(result)-1 {
				floatLine[i], _ = strconv.ParseFloat(val, 64)
			}
		}
		return key, floatLine
	}).Map(func(key string, record []float64) (string, Centroid) {
		mutex.Lock()
		newCentroid := Centroid{C: Stream.AddVectorOnlineStep(record)}
		mutex.Unlock()
		return key, newCentroid
	}).GroupByKey().Map(func(key string, cents []Centroid) [][]float64 {
		fmt.Println(len(cents))
		for _, centSt := range cents {
			mutex.Lock()
			Stream.CentroidCounter.Add(centSt.C)
			mutex.Unlock()
		}
		return Stream.GetCentroids()
	}).AddOutput(outChannel)

	flow.Ready()

	var wg sync.WaitGroup

	goStart(&wg, func() {
		f1.Run()
	})

	var finalCentroids [][]float64
	goStart(&wg, func() {
		initList := list.New()
		lshTotCents := CentroidData{0, initList}
		for out := range outChannel {
			lshTotCents.matrixList.PushBack(out)
		}
		finalCentroids = defaults.NewKMeansSimple(numClusters, lshTotCents.getUnderlyingMatrix()).GetCentroids()
	})

	//records := utils.ReadCsvStreaming(dataPath)
	feedCluster(ch)

	close(ch)
	wg.Wait()
	ts := time.Since(t1)

	// Write the results to the file
	file, err := os.OpenFile(dataResults, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, result := range finalCentroids {
		for indxR := 0; indxR < len(result); indxR++ {
			dimension := result[indxR]
			if indxR < (len(result) - 1) {
				file.WriteString(fmt.Sprintf("%f,", dimension)) //api.Denormalize(dimension)))
			} else {
				file.WriteString(fmt.Sprintf("%f", dimension)) //api.Denormalize(dimension)))
			}
		}
		file.WriteString("\n")
	}
	fmt.Println("Time: " + ts.String())
}

func feedCluster(ch chan [2]string) {
	dataMatrix := utils.ReadCsvStreaming(dataPath)
	curr := 0
	keepRunning := true
	for keepRunning {
		line := dataMatrix.GetNextVector()
		if len(line) > 0 {
			strArr := [2]string{strconv.Itoa(curr % paritions), line}
			ch <- strArr
			curr = curr + 1
		} else {
			keepRunning = false
		}
	}

	/*
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
			if end > start {
				tmpSize := end - start
				currSlice := make([][]float64, tmpSize)
				for indx := 0; indx < tmpSize; indx++ {
					currSlice[indx] = dataMatrix.GetNextVector()
				}
				ch <- currSlice
			}
		}
	*/
}

/*
f1 := source.Map(func(strs [2]string) (string, string) {
	return strs[0], strs[1]
}).Partition(paritions).Map(func(key string, line string) (string, []float64) {
	result := strings.Split(line, ",")
	floatLine := make([]float64, len(result))
	for i, val := range result {
		floatLine[i], _ = strconv.ParseFloat(val, 64)
	}
	return key, floatLine
}).Map(func(key string, record []float64) (string, float64) {
	fmt.Println(record)
	sum := 0.0
	for _, val := range record {
		sum = sum + val
	}
	return key, sum
}).GroupByKey().Map(func(key string, val []float64) Sum {
	fmt.Println(val)
	for _, v := range val {
		MainSum.addValue(v)
	}
	return MainSum
}).AddOutput(outChannel)
*/
