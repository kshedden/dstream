package dstream

type maxchunksize struct {
	xform

	// Maximum chunk size
	size int

	// Position in the current chunk
	pos int

	// If true, need to call source.Next before proceeding
	first bool
}

// MaxChunkSize splits the chunks of the input Dstream so that no
// chunk has more than size rows.
func MaxChunkSize(data Dstream, size int) Dstream {

	sc := &maxchunksize{
		xform: xform{
			source: data,
		},
		size:  size,
		first: true,
	}

	// Read the first source chunk.
	sc.source.Next()

	nvar := sc.source.NumVar()
	sc.names = sc.source.Names()
	sc.bdata = make([]interface{}, nvar)

	return sc
}

func (sc *maxchunksize) Next() bool {

	if sc.first {
		if !sc.source.Next() {
			return false
		}
		sc.first = false
	}

	// Advance or return if there is no current data
	x := sc.source.GetPos(0)
	m := ilen(x) - sc.pos
	if m <= 0 {
		if !sc.source.Next() {
			return false
		}
		sc.pos = 0
		x = sc.source.GetPos(0)
		m = ilen(x) - sc.pos
		if m <= 0 {
			return false
		}
	}
	if m > sc.size {
		m = sc.size
	}

	for j := 0; j < sc.NumVar(); j++ {

		x := sc.source.GetPos(j)

		// TODO generic types
		switch x := x.(type) {
		case []float64:
			sc.bdata[j] = x[sc.pos : sc.pos+m]
		case []string:
			sc.bdata[j] = x[sc.pos : sc.pos+m]
		default:
			panic("unkown type")
		}
	}

	sc.pos += m

	return true
}

func (sc *maxchunksize) Reset() {
	sc.source.Reset()
	sc.first = true
	sc.pos = 0
}
