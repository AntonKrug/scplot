package main

import (
	"flag"
	"github.com/logrusorgru/aurora"
	"io/ioutil"
	"log"
	"os"
)

var timestampFlag = flag.Bool("timestamped_log", true, "enable logger's timestamps")
var colorsFlag = flag.Bool("colors", false, "enable logger's colors (disabled by default)")
var inputFlag = flag.String("input", "", "raw dump file of a variable to be read as input")

var sourceDir string
var outputDir string

var au aurora.Aurora

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
}
