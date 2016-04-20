package utils

import (
    "bytes"
	"io"
	"os"
    "bufio"
    "strings"
	"strconv"
)

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