package dstream

import "fmt"

type addcol struct {
	xform
	newname string
	pos     int
	newdat  []float64
	first   bool
}

// TODO make generic

// Addcol appends a new column of data to a Dstream.
func Addcol(da Dstream, newdat []float64, newname string) Dstream {

	r := &addcol{
		xform: xform{
			source: da,
		},
		newname: newname,
		newdat:  newdat,
	}

	for _, na := range r.source.Names() {
		if newname == na {
			msg := fmt.Sprintf("Addcol: a variable named '%s' already exists.", na)
			panic(msg)
		}
	}

	r.names = append(r.names, r.source.Names()...)
	r.names = append(r.names, r.newname)
	r.first = true
	r.bdata = make([]interface{}, len(r.names))

	return r
}

func (ac *addcol) Reset() {
	ac.pos = 0
	ac.first = true
	ac.source.Reset()
}

func (ac *addcol) Next() bool {

	if !ac.first {
		ac.pos += ilen(ac.bdata[0])
	}
	ac.first = false
	f := ac.source.Next()
	if !f {
		return false
	}
	for k := 0; k < ac.source.NumVar(); k++ {
		ac.bdata[k] = ac.source.GetPos(k)
	}
	csize := ilen(ac.bdata[0])

	ac.bdata[len(ac.bdata)-1] = ac.newdat[ac.pos : ac.pos+csize]

	return true
}
