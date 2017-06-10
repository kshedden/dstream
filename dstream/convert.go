package dstream

import "fmt"

//go:generate go run gen.go -template=convert.template -numeric

type convertType struct {
	xform

	// The name of the variable to be converted
	vname string

	// The position of the variable to be converted
	vpos int

	// We convert lazily, this indicates whether a column in the
	// current chunk has already been converted.
	conv []bool

	// The new type of the variable
	dtype string
}

// Convert returns a Dstream in which the named variable is converted
// to the given type.
func Convert(da Dstream, vname string, dtype string) Dstream {

	vpos := -1
	for k, na := range da.Names() {
		if na == vname {
			vpos = k
			break
		}
	}
	if vpos == -1 {
		msg := fmt.Sprintf("Convert: no variable '%s'\n", vname)
		panic(msg)
	}

	c := &convertType{
		xform: xform{
			source: da,
		},
		vname: vname,
		dtype: dtype,
		vpos:  vpos,
	}

	return c
}

func (c *convertType) Next() bool {

	if !c.source.Next() {
		return false
	}

	if len(c.bdata) == 0 {
		c.bdata = make([]interface{}, c.source.NumVar())
	}
	truncate(c.bdata)

	if len(c.conv) == 0 {
		c.conv = make([]bool, c.source.NumVar())
	} else {
		for j := range c.conv {
			c.conv[j] = false
		}
	}

	return true
}
