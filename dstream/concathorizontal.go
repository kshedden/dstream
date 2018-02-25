package dstream

import "fmt"

type concatHorizontal struct {

	// The streams to be concatenated
	streams []Dstream

	// All names in all streams (in the same order used by
	// GetPos).
	names []string

	// The index of the current stream within streams
	pos int

	// The number of observations in the concatenated stream
	nobs int

	// True if nobs is known yet (nobs is not known until reading
	// the entire concatenated stream).
	nobsKnown bool

	// Map from variable names to stream index
	namestream map[string]int

	// Map from variable names to column position within its stream
	namepos map[string]int

	// Map from external position to stream index
	posstream map[int]int

	// Map from external position to column position within its stream
	pospos map[int]int
}

// ConcatHorizontal concatenates a collection of Dstreams
// horizontally.  The column names of all the Dstreams being combined
// must be distinct.
func ConcatHorizontal(streams ...Dstream) Dstream {

	c := &concatHorizontal{
		streams: streams,
	}

	haveName := make(map[string]bool)

	// Construct the name to position mapping
	c.namepos = make(map[string]int)
	c.namestream = make(map[string]int)
	c.posstream = make(map[int]int)
	c.pospos = make(map[int]int)
	pos := 0
	for j, stream := range streams {
		for k, na := range stream.Names() {

			if haveName[na] {
				msg := fmt.Sprintf("Name '%s' is not unique\n", na)
				panic(msg)
			}
			haveName[na] = true

			c.namestream[na] = j
			c.namepos[na] = k
			c.posstream[pos] = j
			c.pospos[pos] = k
			c.names = append(c.names, na)
			pos++
		}
	}

	return c
}

func (c *concatHorizontal) Close() {

	for _, s := range c.streams {
		s.Close()
	}
}

func (c *concatHorizontal) GetPos(pos int) interface{} {
	return c.streams[c.posstream[pos]].GetPos(c.pospos[pos])
}

func (c *concatHorizontal) NumObs() int {
	if c.nobsKnown {
		return c.nobs
	}
	return -1
}

func (c *concatHorizontal) NumVar() int {
	return len(c.Names())
}

func (c *concatHorizontal) Names() []string {
	return c.names
}

func (c *concatHorizontal) Get(name string) interface{} {
	return c.streams[c.namestream[name]].GetPos(c.namepos[name])
}

func (c *concatHorizontal) Next() bool {

	// Advance all the strings
	var n1, n0 int
	for _, s := range c.streams {
		if s.Next() {
			n1++
		} else {
			n0++
		}
	}

	if n0 == 0 {
		// All streams advanced
		return true
	}

	if n0 > 0 && n1 > 0 {
		panic("Streams have different lengths\n")

	}

	return false
}

func (c *concatHorizontal) Reset() {

	c.pos = 0
	c.nobs = 0
	c.nobsKnown = false

	for _, s := range c.streams {
		s.Reset()
	}
}
