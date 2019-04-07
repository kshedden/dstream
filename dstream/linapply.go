package dstream

import "fmt"

type linapply struct {
	xform

	coeffs   [][]float64
	basename string
}

func (a *linapply) init() {

	snames := a.source.Names()
	names := make([]string, len(snames), len(snames)+len(a.coeffs))
	copy(names, snames)
	for j := 0; j < len(a.coeffs); j++ {
		names = append(names, fmt.Sprintf("%s%d", a.basename, j))
	}
	a.names = names

	a.bdata = make([]interface{}, len(a.names))
}

// Linapply adds new variables to Dstream by taking linear
// combinations of the other variables in the Dstream.
func Linapply(data Dstream, coeffs [][]float64, basename string) Dstream {

	a := &linapply{
		xform: xform{
			source: data,
		},
		coeffs:   coeffs,
		basename: basename,
	}
	a.init()

	return a
}

func (a *linapply) Next() bool {

	f := a.source.Next()
	if !f {
		return false
	}

	// Copy the existing variables directly
	for j := 0; j < a.source.NumVar(); j++ {
		a.bdata[j] = a.source.GetPos(j)
	}

	n := ilen(a.GetPos(0))
	q := a.source.NumVar()

	// Loop over new variables to create
	for j, v := range a.coeffs {

		if len(v) != q {
			msg := fmt.Sprintf("coefs (d=%d) does not conform with variables (d=%d)", len(v), q)
			panic(msg)
		}

		// Create space for new variable
		var z []float64
		if a.bdata[q+j] != nil {
			z = a.bdata[q+j].([]float64)
		}
		z = resizeFloat64(z, n)
		a.bdata[q+j] = z
		zeroFloat(z)

		// Loop over summands
		for k := 0; k < q; k++ {
			switch x := a.bdata[k].(type) {
			case []float64:
				for i := range z {
					z[i] += v[k] * x[i]
				}
			case []string:
				// do nothing
			default:
				panic("unknown type")
			}
		}
	}

	return true
}
