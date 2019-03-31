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
	Type  string
	Utype string
}

var (
	NumTypes = []Dtype{
		Dtype{"uint8", "Uint8"},
		Dtype{"uint16", "Uint16"},
		Dtype{"uint32", "Uint32"},
		Dtype{"uint64", "Uint64"},
		Dtype{"int8", "Int8"},
		Dtype{"int16", "Int16"},
		Dtype{"int32", "Int32"},
		Dtype{"int64", "Int64"},
		Dtype{"float32", "Float32"},
		Dtype{"float64", "Float64"},
	}

	AllTypes = []Dtype{
		Dtype{"string", "String"},
		Dtype{"time.Time", "Time"},
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
