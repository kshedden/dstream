package dstream

//go:generate go run gen.go bcols.template

import (
	"fmt"
	"io"
	"sort"
)

type ByteReaderReader interface {
	io.Reader
	io.ByteReader
}

type bcols struct {
	bpath string

	dtypes map[string]string

	names []string

	bdata []interface{}

	rdrs []ByteReaderReader

	toclose []io.Closer

	chunksize int

	namepos map[string]int

	done bool

	nobsKnown bool
	nobs      int

	include []string
	exclude []string
}

// BCols takes data stored in a column-wise compressed format under
// the given directory path, and returns it via a Dstream.
func NewBCols(dpath string, chunksize int, include, exclude []string) Dstream {

	b := &bcols{
		bpath:     dpath,
		chunksize: chunksize,
		include:   include,
		exclude:   exclude,
	}
	b.init()

	return b
}

// usenames returns variable names to use.
func (b *bcols) usenames() []string {

	var use []string
	for v, _ := range b.dtypes {

		// If a name is in the Exclude list, exclude it
		q := false
		for _, x := range b.exclude {
			if x == v {
				q = true
				break
			}
		}
		if q {
			continue
		}

		// If there is an Include list, and a name is not in
		// it, exclude it.
		if len(b.include) > 0 {
			q := true
			for _, x := range b.include {
				if x == v {
					q = false
					break
				}
			}
			if q {
				continue
			}
		}

		use = append(use, v)
	}

	sort.StringSlice(use).Sort()

	return use
}

func (b *bcols) Names() []string {
	return b.names
}

func (b *bcols) Close() {
	for _, x := range b.toclose {
		x.Close()
	}
}

func (b *bcols) Reset() {
	b.Close()
	b.toclose = b.toclose[0:0]
	b.rdrs = b.rdrs[0:0]
	b.init()
	b.nobsKnown = false
	b.nobs = 0
	b.done = false
}

func (b *bcols) NumVar() int {
	return len(b.names)
}

func (b *bcols) NumObs() int {
	if !b.nobsKnown {
		return -1
	}
	return b.nobs
}

func (b *bcols) GetPos(j int) interface{} {
	return b.bdata[j]
}

func (b *bcols) Get(na string) interface{} {
	pos, ok := b.namepos[na]
	if !ok {
		msg := fmt.Sprintf("Variable '%s' not found", na)
		panic(msg)
	}
	return b.bdata[pos]
}
