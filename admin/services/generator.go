package services

import (
	// "errors"
	"bytes"
	"go/format"
	"io"
	"text/template"
    
	//"fmt"
)

type Format uint

const (
	JSON Format = iota
)

type Metadata struct {
	PackageName string
	TypeData    map[string]string
	// MarshalObject string
	TypeName string
}

type Generator struct {
	Format Format
}

func (g *Generator) Generate(writer io.Writer, metadata Metadata) error {
	tmpl, err := g.template()
	if err != nil {
		return err
	}
	//create in memory buffer to hold result of template execution
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, metadata)
	if err != nil {
		return err
	}
	//format contents of buffer before sending to writer
	formattedBuf, err2 := format.Source(buf.Bytes())
	if err2 != nil {
		return err2
	}
    //send along to file writer passed in as parameter
	_, err3 := bytes.NewBuffer(formattedBuf).WriteTo(writer)
	return err3
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
