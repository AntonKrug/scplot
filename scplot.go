// +Build ignore
//go:generate go run -tags=dev assets_generate.go
package main

import (
	"io/ioutil"
	"log"

	"image"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"

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

func appMain(driver gxui.Driver) {

	p, err := plot.New()
	checkErr(err)

	groupA := plotter.XYs{{20, 35}, {30, 35}, {27, 10}}

	p.X.Label.Text = "index"
	p.Y.Label.Text = "value"

	err = plotutil.AddLinePoints(p, "", groupA)
	checkErr(err)

	width, height := 640, 480
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	c := vgimg.NewWith(vgimg.UseImage(m))
	p.Draw(draw.New(c))

	err = p.Save(640, 480, "buub.png")
	checkErr(err)

	theme := dark.CreateTheme(driver)
	img := theme.CreateImage()
	window := theme.CreateWindow(width, height, "Plot")
	texture := driver.CreateTexture(m, 1)
	img.SetTexture(texture)
	window.AddChild(img)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}
