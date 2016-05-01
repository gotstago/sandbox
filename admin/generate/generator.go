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
    //r, w := io.Pipe()
    var buf bytes.Buffer
    //buf.WriteTo(w)
    //buf.ReadFrom(r)
    //format.Source
	//return tmpl.Execute(writer, metadata)
    tmpl.Execute(&buf, metadata)
    temp,_ := format.Source(buf.Bytes())
    //fmt.Println("buf is ",temp)
    //bytes.NewBuffer(temp)
    _, result := bytes.NewBuffer(temp).WriteTo(writer)
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
