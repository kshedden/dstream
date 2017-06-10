// +build ignore

package main

import (
	"bytes"
	"flag"
	"go/format"
	"html/template"
	"os"
	"strings"
)

type Dtype struct {
	Type string
}

var (
	NumTypes = []Dtype{
		Dtype{Type: "float64"},
		Dtype{Type: "float32"},
		Dtype{Type: "uint64"},
		Dtype{Type: "uint32"},
		Dtype{Type: "uint16"},
		Dtype{Type: "uint8"},
		Dtype{Type: "int64"},
		Dtype{Type: "int32"},
		Dtype{Type: "int16"},
		Dtype{Type: "int8"},
		Dtype{Type: "int"},
	}

	AllTypes = []Dtype{
		Dtype{Type: "string"},
	}
)

func main() {

	noformat := flag.Bool("noformat", false, "format code")
	numeric := flag.Bool("numeric", false, "only use numeric types")
	templatefile := flag.String("template", "", "template file")
	flag.Parse()

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
	if !*noformat {
		p, err = format.Source(buf.Bytes())
		if err != nil {
			panic(err)
		}
	} else {
		p = buf.Bytes()
	}

	outname := strings.Replace(*templatefile, ".template", "_gen.go", 1)
	out, err := os.Create(outname)
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
