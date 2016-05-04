package utils

import (
    "bytes"
	"io"
	"os"
    "bufio"
    "strings"
	"strconv"
)
type DataFileReader struct {
  reader          	*bufio.Reader
  file				*os.File
  buffer			*bytes.Buffer
  hasNext			bool
  part				[]byte
  prefix			bool
  err 				error
}


func NewDataFileReader(path string) *DataFileReader {
	var (
		file   *os.File
		err    error
	)
	if file, err = os.Open(path); err != nil {
		panic(err)
	}
	
	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	hasNext := true
	return &DataFileReader{
		reader:       	   reader,
		buffer:            buffer,
		hasNext:           hasNext,
		file:			   file,
  }
}

func (this *DataFileReader) HasNext() bool {
	return this.hasNext
}

func (this *DataFileReader) Next() []float64 {
		for {
			if this.part, this.prefix, this.err = this.reader.ReadLine(); this.err != nil {
				this.hasNext = false
				return nil
			}
			this.buffer.Write(this.part)
			if !this.prefix {
				line := strings.Fields(this.buffer.String())
				this.buffer.Reset()
				return StringLineToFloatLine(line)
			}
		}
}

func ReadXLines(path string, x int) (lines [][]string, err error) {
	// Read a whole file into the memory and store it as array of lines
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for i := 0; i < x; i ++ {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			line := strings.Fields(buffer.String())
			lines = append(lines, line)
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func StringArrayToFloatArray(lines [][]string) (result [][]float64) {
    result = make([][]float64, len(lines), len(lines))
    for i, line := range lines {
        result[i] = make([]float64, len(lines[i]), len(lines[i]))
        for j, toFloat := range line {
			float, err := strconv.ParseFloat(toFloat, 64)
			if err != nil {
				panic(err)
			}
           	result[i][j] = float
        }
    }
    return result;
}

func StringLineToFloatLine(line []string) (result []float64) {
	    result = make([]float64, len(line), len(line))
        for j, toFloat := range line {
			float, err := strconv.ParseFloat(toFloat, 64)
			if err != nil {
				panic(err)
			}
           	result[j] = float
        }
		return result;
}

func ReadLines(path string) (lines [][]string, err error) {
	// Read a whole file into the memory and store it as array of lines
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			line := strings.Fields(buffer.String())
			lines = append(lines, line)
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}