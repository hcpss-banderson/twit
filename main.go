package main

import (
	"fmt"
	htmlTemplate "html/template"
	"io/ioutil"
	"log"
	"os"
	textTemplate "text/template"

	flag "github.com/ogier/pflag"
)

// VERSION of the application.
const VERSION = "v0.1.0"

var (
	source, destination string
	params              map[string]interface{}
	html                bool
)

func init() {
	var paramsPath, paramsJSON string
	var version bool

	flag.StringVarP(&paramsPath, "params-file", "f", "", "")
	flag.StringVarP(&paramsJSON, "params", "p", "", "")
	flag.BoolVarP(&html, "html", "m", false, "")
	flag.BoolVarP(&version, "version", "v", false, "")

	flag.Usage = twitUsage

	flag.Parse()

	if version {
		fmt.Printf("twit %s\n", VERSION)
		os.Exit(0)
	}

	if flag.NArg() < 2 {
		log.Fatal("Not enough arguments.")
	}

	source = flag.Arg(0)
	destination = flag.Arg(1)
	setJSONParams(paramsJSON)
	setFileParams(paramsPath)
}

func main() {
	destinationFile, err := os.Create(destination)
	if err != nil {
		log.Fatalf(
			"Could not create destination file %s. The error was %s",
			destination,
			err.Error(),
		)
	}
	defer destinationFile.Close()

	templateBytes, err := ioutil.ReadFile(source)
	if err != nil {
		log.Fatalf(
			"Could not read the template %s. The error was %s",
			source,
			err.Error(),
		)
	}

	render(string(templateBytes), destinationFile, params)
}

func render(template string, destFile *os.File, params map[string]interface{}) {
	if html {
		htmlTemplate.
			Must(htmlTemplate.New("html").Parse(template)).
			Execute(destFile, params)
	} else {
		textTemplate.
			Must(textTemplate.New("text").Parse(template)).
			Execute(destFile, params)
	}
}
