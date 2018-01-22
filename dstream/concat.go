package dstream

type concat struct {
	streams []Dstream

	pos int

	nobs int

	nobsKnown bool

	namepos map[string]int
}

// Concat concatenates a collection of Dstreams.  The columns names of
// types of all the Dstreams being combined should be identical.
func Concat(streams []Dstream) Dstream {

	c := &concat{
		streams: streams,
	}

	c.namepos = make(map[string]int)
	for k, n := range streams[0].Names() {
		c.namepos[n] = k
	}

	return c
}

func (c *concat) Close() {

	for _, s := range c.streams {
		s.Close()
	}
}

func (c *concat) GetPos(pos int) interface{} {
	return c.streams[c.pos].GetPos(pos)
}

func (c *concat) NumObs() int {
	if c.nobsKnown {
		return c.nobs
	}
	return -1
}

func (c *concat) NumVar() int {
	return len(c.Names())
}

func (c *concat) Names() []string {
	return c.streams[0].Names()
}

func (c *concat) Get(name string) interface{} {
	return c.GetPos(c.namepos[name])
}

func (c *concat) Next() bool {

	if c.streams[c.pos].Next() {
		return true
	}

	c.nobs += c.streams[c.pos].NumObs()
	c.pos++
	if c.pos >= len(c.streams) {
		c.nobsKnown = true
		return false
	}

	c.streams[c.pos].Next()
	return true
}

func (c *concat) Reset() {

	c.pos = 0
	c.nobs = 0
	c.nobsKnown = false

	for _, s := range c.streams {
		s.Reset()
	}
}
