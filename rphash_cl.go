package main

import (
	"container/list"
	"encoding/gob"
	"flag"
	"fmt"
	"os"
	"strconv"
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
	numClusters        = flag.Int("k", 0, "Number of Clusters")
	dataDimen          = flag.Int("dim", 0, "Dimension of Dataset")
	dataPath           = flag.String("data", "", "Data File")
	dataResults        = flag.String("output", "Test/output.rphash", "Results File")
	dataTimeResults    = flag.String("time", "Test/output.rphash.time", "Timing Results")
	paritions          = flag.Int("div", 1, "Number of Partitions")
	vectChunks         = flag.Int("vects", 100, "Number of Vectors in a Chunk")
)

type Centroid struct {
	C *itemset.Centroid
}

type CentroidData struct {
	dimensions int
	matrixList *list.List
}

type CountsData struct {
	arrayList *list.List
}

type LshCentObj struct {
	Cents  [][]float64
	Counts []int64
}

type DataInput struct {
	Key string
	Val [][]float64
}

type FloatHolder struct {
	Val []float64
}

type CentroidArr struct {
	Val []Centroid
}

type CentroidMat struct {
	Val [][]Centroid
}

func init() {
	gob.Register(itemset.Centroid{})
	gob.Register(utils.Hash64Set{})
	gob.Register(DataInput{})
	gob.Register(Centroid{})
	gob.Register(LshCentObj{})
	gob.Register(FloatHolder{})
	gob.Register(CentroidArr{})
	gob.Register(CentroidMat{})
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

// Function used to combine the data in the list (of the CentroidData structure)
// Into a single matrix (type: [][]float64).
func (this *CountsData) getUnderlyingArray() []int64 {

	// Determine the final dimensions of the matrixList
	totalSize := 0
	for e := this.arrayList.Front(); e != nil; e = e.Next() {
		totalSize += len(e.Value.([]int64))
	}

	// Allocate the new matrix
	newCounts := make([]int64, totalSize)
	//for i := range newMatrix {
	//	newMatrix[i] = make([]float64, this.dimensions)
	//}

	// Transfer over the list items to the matrix
	indx := 0
	for e := this.arrayList.Front(); e != nil; e = e.Next() {
		tmpMatrix := e.Value.([]int64)
		for i := 0; i < len(tmpMatrix); i++ {
			newCounts[indx] = tmpMatrix[i]
			indx++
		}
	}

	// Return the new marix
	return newCounts
}

func goStart(wg *sync.WaitGroup, fn func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fn()
	}()
}

func main() {
	flag.Parse()
	t1 := time.Now()
	/*
		Create Streams
		  -> Group vectors by key
			- Partition groups (by key)
			-> Add vector groups (online) to each stream (again by key)
			-> Combine/Merge Results?
			-> Perform KMeans on single resulting stream
	*/

	Object := reader.NewStreamObject(*dataDimen, *numClusters)
	Stream := stream.NewStream(Object)
	var mutex = &sync.Mutex{}

	outChannel := make(chan stream.SimplifiedStream)

	ch := make(chan DataInput)

	source := flow.New().Channel(ch)

	f1 := source.Map(func(inputStruct DataInput) (string, DataInput) {
		return inputStruct.Key, inputStruct
	}).Partition(*paritions).Map(func(key string, dataChunk DataInput) stream.SimplifiedStream {
		newCentroid := make([]Centroid, len(dataChunk.Val))
		for i, vect := range dataChunk.Val {
			mutex.Lock()
			newCentroid[i] = Centroid{C: Stream.AddVectorOnlineStep(vect)}
			mutex.Unlock()
		}
		for _, centSt := range newCentroid {
			mutex.Lock()
			Stream.CentroidCounter.Add(centSt.C)
			mutex.Unlock()
		}
		return Stream.GetSimplifiedStream()
	}).MergeReduce(func(stream1, stream2 stream.SimplifiedStream) stream.SimplifiedStream {
		return stream1.CombineSimpleStreams(stream2)
	}).AddOutput(outChannel)

	/*

		f1 := source.Map(func(inputStruct DataInput) (string, FloatHolder) {
			return inputStruct.Key, FloatHolder{inputStruct.Val}
		}).Partition(*paritions).Map(func(key string, dataChunk FloatHolder) (string, CentroidArr) {
			newCentroid := make([]Centroid, len(dataChunk.Val))
			for i, vect := range dataChunk.Val {
				mutex.Lock()
				newCentroid[i] = Centroid{C: Stream.AddVectorOnlineStep(vect)}
				mutex.Unlock()
			}
			return key, CentroidArr{newCentroid}
		}).GroupByKey().Map(func(key string, cents []CentroidArr) LshCentObj {
			for _, centL := range cents {
				for _, centSt := range centL.Val {
					mutex.Lock()
					Stream.CentroidCounter.Add(centSt.C)
					mutex.Unlock()
				}
			}
			Stream.ProcessCentroids()
			return LshCentObj{Stream.GetLshCentroids(), Stream.GetLshCounts()}
		}).AddOutput(outChannel)

	*/

	flow.Ready()

	var wg sync.WaitGroup

	goStart(&wg, func() {
		f1.Run()
	})

	// Extract the output from the output channel
	var finalCentroids [][]float64
	goStart(&wg, func() {

		// Create lists to store the result chunks in
		//initList := list.New()
		//lshTotCents := CentroidData{0, initList}
		//initList2 := list.New()
		//countList := CountsData{initList2}

		// Put all result chunks into a single list
		//for out := range outChannel {
		//	lshTotCents.matrixList.PushBack(out.Cents)
		//	countList.arrayList.PushBack(out.Counts)
		//}
		//finalCentroids = defaults.NewKMeansWeighted(*numClusters, lshTotCents.getUnderlyingMatrix(), countList.getUnderlyingArray()).GetCentroids()
		sStream := <-outChannel
		finalCentroids = defaults.NewKMeansWeighted(sStream.K, sStream.Centroids, sStream.Counts).GetCentroids()
	})

	//records := utils.ReadCsvStreaming(dataPath)
	feedCluster(ch)

	close(ch)
	wg.Wait()
	ts := time.Since(t1)

	// Write the results to the file
	file, err := os.OpenFile(*dataResults, os.O_WRONLY|os.O_CREATE, 0644)
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

func feedCluster(ch chan DataInput) {
	dataMatrix := utils.ReadCsvStreaming(*dataPath)

	// Define metadata for dividing up the data
	size := dataMatrix.GetDataSetSize()
	divisions := size / ((*paritions) * (*vectChunks))

	// Manage data dispatch to compute nodes
	curr := 0
	for i := 0; i < divisions; i++ {

		// Determine the bounds for each chunk
		start := i * (*vectChunks)
		end := (i + 1) * (*vectChunks)
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
			ch <- DataInput{strconv.Itoa(curr), currSlice}
			curr = (curr + 1) % *paritions
		}
	}
}
