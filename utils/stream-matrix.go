package utils

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

type StreamMatrix struct {
	reader      *bufio.Reader
	file        *os.File
	buffer      *bytes.Buffer
	sizeChecked bool
	prevLine    string
	sizeVal     int
	filename    string
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
	return &StreamMatrix{readerM, fileM, bufferM, false, "", -1, path}
}

func (this *StreamMatrix) GetDataSetSize() int {
	fileTmp, _ := os.Open(this.filename)
	size, _ := lineCounter(fileTmp)
	return size
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func (this *StreamMatrix) GetVectSize() int {
	if this.sizeVal < 0 {
		this.sizeVal = len(this.GetNextVector())
		this.sizeChecked = true
	}
	return this.sizeVal
}

func (this *StreamMatrix) GetNextVector() string {

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
		return ""
	}

	// Write the next line into the buffer
	this.buffer.Write(part)
	if !prefix {
		line := strings.Fields(this.buffer.String())
		this.buffer.Reset()

		// Convert the line to floats and return the vector
		this.prevLine = line[0]
		return this.prevLine
	}

	// Return nil if we get here (something went wrong)
	return this.GetNextVector()
}
