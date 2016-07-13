package main

import (
	"fmt"
	"log"
	"os"

	flag "github.com/ogier/pflag"
)

// VERSION of the application.
const VERSION = "v1.0.0"

var app *Twit

func init() {
	var version, noEscape bool
	var templateParams TemplateParams

	flag.VarP(&templateParams, "params", "p", "")
	flag.BoolVarP(&version, "version", "v", false, "")
	flag.BoolVarP(&noEscape, "no-escape", "n", false, "")

	flag.Usage = twitUsage

	flag.Parse()

	if version {
		fmt.Printf("twit %s\n", VERSION)
		os.Exit(0)
	}

	if flag.NArg() < 2 {
		log.Fatal("Not enough arguments.")
	}

	twit, err := NewTwit(flag.Arg(0), flag.Arg(1), templateParams, !noEscape)
	if err != nil {
		panic(err)
	}

	app = twit
}

func main() {
	app.Render()
}
