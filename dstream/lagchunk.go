package dstream

import (
	"fmt"
)

// TODO lagChunk should use xform
type lagChunk struct {
	SourceData Dstream

	// Lags maps variable names to the number of lags to include
	// for the variable.  Variables not included in the map are
	// retained with no lags.
	Lags map[string]int

	// Backing data for slices returned by Get()
	ldata []interface{}

	// Names of the variables
	names []string

	namespos map[string]int

	nobs      int  // total sample size
	nobsKnown bool // indicates whether the sample size is available
	maxlag    int  // maximum of all requested lags
	doneInit  bool // init has run
}

// LagChunk returns a new Dstream in which specified variables are
// included with lagged values.  Lagged values are only computed
// within a chunk, not across chunk boundaries, and the first m values
// of each chunk are omitted, where m is the maximum lag value.
func LagChunk(data Dstream, lags map[string]int) Dstream {
	return &lagChunk{
		SourceData: data,
		Lags:       lags,
	}
}

func (lc *lagChunk) Names() []string {
	if !lc.doneInit {
		lc.init()
	}
	return lc.names
}

func (lc *lagChunk) init() {
	maxlag := 0
	for _, v := range lc.Lags {
		if v > maxlag {
			maxlag = v
		}
	}
	lc.maxlag = maxlag

	// Create the names of the new variables
	var names []string
	for _, a := range lc.SourceData.Names() {
		if lc.Lags[a] == 0 {
			names = append(names, a)
		} else {
			for j := 0; j <= lc.Lags[a]; j++ {
				b := fmt.Sprintf("%s[%d]", a, -j)
				names = append(names, b)
			}
		}
	}

	lc.names = names

	lc.namespos = make(map[string]int)
	for pos, na := range lc.names {
		lc.namespos[na] = pos
	}

	lc.doneInit = true
}

func (lc *lagChunk) GetPos(j int) interface{} {
	return lc.ldata[j]
}

func (lc *lagChunk) Get(na string) interface{} {

	pos, ok := lc.namespos[na]
	if !ok {
		msg := fmt.Sprintf("Variable '%s' not found", na)
		panic(msg)
	}

	return lc.ldata[pos]
}

func (lc *lagChunk) NumObs() int {
	if lc.nobsKnown {
		return lc.nobs
	} else {
		return -1
	}
}

func (lc *lagChunk) NumVar() int {
	if !lc.doneInit {
		lc.init()
	}
	return len(lc.names)
}

func (lc *lagChunk) Reset() {
	lc.SourceData.Reset()
	lc.doneInit = false
}

func (lc *lagChunk) Next() bool {

	if !lc.doneInit {
		lc.init()
	}

	if !lc.SourceData.Next() {
		lc.nobsKnown = true
		return false
	}

	if lc.ldata == nil {
		lc.ldata = make([]interface{}, len(lc.names))
	}

	// Loop over the original data columns
	jj := 0
	maxlag := lc.maxlag
	for j, oname := range lc.SourceData.Names() {

		v := lc.SourceData.GetPos(j)
		if ilen(v) <= maxlag {
			// Segment is too short to use
			continue
		}

		q := lc.Lags[oname]
		switch v := v.(type) {
		case []float64:
			n := len(v)
			lc.nobs += n - maxlag
			for k := 0; k <= q; k++ {
				lc.ldata[jj] = v[(maxlag - k):(n - k)]
				jj++
			}
		case []string:
			n := len(v)
			lc.nobs += n - maxlag
			for k := 0; k <= q; k++ {
				lc.ldata[jj] = v[(maxlag - k):(n - k)]
				jj++
			}
		default:
			panic("unkown data type")

		}
	}

	return true
}
