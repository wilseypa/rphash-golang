package utils

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"github.com/wilseypa/rphash-golang/parse"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	fixedDecimalPoint = 18
	weightMax         = math.Abs(parse.ToFixed(math.MaxFloat64, fixedDecimalPoint))
	weightMin         = float64(0)
)

func NormalizeSlice(records [][]float64) [][]float64 {
	data := make([][]float64, len(records))
	for i, record := range records {
		data[i] = make([]float64, len(record))
		for j, entry := range record {
			data[i][j] = parse.Normalize(entry)
		}
	}
	return data
}

func ReadCSV(path string) [][]float64 {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	lines, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	lines = lines[1:]
	return NormalizeSlice(StringArrayToFloatArray(lines))
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
	for i := 0; i < x; i++ {
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
        continue
			}
			result[i][j] = float
		}
	}
	return result
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
