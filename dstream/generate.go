package dstream

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
