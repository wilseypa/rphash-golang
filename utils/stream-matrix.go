package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type StreamMatrix struct {
	reader      *bufio.Reader
	file        *os.File
	buffer      *bytes.Buffer
	sizeChecked bool
	prevLine    []float64
}

func SetupStreamMatrix(path string) *StreamMatrix {

	// Setup variables
	var (
		fileM *os.File
		err   error
	)

	// Open the file
	if fileM, err = os.Open(path); err != nil {
		panic("Unable to Open File: " + path)
	}
	//defer fileM.Close()

	// Assign class variables
	readerM := bufio.NewReader(fileM)
	bufferM := bytes.NewBuffer(make([]byte, 0))

	// Setup the class and return it
	return &StreamMatrix{readerM, fileM, bufferM, false, nil}
}

func (this *StreamMatrix) GetVectSize() int {
	size := len(this.GetNextVector())
	this.sizeChecked = true
	return size
}

func (this *StreamMatrix) GetNextVector() []float64 {
	fmt.Println("START")

	// Return the previous vector, if the size was checked
	if this.sizeChecked {
		this.sizeChecked = false
		return this.prevLine
	}

	// Setup variables
	var (
		part   []byte
		prefix bool
		err    error
	)

	// Read the next line in the file
	if part, prefix, err = this.reader.ReadLine(); err != nil {

		// Return nil if the file end has been reached
		if err == io.EOF {
			err = nil
		}
		fmt.Println("ERR")
		return nil
	}

	// Write the next line into the buffer
	this.buffer.Write(part)
	if !prefix {
		line := strings.Fields(this.buffer.String())
		this.buffer.Reset()

		// Convert the line to floats and return the vector
		this.prevLine = StringLineToFloatLine(line)
		fmt.Println(line)
		fmt.Println("END")
		return this.prevLine
	}

	// Return nil if we get here (something went wrong)
	fmt.Println("EXI")
	return nil
}
