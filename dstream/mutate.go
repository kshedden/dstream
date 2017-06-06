package dstream

import "fmt"

type MutateFunc func(interface{})

type mutated struct {
	xform

	// The name of the variable to be mutated
	vname string

	// The position of the variable to be mutated
	vpos int

	// The function that performs the mutation
	f MutateFunc
}

// Mutate returns a Dstream in which the variable with the given name
// is transformed using the given function.
func Mutate(ds Dstream, name string, f MutateFunc) Dstream {

	// Find the variable's position
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

	// Call the mutating function on the variable to be
	// transformed.
	m.f(m.GetPos(m.vpos))

	return true
}
