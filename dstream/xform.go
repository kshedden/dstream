package dstream

import "fmt"

// xform is embedded into any transformer object that converts a
// Dstream into another Dstream.  It implements the basic
// functionality that is common to most such transformers.
type xform struct {

	// The data to be transformed.
	source Dstream

	// The transformed data, held as references or copies (if
	// needed) of the source data.
	bdata []interface{}

	// List of variable names, used if names for the xform are
	// different from the names of its source, otherwise use
	// source.Names().
	names []string

	// Map from variable name to column number
	namepos map[string]int
}

func (x *xform) Close() {
}

// setNamePos constructs a map from variable names to their column
// positions.
func (x *xform) setNamePos() {
	x.namepos = make(map[string]int)
	for k, n := range x.Names() {
		x.namepos[n] = k
	}
}

func (x *xform) Reset() {
	x.source.Reset()
	if x.bdata != nil {
		truncate(x.bdata)
	}
}

func (x *xform) Next() bool {
	return x.source.Next()
}

func (x *xform) NumVar() int {
	return len(x.Names())
}

func (x *xform) NumObs() int {
	return x.source.NumObs()
}

func (x *xform) Names() []string {
	// If there is a names, then the xform has modified the names
	// so we need to return the modified names, otherwise pass
	// through to the source.
	if x.names != nil {
		return x.names
	} else {
		return x.source.Names()
	}
}

func (x *xform) GetPos(j int) interface{} {
	// If there is a bdata, then the xform has modified the source
	// data so we need to return the modified data, otherwise pass
	// through to the source.
	if x.bdata != nil {
		return x.bdata[j]
	}
	return x.source.GetPos(j)
}

func (x *xform) Get(na string) interface{} {

	if x.namepos == nil {
		x.setNamePos()
	}

	pos, ok := x.namepos[na]
	if !ok {
		msg := fmt.Sprintf("Variable '%s' not found", na)
		panic(msg)
	}
	return x.bdata[pos]
}
