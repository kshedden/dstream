package dstream

import "fmt"

type mfunc func(interface{})

type mutated struct {
	xform

	vname string
	vpos  int

	f mfunc
}

// Mutate returns a Dstream in which the variable with the given name
// is transformed using the given function.
func Mutate(ds Dstream, name string, f mfunc) Dstream {

	vpos := -1
	for j, n := range ds.Names() {
		if n == name {
			vpos = j
			break
		}
	}
	if vpos == -1 {
		msg := fmt.Sprintf("Mutate: variable '%s' not found", name)
		panic(msg)
	}

	m := &mutated{
		xform: xform{
			source: ds,
		},
		vname: name,
		vpos:  vpos,
		f:     f,
	}

	return m
}

func (m *mutated) Next() bool {
	if !m.source.Next() {
		return false
	}
	m.f(m.GetPos(m.vpos))
	return true
}
