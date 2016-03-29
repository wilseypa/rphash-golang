package main

import (
  "bufio"
  "github.com/chrislusf/glow/flow"
  "github.com/wilseypa/rphash-golang/reader"
  "github.com/wilseypa/rphash-golang/stream"
  "github.com/wilseypa/rphash-golang/types"
  "github.com/wilseypa/rphash-golang/parse"
  "github.com/gonum/plot";
  "github.com/gonum/plot/plotter";
  "github.com/gonum/plot/plotutil";
  "github.com/gonum/plot/vg";
  //  _ "github.com/chrislusf/glow/driver"
  "log"
  "io"
  "os"
  "bytes"
  "strconv"
  "strings"
)

var (
  dataFilePath       = "gisette_train.data"
  f                  = flow.New()
  expectedDimensions = -1
  numClusters        = 10
  numShards          = 8
)

type Vector struct {
  Data []float64
}

func GeneratePlots(x, y [][]float64, title, xLabel, yLabel, fileName string, legendLabel []string) {
  outPlotPoints := make([]plotter.XYs, len(x));
  outPlots := make([]*plot.Plot, len(x))

  for i, _ := range outPlotPoints {
    outPlot, err := plot.New();
    outPlots[i] = outPlot
    outPlots[i].Title.Text = title;
    outPlots[i].X.Label.Text = xLabel;
    outPlots[i].Y.Label.Text = yLabel;
    outPlotPoints[i] = make(plotter.XYs, len(x[0]));
    for j, _ := range x[0] {
      outPlotPoints[i][j].X = x[i][j];
      outPlotPoints[i][j].Y = y[i][j];
    }
    err = plotutil.AddLines(outPlots[i],
      legendLabel[i], outPlotPoints[i]);
    if err != nil {
      panic(err);
    }

    if err = outPlot.Save(6 * vg.Inch, 6 * vg.Inch, (fileName + strconv.FormatInt(int64(i), 16)) + ".png"); err != nil {
      panic(err);
    }
  }
};

func ReadLines(path string) (lines [][]string, err error) {
  // Read a whole file into the memory and store it as array of lines
  var (
    file *os.File
    part []byte
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

func main() {
  var rphashObject *reader.StreamObject
  var rphashStream *stream.Stream
  var centroids []types.Centroid

  // Split the data into shards and send them to the Agents to work on.
  f.Source(func(out chan Vector) {
    log.Println("Reading Data File...")
    records, err := ReadLines(dataFilePath)
    log.Println("Finished Reading File with", len(records), "entries and", len(records[0]), "dimensions √")
    if err != nil {
      log.Println(err)
      os.Exit(1)
    }
    // Convert the record to standard floating points.
    log.Println("Parsing Data...")
    for i, record := range records {
      if i == 0 {
        // Create a new RPHash stream.
        rphashObject = reader.NewStreamObject(len(record), numClusters)
        rphashStream = stream.NewStream(rphashObject)
        rphashStream.RunCount = 1
      }
      data := make([]float64, len(record))
      for j, entry := range record {
        f, err := strconv.ParseFloat(entry, 64)
        f = parse.Normalize(f)
        if err != nil {
          log.Println(err)
          os.Exit(1)
        }
        data[j] = f
      }
      out <- Vector{Data:data}
    }
    log.Println("Finished Parsing Data √")
  }, numShards).Map(func(vec Vector) {
    log.Println(vec)
    centroids = append(centroids, rphashStream.AddVectorOnlineStep(vec.Data))
  }).Run()

  log.Println("Getting Best Centroid Candidates...")
  for _, cent := range centroids {
    rphashStream.CentroidCounter.Add(cent)
  }
  log.Println("Got Centroids √")
  log.Println("Denormalizing...")
  normalizedResults := rphashStream.GetCentroids()
  denormalizedResults := make([][]float64, len(normalizedResults))
  for i, result := range normalizedResults {
    row := make([]float64, len(result));
    for j, dimension := range result {
      row[j] = parse.DeNormalize(dimension)
    }
    denormalizedResults[i] = row
  }
  log.Println("Displaying Results ...")
  labels := make([]string, len(denormalizedResults))
  xPlotValues := make([][]float64, len(denormalizedResults))
  yPlotValues := make([][]float64, len(denormalizedResults))
  for i, result := range denormalizedResults {
    xPlotValues[i] = make([]float64, len(result))
    yPlotValues[i] = make([]float64, len(result))
    for j, val := range result {
      xPlotValues[i][j] = float64(j)
      yPlotValues[i][j] = val
    }
    sI := strconv.FormatInt(int64(i), 16)
    labels[i] = "Digit " + sI + " (by Classifier Centroid)"
  }
  GeneratePlots(xPlotValues, yPlotValues, "High Dimension Handwritting Digits 0-9 Classification", "Centroid", "Strength of Visual Pixel Recognition (0-1000)", "digits", labels)
}
