package dstream

// FilterFunc is a filtering function for use with Filter. The first
// argument holds a map from variable names to data slices. The second
// argument is a boolean slice that is initialized to all 'true'.  The
// FilterFunc should set elements of the boolean slice to false wherever
// the corresponding dstream record should be excluded.
type FilterFunc func(map[string]interface{}, []bool)

type filterCol struct {
	xform

	filter    FilterFunc
	keep      []bool
	keeppos   []int
	nobs      int
	nobsKnown bool
}

// Filter selects rows from the dstream.  A filtering function determines
// which rows are selected.
func Filter(data Dstream, f FilterFunc) Dstream {

	fc := &filterCol{
		xform: xform{
			source: data,
		},
		filter: f,
	}
	fc.init()

	return fc
}

func (fc *filterCol) init() {
	fc.names = fc.source.Names()

	// We need to know this early so call explicitly
	fc.setNamePos()

	fc.bdata = make([]interface{}, len(fc.names))
	fc.nobsKnown = false
	fc.nobs = 0
}

func (fc *filterCol) Reset() {
	fc.source.Reset()
	fc.nobs = 0
	fc.nobsKnown = false
	fc.init()
}

func (fc *filterCol) NumObs() int {
	if fc.nobsKnown {
		return fc.nobs
	}
	return -1
}
