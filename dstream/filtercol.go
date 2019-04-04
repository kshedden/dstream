package dstream

// FilterFunc is a filtering function for use with Filter. The empty
// interface holds a slice of values of length n, the boolean array,
// denoted "Keep" below, also has length n.  The FilterFunc should set
// elements of Keep to false wherever the corresponding Dstream record
// is to be excluded.  Never set any element of Keep to true, as this
// may interfere with other FilterFuncs acting jointly with this
// one. The returned boolean indicates whether the FilterFunc entered
// any new false values into Keep.
type FilterFunc func(map[string]interface{}, []bool)

type filterCol struct {
	xform

	filter    FilterFunc
	keep      []bool
	keeppos   []int
	nobs      int
	nobsKnown bool
}

// Filter applies filtering functions to one or more data columns,
// and retains only the rows where all filtering functions are true.
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
