// +Build ignore
//go:generate go run -tags=dev assets_generate.go
package main

import (
	"flag"
	"github.com/logrusorgru/aurora"
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	// ndraw "gonum.org/v1/plot/vg/draw"
	// "gonum.org/v1/plot/vg/vgimg"
	"gonum.org/v1/plot/vg"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

var timestampFlag = flag.Bool("timestamped_log", true, "enable logger's timestamps")
var colorsFlag = flag.Bool("colors", false, "enable logger's colors (disabled by default)")
var inputFlag = flag.String("input", "", "raw dump file of a variable to be read as input")
var scaleFlag = flag.Float64("scale", 3.0, "Maximum scaling factor used (if smaller window used, it will scale down to fit anyway)")
var widthFlag = flag.Int("width", 640, "Render width resolution for the plot")
var heightFlag = flag.Int("height", 480, "Render height resolution for the plot")

var sourceWithoutExtension string
var pngFilename string

var au aurora.Aurora

func fileWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(au.Red(err))
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

	// m := image.NewRGBA(image.Rect(0, 0, width, height))
	// c := vgimg.NewWith(vgimg.UseImage(m))
	// p.Draw(ndraw.New(c))

	err = p.Save(vg.Length(*widthFlag), vg.Length(*heightFlag), pngFilename)
	checkErr(err)

	f, err := os.Open(pngFilename)
	checkErr(err)

	source, _, err := image.Decode(f)
	checkErr(err)

	theme := dark.CreateTheme(driver)
	window := theme.CreateWindow(*widthFlag, *heightFlag, "Plot")
	window.SetScale(float32(*scaleFlag))
	img := theme.CreateImage()

	// texture := driver.CreateTexture(m, 1)
	// img.SetTexture(texture)

	rgba := image.NewRGBA(source.Bounds())
	draw.Draw(rgba, source.Bounds(), source, image.ZP, draw.Src)
	// draw.Draw(rgba, m.Bounds(), m, image.ZP, draw.Src)
	texture := driver.CreateTexture(rgba, 1)
	img.SetTexture(texture)

	window.AddChild(img)
	window.OnClose(driver.Terminate)
}

func init() {
	versionFlag := flag.Bool("v", false, "prints current version")
	quietFlag := flag.Bool("quiet", false, "do not print log messages (by default logger is noisy)")

	flag.Parse()

	au = aurora.NewAurora(*colorsFlag)

	log.SetOutput(os.Stdout)

	if *timestampFlag {
		log.SetFlags(log.Ldate | log.Ltime)
	}

	if *versionFlag {
		log.Println("SCplot version", au.Bold(SCPLOT_VERSION))
		os.Exit(0)
	}

	if *quietFlag {
		log.SetOutput(ioutil.Discard)
		log.SetFlags(0)
	}

	if *inputFlag == "" {
		// No input argument
		log.Println(au.Red("The '-input' argument is mandatory, use -h for help!"))
		os.Exit(0)
	} else if !fileExists(*inputFlag) {
		// Wrong input argument
		log.Println(au.Red("The file " + *inputFlag + " is not accessible, check if it exists and if it has correct permissions."))
		os.Exit(0)
	} else {
		// Correct input argument
		sourceWithoutExtension = fileWithoutExtension(*inputFlag)
		pngFilename = sourceWithoutExtension + ".png"
	}
}

func main() {
	gl.StartDriver(appMain)
}
