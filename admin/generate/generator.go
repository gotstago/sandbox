package main

import (
	// "errors"
	"io"
	"text/template"
    "go/format"
    "bytes"
    //"fmt"
)

type Format uint

const (
	JSON Format = iota
)

type Metadata struct {
	PackageName   string
	TypeData        map[string]string
	// MarshalObject string
	TypeName          string
}

type Generator struct {
	Format Format
}

func (g *Generator) Generate(writer io.Writer, metadata Metadata) error {
	tmpl, err := g.template()
	if err != nil {
		return nil
	}
    //create in memory buffer to hold result of template execution
    var buf bytes.Buffer
    tmpl.Execute(&buf, metadata)
    //format contents of buffer before sending to writer
    formattedBuf,_ := format.Source(buf.Bytes())
    //send along to file writer passed in as parameter
    _, result := bytes.NewBuffer(formattedBuf).WriteTo(writer)
    return result
}

func (g *Generator) template() (*template.Template, error) {
	// if g.Format != JSON {
	// 	return nil, errors.New("Unsupported format")
	// }

	resource, err := Asset("templates/model.tmpl")
	if err != nil {
		return nil, err
	}

	tmpl := template.New("template")
	return tmpl.Parse(string(resource))
}
