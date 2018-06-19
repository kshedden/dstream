package dstream

import "fmt"

// GenerateFunc is a function that can be used to generate a new variable
// from existing variables.  The first argument, say m, is a map from
// variable names to data (whose concrete types are slices held as
// empty interfaces).  The second argument is a pre-allocated array (a
// slice provided as an interface{}) into which the new variable's
// values are to be written.  The destination array is not set to
// zeros before passing to the function.
type GenerateFunc func(map[string]interface{}, interface{})

type generate struct {
	xform

	gfunc      GenerateFunc
	newvarname string
	dtype      string
}

func (gen *generate) init() {

	for _, na := range gen.source.Names() {
		if gen.newvarname == na {
			msg := fmt.Sprintf("Generate: variable '%s' already exists.\n", gen.newvarname)
			panic(msg)
		}
	}

	gen.names = append(gen.source.Names(), gen.newvarname)
	gen.bdata = make([]interface{}, len(gen.names))

	// TODO make type generic
	switch gen.dtype {
	case "float64":
		gen.bdata[len(gen.bdata)-1] = make([]float64, 0)
	case "uint64":
		gen.bdata[len(gen.bdata)-1] = make([]float64, 0)
	case "string":
		gen.bdata[len(gen.bdata)-1] = make([]string, 0)
	default:
		panic("Generate: unknown dtype")
	}
}

// Generate appends a new variable to a Dstream, obtaining its values by
// applying the given function to the other variables in the Dstream.
// The new variable must not already exist in the Dstream.
func Generate(data Dstream, name string, fnc GenerateFunc, dtype string) Dstream {

	g := &generate{
		xform: xform{
			source: data,
		},
		newvarname: name,
		gfunc:      fnc,
		dtype:      dtype,
	}
	g.init()

	return g
}

func (gen *generate) Next() bool {

	if !gen.source.Next() {
		return false
	}

	// All but new variable
	for j := 0; j < gen.source.NumVar(); j++ {
		gen.bdata[j] = gen.source.GetPos(j)
	}

	n := ilen(gen.GetPos(0))
	q := len(gen.names) - 1

	// TODO make type generic
	switch x := gen.bdata[q].(type) {
	case []float64:
		gen.bdata[q] = resizefloat64(x, n)
	case []uint64:
		gen.bdata[q] = resizeuint64(x, n)
	case []string:
		gen.bdata[q] = resizestring(x, n)
	default:
		panic("unkown type")
	}

	mp := make(map[string]interface{})
	for k, na := range gen.names {
		if na != gen.newvarname {
			mp[na] = gen.bdata[k]
		}
	}

	gen.gfunc(mp, gen.bdata[q])
	return true
}
