package utils

import (
  "image/color"
  "github.com/gonum/matrix/mat64"
  "github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
  "github.com/gonum/plot/palette"
	"github.com/gonum/plot/vg"
	"strconv"
  "math"
)

var (
  grayscale = Grayscale{Min:0, Max:1}
)

// Linearly maps the colors between light gray and black
type Grayscale struct {
	Min float64
	Max float64
}

func (c *Grayscale) Color(z float64) color.Color {
	val := (1 - (z-c.Min)/(c.Max-c.Min))
	if val < 0 {
		val = 0
	} else if val > 1 {
		val = 1
	}
	val *= 255 * 0.9
	u8v := uint8(val)
	return color.RGBA{u8v, u8v, u8v, 255}
}

// Util for generating plots
func GeneratePlots(x, y [][]float64, title, xLabel, yLabel, fileName string, legendLabel []string) {
	outPlotPoints := make([]plotter.XYs, len(x))
	outPlots := make([]*plot.Plot, len(x))

	for i, _ := range outPlotPoints {
		outPlot, err := plot.New()
		outPlots[i] = outPlot
		outPlots[i].Title.Text = title
		outPlots[i].X.Label.Text = xLabel
		outPlots[i].Y.Label.Text = yLabel
		outPlotPoints[i] = make(plotter.XYs, len(x[0]))
		for j, _ := range x[0] {
			outPlotPoints[i][j].X = x[i][j]
			outPlotPoints[i][j].Y = y[i][j]
		}
		err = plotutil.AddLines(outPlots[i],
			legendLabel[i], outPlotPoints[i])
		if err != nil {
			panic(err)
		}

		if err = outPlot.Save(6*vg.Inch, 6*vg.Inch, (fileName+strconv.FormatInt(int64(i), 16))+".png"); err != nil {
			panic(err)
		}
	}
}

type Grid struct {
  Matrix *mat64.Dense
}

func NewGrid(matrix *mat64.Dense) *Grid {
  return &Grid{
    Matrix: matrix,
  }
}

// Dims returns the dimensions of the grid.
func (this *Grid) Dims() (c, r int) {
  return this.Matrix.Dims()
}

// Z returns the value of a grid value at (c, r).
// It will panic if c or r are out of bounds for the grid.
func (this *Grid) Z(c, r int) float64 {
  return this.Matrix.ColView(c).At(r, 0)
}

// X returns the coordinate for the column at the index x.
// It will panic if c is out of bounds for the grid.
func (this *Grid) X(c int) float64 {
  return this.Matrix.ColView(c).At(c, 0)
}

// Y returns the coordinate for the row at the index r.
// It will panic if r is out of bounds for the grid.
func (this *Grid) Y(r int) float64 {
  return this.Matrix.RowView(r).At(r, 0)
}

func MaxFloat(collection []float64) float64 {
  max := collection[0]
  for _, value := range collection {
    if value > max {
      max = value
    }
  }
  return max
}

func HeatMap(image []float64, index int) {
  dim := int(math.Sqrt(float64(len(image))))

  m := NewGrid(mat64.NewDense(dim, dim, image))

  h := plotter.NewHeatMap(m, palette.Heat(10, 1))

  p, err := plot.New()
  if err != nil {
    panic(err)
  }
  p.Title.Text = "Heat map"
  p.Y.Max = MaxFloat(image)
  p.X.Max = MaxFloat(image)
  p.Add(h)

  err = p.Save(6*vg.Inch, 6*vg.Inch, "plots/heatmap"+ strconv.FormatInt(int64(index), 16) +".png")
  if err != nil {
    panic(err)
  }
}

// 784 Bits
func Paint(image []float64, imageId int, fileName string) {
	outPlotPoints := make(plotter.XYs, len(image))
	outPlot, err := plot.New()
	if err != nil {
		panic(err)
	}
	x := 0
  y := 0
	for i, bit := range image {
		outPlotPoints[i].X = float64(x)

    if bit > 0.3 {
      outPlotPoints[i].Y = float64(y)
    } else {
      outPlotPoints[i].Y = 0
    }

		if i%int(math.Sqrt(float64(len(image)))) == 0 {
			x = 0
      y++
		} else {
			x++
		}
	}
  outPlot.Add(plotter.NewGrid())
  s, _ := plotter.NewScatter(outPlotPoints)
	outPlot.Add(s)
	if err = outPlot.Save(6*vg.Inch, 6*vg.Inch, fileName+strconv.FormatInt(int64(imageId), 16)+".png"); err != nil {
		panic(err)
	}
}
