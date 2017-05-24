// +build ignore

package main

import (
	"bytes"
	"go/format"
	"os"
	"strings"
	"text/template"
)

const (
	doformat = true
)

type Dtype struct {
	Type string
}

var (
	Dtypes = []Dtype{
		Dtype{Type: "float64"},
		Dtype{Type: "float32"},
		Dtype{Type: "uint64"},
		Dtype{Type: "uint32"},
		Dtype{Type: "uint16"},
		Dtype{Type: "uint8"},
		Dtype{Type: "string"},
	}
)

func main() {

	if len(os.Args) != 2 {
		panic("wrong number of arguments")
	}

	tmpl, err := template.ParseFiles(os.Args[1])
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, Dtypes)
	if err != nil {
		panic(err)
	}

	var p []byte
	if doformat {
		p, err = format.Source(buf.Bytes())
		if err != nil {
			panic(err)
		}
	} else {
		p = buf.Bytes()
	}

	outname := strings.Replace(os.Args[1], ".template", "_gen.go", 1)
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
