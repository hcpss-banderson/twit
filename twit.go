package main

import (
	hTemplate "html/template"
	"io/ioutil"
	"io"
	"os"
	tTemplate "text/template"

	// "github.com/leekchan/gtf"
	"gtf"
)

// Twit is the structural representation of the Twit application.
type Twit struct {
	TemplateParams TemplateParams
	Source         string
	Target         io.Writer
	HTML           bool
}

// NewTwit create a new Twit instance.
func NewTwit(source string, dest io.Writer, params TemplateParams, html bool) (*Twit, error) {
	t := &Twit{
		TemplateParams: params,
		HTML:           html,
	}

	t.Target = dest

	err := t.SetSourceFromPath(source)
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

// Render renders the template.
func (t *Twit) Render() {
	params := t.TemplateParams.ToMap()
	
	ta, ok := t.Target.(*os.File)
	if ok {
		ta.Truncate(0)
		ta.Seek(0,0)
	}

	if t.HTML {
		hTemplate.
			Must(gtf.New("html").Parse(t.Source)).
			Execute(t.Target, params)
	} else {
		// The Gtf package provides a nice set of filters, but as a
		// html/template.FuncMap and we need a text/template.FuncMap.
		funcMap := make(tTemplate.FuncMap)
		for name, function := range gtf.GtfFuncMap {
			funcMap[name] = function
		}

		tTemplate.
			Must(tTemplate.New("text").Funcs(funcMap).Parse(t.Source)).
			Execute(t.Target, params)
	}
}
