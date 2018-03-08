package dstream

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

type center struct {
	source Dstream

	// Positions of variables to center
	pos []int

	// Map from variable name to variable position
	vpos map[string]int

	// Query if something is in pos
	cenpos map[int]int

	// Means of variables to center
	means []float64
}

func (c *center) Close() {
	c.source.Close()
}

func (c *center) Reset() {
	c.source.Reset()
}

// Center returns a new Dstream in which the given columns have been
// mean-centered.  Currently only works with float64 type data.
func Center(source Dstream, names ...string) Dstream {

	means := make([]float64, len(names))

	// Map from variable names to position
	vpos := make(map[string]int)
	for k, na := range source.Names() {
		vpos[na] = k
	}

	// Positions of variables to center.  cenpos[j], if it exists,
	// is the position within means where the mean of the variable
	// j is found.  If cenpos[j] does not exist, the variable
	// should not be centered.
	var pos []int
	cenpos := make(map[int]int)
	for j, na := range names {
		q, ok := vpos[na]
		if !ok {
			msg := fmt.Sprintf("Center: variable '%s' not found.\n", na)
			panic(msg)
		}
		pos = append(pos, q)
		cenpos[q] = j
	}

	// Get the means of the variables to center
	source.Reset()
	n := 0
	for source.Next() {
		for i, j := range pos {
			v := source.GetPos(j).([]float64)
			if i == 0 {
				n += len(v)
			}
			means[i] += floats.Sum(v)
		}
	}
	floats.Scale(1/float64(n), means)

	return &center{
		source: source,
		pos:    pos,
		means:  means,
		cenpos: cenpos,
		vpos:   vpos,
	}
}

func (c *center) GetPos(j int) interface{} {

	x := c.source.GetPos(j).([]float64)

	q, ok := c.cenpos[j]
	if !ok {
		return x
	}

	m := c.means[q]

	z := make([]float64, len(x))
	for i := range x {
		z[i] = x[i] - m
	}

	return z
}

func (c *center) Names() []string {
	return c.source.Names()
}

func (c *center) Next() bool {
	return c.source.Next()
}

func (c *center) NumObs() int {
	return c.source.NumObs()
}

func (c *center) NumVar() int {
	return c.source.NumVar()
}

func (c *center) Get(na string) interface{} {
	return c.GetPos(c.vpos[na])
}
