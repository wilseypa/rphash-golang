package utils

import (
  "github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
	"strconv"
)

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
