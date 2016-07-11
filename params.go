package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func setJSONParams(jsonString string) {
	if jsonString == "" {
		return
	}

	err := resolveParams(jsonString, json.Unmarshal)
	if err != nil {
		log.Fatalf(
			"Could not read parameters. The error was: %s",
			err.Error(),
		)
	}
}

func setFileParams(paramsPath string) {
	if paramsPath == "" {
		return
	}
	paramsBytes, err := ioutil.ReadFile(paramsPath)
	if err != nil {
		log.Fatalf(
			"Could not read file %s. The error was: %s",
			paramsPath,
			err.Error(),
		)
	}

	err = resolveParams(string(paramsBytes), yaml.Unmarshal)
	if err != nil {
		log.Fatalf(
			"Could not read parameters from %s. The error was: %s",
			paramsPath,
			err.Error(),
		)
	}
}

type unmarshaller func([]byte, interface{}) error

func resolveParams(
	paramsString string,
	callback func([]byte, interface{}) error,
) (err error) {
	var localParams map[string]interface{}

	if params == nil {
		params = make(map[string]interface{})
	}

	err = callback([]byte(paramsString), &localParams)
	if err != nil {
		return
	}

	for k, v := range localParams {
		params[k] = v
	}

	return
}
