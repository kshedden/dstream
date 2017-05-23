package dstream

type sizechunk struct {
	xform

	size int

	stash []interface{}
}

func SizeChunk(data Dstream, size int) Dstream {

	sc := &sizechunk{
		xform: xform{
			source: data,
		},
		size: size,
	}

	nvar := data.NumVar()
	sc.names = sc.source.Names()
	sc.bdata = make([]interface{}, nvar)
	sc.stash = make([]interface{}, nvar)

	return sc
}

// setFromStash sets bdata using the stash.  Only can be called if the
// stash arrays are big enough to hold a chunk.
func (sc *sizechunk) setFromStash() {

	nvar := sc.NumVar()
	n := sc.size
	for j := 0; j < nvar; j++ {
		switch x := sc.stash[j].(type) {
		case []float64:
			sc.bdata[j] = x[0:n]
			sc.stash[j] = x[n:len(x)]
		case []string:
			sc.bdata[j] = x[0:n]
			sc.stash[j] = x[n:len(x)]
		default:
			panic("unkown type")
		}
	}

}

func (sc *sizechunk) Next() bool {

	for {
		if ilen(sc.stash[0]) >= sc.size {
			sc.setFromStash()
			return true
		}

		f := sc.source.Next()
		if !f {
			// What remains is final chunk
			copy(sc.bdata, sc.stash)
			for j, _ := range sc.stash {
				sc.stash[j] = nil
			}
			return false
		}

		for j := 0; j < sc.NumVar(); j++ {
			x := sc.source.GetPos(j)
			switch x := x.(type) {
			case []float64:
				var y []float64
				if sc.stash[j] != nil {
					y = sc.stash[j].([]float64)
				}
				sc.stash[j] = append(y, x...)
			case []string:
				var y []string
				if sc.stash[j] != nil {
					y = sc.stash[j].([]string)
				}
				sc.stash[j] = append(y, x...)
			default:
				panic("unkown type")
			}
		}
	}
}
