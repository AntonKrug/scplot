// +Build ignore
//go:generate go run -tags=dev assets_generate.go
package main

import (
	"io/ioutil"
	"log"

	"image"
	"image/color"
	"image/draw"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

type ead_file struct {
	filename string
	metadata string
}

var ead_files []ead_file
var dictionary map[string]string

func checkErr(err error) {
	if err != nil {
		// log.Fatal(au.Red(err))
		log.Fatal(err)
	}
}

func readFileRealRawToString(filename string) string {
	content, err := ioutil.ReadFile(filename)
	checkErr(err)
	return string(content)
}

func generatePlot() {
	p, err := plot.New()
	checkErr(err)

	groupA := plotter.XYs{{20, 35}, {30, 35}, {27, 10}}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p, "", randomPoints(15))

	checkErr(err)

	// Save the plot to a PNG file.
	err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png")

	checkerr(err)
}

func appMain(driver gxui.Driver) {
	width, height := 640, 480
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	theme := dark.CreateTheme(driver)
	img := theme.CreateImage()
	window := theme.CreateWindow(width, height, "Plot")
	texture := driver.CreateTexture(m, 1.0)
	img.SetTexture(texture)
	window.AddChild(img)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}
