package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// TemplateParams is the representation of template parameters.
type TemplateParams struct {
	flags  []string
	params map[string]interface{}
}

// String is the string representation of the TemplateParams.
func (tp *TemplateParams) String() string {
	return strings.Join(tp.flags, ", ")
}

// Set sets a new value from a flag.
func (tp *TemplateParams) Set(value string) error {
	// Let's keep track of the flags.
	tp.flags = append(tp.flags, value)

	err := tp.AddParamsFromFlag(value)

	return err
}

// AddParamsFromFlag adds params from a flag. Flag is either a path to a YAML
// encoded file or a JSON encoded string.
func (tp *TemplateParams) AddParamsFromFlag(flag string) error {
	params := make(map[string]interface{})

	jsonErr := json.Unmarshal([]byte(flag), &params)

	if jsonErr != nil {
		// We could not unmarshal as json so let's try as a YAML file.
		paramBytes, readErr := ioutil.ReadFile(flag)
		if readErr != nil {
			return readErr
		}

		yamlErr := yaml.Unmarshal(paramBytes, &params)
		if yamlErr != nil {
			return yamlErr
		}
	}

	if tp.params == nil {
		tp.params = make(map[string]interface{})
	}

	for k, v := range params {
		tp.params[k] = v
	}

	return nil
}

// ToMap returns the map of template parameters.
func (tp *TemplateParams) ToMap() map[string]interface{} {
	return tp.params
}
