package dstream

import "fmt"

// xform can be embedded into any transformer object that converts a
// Dstream into another Dstream.  It implements the basic behavior
// that is common to all transformers.  While xform implements the
// Dstream interface in a trivial way, Types that embed xform can
// provide their own implementation of any of these methods.
//
// In most cases, it will be sufficient to (i) set names properly, and
// (ii) implement Next to set bdata correctly (bdata holds the data
// slices by column).  In some cases it may be desirable to implement
// other Dstream methods.  Note that due to the way that Go embedding
// works, if a type implements its own, say, Names function, all the
// code below will continue to call the xform Names implementation,
// not the customized Names implementation.  For this reason, it is
// usually better to modify the bdata and names members rather than to
// provide new implementations of Get, etc.  If Get and other Dstream
// methods need to be implemented by a type, it may be necessary to
// implement all Dstream interface components.
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
	}
	return x.source.Names()
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
	return x.GetPos(pos)
}
