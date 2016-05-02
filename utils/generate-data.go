package utils

import (
	"os"
    "bufio"
    "math/rand"
	"strconv"
)

func GenerateData(path string, dimensionality int, numRows int) () {
    file, err := os.Create(path)
    defer file.Close();
    if err != nil {
        panic(err);
    }
    w := bufio.NewWriter(file)
    for i := 0; i < numRows; i++ {
        for j := 0; j < dimensionality; j++ {
            w.WriteString(strconv.FormatFloat(rand.Float64(), 'f', -1, 64))
            w.WriteString(" ")
        }
        w.WriteString("\n")
    }
    w.Flush()
}