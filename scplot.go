// +Build ignore
//go:generate go run -tags=dev assets_generate.go
package main

import (
	"io/ioutil"
	"log"
)

type ead_file struct {
	filename string
	metadata string
}

var ead_files []ead_file
var dictionary map[string]string

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

func main() {

}
