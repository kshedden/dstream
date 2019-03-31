package dstream

import "fmt"

// FilterFunc is a filtering function for use with Filter. The empty
// interface holds a slice of values of length n, the boolean array,
// denoted "Keep" below, also has length n.  The FilterFunc should set
// elements of Keep to false wherever the corresponding Dstream record
// is to be excluded.  Never set any element of Keep to true, as this
// may interfere with other FilterFuncs acting jointly with this
// one. The returned boolean indicates whether the FilterFunc entered
// any new false values into Keep.
type FilterFunc func(interface{}, []bool) bool

type filterCol struct {
	xform

	filters   map[string]FilterFunc
	keep      []bool
	keeppos   []int
	nobs      int
	nobsKnown bool
}

// Filter applies filtering functions to one or more data columns,
// and retains only the rows where all filtering functions are true.
func Filter(data Dstream, funcs map[string]FilterFunc) Dstream {
	fc := &filterCol{
		xform: xform{
			source: data,
		},
		filters: funcs,
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

func (fc *filterCol) Next() bool {

	if !fc.source.Next() {
		fc.nobsKnown = true
		return false
	}

	nvar := fc.source.NumVar()
	n := ilen(fc.source.GetPos(0))
	if n == 0 {
		return true
	}

	fc.keep = resizeBool(fc.keep, n)
	for i := 0; i < n; i++ {
		fc.keep[i] = true
	}

	anydrop := false
	for na, fu := range fc.filters {
		j, ok := fc.namepos[na]
		if !ok {
			msg := fmt.Sprintf("Variable '%s' not found.", na)
			panic(msg)
		}
		x := fc.source.GetPos(j)
		ad := fu(x, fc.keep)
		anydrop = anydrop || ad
	}

	nkp := n
	if anydrop {
		fc.keeppos = fc.keeppos[0:0]
		for i, b := range fc.keep {
			if b {
				fc.keeppos = append(fc.keeppos, i)
			}
		}
		nkp = len(fc.keeppos)
		fc.nobs += nkp
	} else {
		fc.nobs += n
	}

	for j := 0; j < nvar; j++ {
		v := fc.source.GetPos(j)

		//TODO needs all types
		switch v := v.(type) {
		case []float64:
			if anydrop {
				var u []float64
				if fc.bdata[j] != nil {
					u = fc.bdata[j].([]float64)
				}
				u = resizeFloat64(u, nkp)
				for i, j := range fc.keeppos {
					u[i] = v[j]
				}
				fc.bdata[j] = u
			} else {
				fc.bdata[j] = v
			}
		case []string:
			if anydrop {
				var u []string
				if fc.bdata[j] != nil {
					u = fc.bdata[j].([]string)
				}
				u = resizeString(u, nkp)
				for i, j := range fc.keeppos {
					u[i] = v[j]
				}
				fc.bdata[j] = u
			} else {
				fc.bdata[j] = v
			}
		default:
			msg := fmt.Sprintf("unkown type %T", fc.bdata[j])
			panic(msg)
		}
	}

	return true
}
