package dstream

// ApplyFunc is a function that can be used to generate a new variable
// from existing variables.  The first argument, say m, is a map from
// variable names to data (whose concrete types are slices held as
// empty interfaces).  The second argument is a pre-allocated array
// (also a slice held as an empty interface) into which the new
// variable's values are to be written.
type ApplyFunc func(map[string]interface{}, interface{})

type apply struct {
	xform

	afunc      ApplyFunc
	newvarname string
	dtype      string
}

func (a *apply) init() {

	a.names = append(a.source.Names(), a.newvarname)
	a.bdata = make([]interface{}, len(a.names))

	// TODO make type generic
	switch a.dtype {
	case "float64":
		a.bdata[len(a.bdata)-1] = make([]float64, 0)
	case "string":
		a.bdata[len(a.bdata)-1] = make([]string, 0)
	default:
		panic("Apply: unkown dtype")
	}
}

// Apply appends a new variable to a Dstream, obtaining its values by
// applying the given function to the other variables in the Dstream.
func Apply(data Dstream, name string, fnc ApplyFunc, dtype string) Dstream {

	a := &apply{
		xform: xform{
			source: data,
		},
		newvarname: name,
		afunc:      fnc,
		dtype:      dtype,
	}
	a.init()

	return a
}

func (a *apply) Next() bool {

	if !a.source.Next() {
		return false
	}

	// All but new variable
	for j := 0; j < a.source.NumVar(); j++ {
		a.bdata[j] = a.source.GetPos(j)
	}

	n := ilen(a.GetPos(0))
	q := len(a.names) - 1

	// TODO make type generic
	switch x := a.bdata[q].(type) {
	case []float64:
		a.bdata[q] = resizefloat64(x, n)
	case []string:
		a.bdata[q] = resizestring(x, n)
	default:
		panic("unkown type")
	}

	mp := make(map[string]interface{})
	for k, na := range a.names {
		if na != a.newvarname {
			mp[na] = a.bdata[k]
		}
	}

	a.afunc(mp, a.bdata[q])
	return true
}
