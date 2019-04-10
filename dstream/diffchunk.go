package dstream

import (
	"fmt"
)

type diffChunk struct {
	xform

	order     map[string]int
	maxorder  int
	nobs      int  // total sample size
	nobsKnown bool // indicates whether the sample size is available
	doneInit  bool // init has run
}

// DiffChunk returns a new Dstream in which specified variables are
// differenced.  The differenced values are only computed within a
// chunk, not across chunk boundaries, and the first value of each
// chunk is omitted.
func DiffChunk(data Dstream, order map[string]int) Dstream {
	d := &diffChunk{
		xform: xform{
			source: data,
		},
		order: order,
	}
	d.init()
	return d
}

func (df *diffChunk) init() {
	maxorder := 0
	for _, v := range df.order {
		if v > maxorder {
			maxorder = v
		}
	}
	df.maxorder = maxorder

	// Create the names of the new variables
	var names []string
	for _, a := range df.source.Names() {
		names = append(names, a)
		q := df.order[a]
		if q > 0 {
			b := fmt.Sprintf("%s$d%d", a, q)
			names = append(names, b)
		}
	}

	df.bdata = make([]interface{}, len(names))
	df.names = names
	df.doneInit = true
}

func (df *diffChunk) Nobs() int {
	if df.nobsKnown {
		return df.nobs
	}
	return -1
}

func (df *diffChunk) Next() bool {

	if !df.doneInit {
		df.init()
	}

	if !df.source.Next() {
		df.nobsKnown = true
		return false
	}

	if df.bdata == nil {
		df.bdata = make([]interface{}, len(df.names))
	}

	// Loop over the original data columns
	jj := 0
	maxorder := df.maxorder
	for j, oname := range df.source.Names() {

		v := df.source.GetPos(j)
		if ilen(v) <= maxorder {
			// Segment is too short to use
			continue
		}

		q := df.order[oname]
		switch v := v.(type) {
		case []float64:
			n := len(v)
			df.nobs += n - maxorder
			df.bdata[jj] = v[maxorder:]
			jj++
			if q > 0 {
				var y []float64
				if df.bdata[jj] != nil {
					y = df.bdata[jj].([]float64)
				}
				y = resizeFloat64(y, n)
				copy(y, v)
				y = diff(y, q)
				if q < maxorder {
					y = y[maxorder-q:]
				}
				df.bdata[jj] = y
				jj++
			}
		case []string:
			n := len(v)
			df.nobs += n - maxorder
			df.bdata[jj] = v[maxorder:]
			jj++
		default:
			msg := fmt.Sprintf("unknown data type: %T", v)
			panic(msg)
		}
	}

	return true
}

func diff1(x []float64) []float64 {
	for i := len(x) - 1; i > 0; i-- {
		x[i] -= x[i-1]
	}
	return x[1:]
}

func diff(x []float64, ord int) []float64 {
	for j := 0; j < ord; j++ {
		x = diff1(x)
	}
	return x
}
