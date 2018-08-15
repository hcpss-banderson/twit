package main

import (
	"fmt"
	"log"
	"os"

	flag "github.com/ogier/pflag"
	"github.com/fsnotify/fsnotify"
)

// VERSION of the application.
const VERSION = "v1.1.0"

var version, noEscape, watch bool
var source, destination, params string

func init() {
	flag.StringVarP(&params, "params", "p", "", "")
	flag.BoolVarP(&version, "version", "v", false, "")
	flag.BoolVarP(&noEscape, "no-escape", "n", false, "")
	flag.BoolVarP(&watch, "watch", "w", false, "")

	flag.Usage = twitUsage

	flag.Parse()
	
	source = flag.Arg(0)
	destination = flag.Arg(1)

	if version {
		fmt.Printf("twit %s\n", VERSION)
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		log.Fatal("Not enough arguments.")
	}
	
	if watch && flag.NArg() < 2 {
		log.Fatal("To use watch, you have to specify a destination.")
	}
}

func rerender(twit *Twit, name string) {
	templateParams := TemplateParams{}
	templateParams.Set(params)
	fmt.Printf("Changes detected in %#v\n", name)
	fmt.Println("Rewriting template")
	templateParams.Set(params)
	twit.TemplateParams = templateParams
	twit.SetSourceFromPath(source)
	twit.Render()
}

func main() {
	templateParams := TemplateParams{}
	templateParams.Set(params)
	
	output := os.Stdout
	if destination == "" {
		output = os.Stdout
	} else {
		output, _ = os.Create(destination)
	}
	
	twit, err := NewTwit(source, output, templateParams, !noEscape)	
	if err != nil {
		panic(err)
	}
	
	twit.Render()
	
	if watch {
		templateWatcher, err := fsnotify.NewWatcher()
		if err != nil {
			panic(err)
		}
		defer templateWatcher.Close()
		
		paramsWatcher, err := fsnotify.NewWatcher()
		if err != nil {
			panic(err)
		}
		defer paramsWatcher.Close()
		
		done := make(chan bool)
		
		go func() {
			for {
				select {
				case event := <-templateWatcher.Events:
					rerender(twit, event.Name)
				case err := <-templateWatcher.Errors:
					panic(err)
				case event := <-paramsWatcher.Events:
					rerender(twit, event.Name)
				case err := <-paramsWatcher.Errors:
					panic(err)
				}
			}
		}()
		
		if err := templateWatcher.Add(source); err != nil {
			panic(err)
		} else {
			fmt.Println("Watching" + source + " ...")
		}
		
		if err := paramsWatcher.Add(params); err != nil {
			// This probably means the params were passed in as json.
		} else {
			fmt.Println("Watching" + params + " ...")
		}
		
		<-done
	}
}
