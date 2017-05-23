package dstream

import "math"

func DropNA(data Dstream) Dstream {
	dna := &dropNA{
		xform: xform{
			source: data,
		},
	}

	return dna
}

type dropNA struct {
	xform

	mask      []bool
	pos       []int
	nobs      int
	nobsKnown bool
}

func (dna *dropNA) GetPos(j int) interface{} {
	return dna.bdata[j]
}

func (dna *dropNA) Names() []string {
	return dna.source.Names()
}

func (dna *dropNA) NumVar() int {
	return dna.source.NumVar()
}

func (dna *dropNA) NumObs() int {
	if dna.nobsKnown {
		return dna.nobs
	} else {
		return -1
	}
}

func (dna *dropNA) Reset() {
	dna.source.Reset()
	dna.nobsKnown = false

}

func (dna *dropNA) Next() bool {

	if !dna.source.Next() {
		dna.nobsKnown = true
		return false
	}

	// The size of the source data chunk
	n := ilen(dna.source.GetPos(0))

	dna.mask = resizeBool(dna.mask, n)
	for j, _ := range dna.mask {
		dna.mask[j] = false
	}

	// Get the missing mask (true = missing)
	nvar := dna.source.NumVar()
	for j := 0; j < nvar; j++ {
		v := dna.source.GetPos(j)
		switch v := v.(type) {
		case []float64:
			for i, y := range v {
				if math.IsNaN(y) {
					dna.mask[i] = true
				}
			}
		case []string:
			// Currently strings cannot be missing
			// TODO: allow specified labels to denote missing
		}
	}

	dna.pos = wherefalse(dna.mask, dna.pos)
	m := len(dna.pos)
	dna.nobs += m

	data := dna.bdata
	if data == nil {
		data = make([]interface{}, nvar)
	}
	for j := 0; j < nvar; j++ {
		v := dna.source.GetPos(j)
		switch v := v.(type) {
		case []float64:
			var x []float64
			if data[j] != nil {
				x = data[j].([]float64)
			}
			x = resizefloat64(x, m)
			for i, p := range dna.pos {
				x[i] = v[p]
			}
			data[j] = x
		case []string:
			var x []string
			if data[j] != nil {
				x = data[j].([]string)
			}
			x = resizestring(x, m)
			for i, p := range dna.pos {
				x[i] = v[p]
			}
			data[j] = x
		default:
			panic("unknown data type")
		}
	}

	dna.bdata = data
	return true
}
