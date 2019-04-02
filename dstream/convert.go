package dstream

import "fmt"

type convert struct {
	xform

	// The name of the variable to be converted
	vname string

	// The position of the variable to be converted
	vpos int

	// The new type of the variable
	dtype Dtype
}

// Convert returns a Dstream in which the named variable is converted
// to the given type.
func Convert(da Dstream, vname string, dtype Dtype) Dstream {

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

	c := &convert{
		xform: xform{
			source: da,
		},
		vname: vname,
		dtype: dtype,
		vpos:  vpos,
	}

	c.bdata = make([]interface{}, c.source.NumVar())

	return c
}
