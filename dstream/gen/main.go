// +build ignore

package main

import (
	"bytes"
	"flag"
	"go/format"
	"os"
	"path"
	"strings"
	"text/template"
)

type Dtype struct {
	Type      string
	Utype     string
	ConvGroup string
}

var (
	NumTypes = []Dtype{
		Dtype{"uint8", "Uint8", "i"},
		Dtype{"uint16", "Uint16", "i"},
		Dtype{"uint32", "Uint32", "i"},
		Dtype{"uint64", "Uint64", "i"},
		Dtype{"int8", "Int8", "i"},
		Dtype{"int16", "Int16", "i"},
		Dtype{"int32", "Int32", "i"},
		Dtype{"int64", "Int64", "i"},
		Dtype{"float32", "Float32", "f"},
		Dtype{"float64", "Float64", "f"},
	}

	AllTypes = []Dtype{
		Dtype{"string", "String", "s"},
		Dtype{"time.Time", "Time", "t"},
	}
)

func main() {

	noformat := flag.Bool("noformat", false, "format code")
	numeric := flag.Bool("numeric", false, "only use numeric types")
	templatefile := flag.String("template", "", "template file")
	flag.Parse()

	if *templatefile == "" {
		panic("'template' is a required argument")
	}

	AllTypes = append(AllTypes, NumTypes...)

	tmpl, err := template.ParseFiles(*templatefile)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if *numeric {
		err = tmpl.Execute(&buf, NumTypes)
	} else {
		err = tmpl.Execute(&buf, AllTypes)
	}
	if err != nil {
		panic(err)
	}

	var p []byte
	if *noformat {
		p = buf.Bytes()
	} else {
		p, err = format.Source(buf.Bytes())
		if err != nil {
			panic(err)
		}
	}

	outname := strings.Replace(*templatefile, ".tmpl", "_gen.go", 1)
	out, err := os.Create(path.Join("..", outname))
	if err != nil {
		panic(err)
	}
	out.WriteString("// GENERATED CODE, DO NOT EDIT\n\n")
	_, err = out.Write(p)
	if err != nil {
		panic(err)
	}
	out.Close()
}
