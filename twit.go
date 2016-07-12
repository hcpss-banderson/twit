package main

import (
	"html/template"
	"io/ioutil"
	"os"

	"github.com/leekchan/gtf"
)

// Twit is the structural representation of the Twit application.
type Twit struct {
	TemplateParams TemplateParams
	Source         string
	Target         *os.File
	Filters        map[string]interface{}
}

// NewTwit create a new Twit instance.
func NewTwit(source, dest string, params TemplateParams) (*Twit, error) {
	t := &Twit{
		TemplateParams: params,
	}

	err := t.SetTargetFromPath(dest)
	if err != nil {
		return nil, err
	}

	err = t.SetSourceFromPath(source)
	if err != nil {
		return nil, err
	}

	return t, err
}

// SetSourceFromPath sets the source from a path to the file.
func (t *Twit) SetSourceFromPath(source string) error {
	templateBytes, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	t.Source = string(templateBytes)

	return nil
}

// SetTargetFromPath sets the target from a path to a file.
func (t *Twit) SetTargetFromPath(dest string) error {
	target, err := os.Create(dest)
	if err != nil {
		return err
	}

	t.Target = target

	return nil
}

// Render renders the template.
func (t *Twit) Render() {
	params := t.TemplateParams.ToMap()

	template.
		Must(gtf.New("text").Funcs(t.Filters).Parse(t.Source)).
		Execute(t.Target, params)
}
