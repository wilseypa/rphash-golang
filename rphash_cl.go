package main

import (
	"fmt"
	"os"
	"time"

	"github.com/wilseypa/rphash-golang/api"
)

// go run rphash ./dataset.csv ./results.txt
func main() {

	// Check command-line arguemnts
	if len(os.Args) <= 1 {
		fmt.Print("Invalid Input Arguemnts\n")
		return
	}

	t1 := time.Now()
	normalizedResults := api.ClusterFile(os.Args[1])
	ts := time.Since(t1)

	file, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, result := range normalizedResults {
		for _, dimension := range result {
			file.WriteString(fmt.Sprintf("%f ", api.Denormalize(dimension)))
		}
		file.WriteString("\n")
	}
	file.WriteString("Time: " + ts.String())
}
