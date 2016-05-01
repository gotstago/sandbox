package main

import (
	"errors"
	"io"
	"text/template"
    "go/format"
    "bytes"
)

type Format uint

const (
	JSON Format = iota
)

type Metadata struct {
	PackageName   string
	Object        string
	MarshalObject string
	Type          string
}

type Generator struct {
	Format Format
}

func (g *Generator) Generate(writer io.Writer, metadata Metadata) error {
	tmpl, err := g.template()
	if err != nil {
		return nil
	}
    //r, w := io.Pipe()
    buf := bytes.Buffer
    //buf.WriteTo(w)
    //buf.ReadFrom(r)
    //format.Source
	//return tmpl.Execute(writer, metadata)
    tmpl.Execute(buf, metadata)
    buf.WriteTo(writer)
}

func (g *Generator) template() (*template.Template, error) {
	if g.Format != JSON {
		return nil, errors.New("Unsupported format")
	}

	resource, err := Asset("templates/writeto_json.tmpl")
	if err != nil {
		return nil, err
	}

	tmpl := template.New("template")
	return tmpl.Parse(string(resource))
}
