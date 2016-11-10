package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/wilseypa/rphash-golang/api"
)

// go run rphash ./dataset.csv ./results.txt
func main() {

	// Check command-line arguemnts
	if len(os.Args) < 3 {
		fmt.Print("Invalid Input Arguemnts\n")
		// TODO - print correct usage.
		return
	}

	// Process input arguemnts
	// TODO - revise.  Simple for now.
	isDistributed := false
	clusterNodeCount := 0
	if len(os.Args) > 3 {
		isDistributed = true
		clusterNodeCount, _ = strconv.Atoi(os.Args[3])
	}

	// Keep track of timing for performance metrics
	t1 := time.Now()

	// Perform the clustering, either in distributed or local form
	// depending on the input arguments
	normalizedResults := api.ClusterFile(os.Args[1], 6, isDistributed, clusterNodeCount)

	// Determine the elapsed time
	ts := time.Since(t1)

	// Remove the output file if it exists already
	os.Remove(os.Args[2])

	// Write the results to the file
	file, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, result := range normalizedResults {
		result = result[0:(len(result) - 1)]
		for indxR := 0; indxR < len(result); indxR++ {
			dimension := result[indxR]
			if indxR < (len(result) - 1) {
				file.WriteString(fmt.Sprintf("%f,", api.Denormalize(dimension)))
			} else {
				file.WriteString(fmt.Sprintf("%f", api.Denormalize(dimension)))
			}
		}
		file.WriteString("\n")
	}

	// Print the timing metrics to the screen/terminal
	file.WriteString(ts.String())
	fmt.Println("Time: " + ts.String())
}
