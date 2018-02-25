package dstream

type concatVertical struct {

	// The streams to be concatenated
	streams []Dstream

	// The index of the current stream within streams
	pos int

	// The number of observations in the concatenated stream
	nobs int

	// True if nobs is known yet (nobs is not known until reading
	// the entire concatenated stream).
	nobsKnown bool

	// Map from variable names to column positions
	namepos map[string]int
}

// ConcatVertical concatenates a collection of Dstreams vertically
// (appending additional observations).  The column names and data
// types of all the Dstreams being combined must be identical.
func ConcatVertical(streams ...Dstream) Dstream {

	c := &concatVertical{
		streams: streams,
	}

	// Construct the name to position mapping
	c.namepos = make(map[string]int)
	for k, n := range streams[0].Names() {
		c.namepos[n] = k
	}

	return c
}

func (c *concatVertical) Close() {

	for _, s := range c.streams {
		s.Close()
	}
}

func (c *concatVertical) GetPos(pos int) interface{} {
	return c.streams[c.pos].GetPos(pos)
}

func (c *concatVertical) NumObs() int {
	if c.nobsKnown {
		return c.nobs
	}
	return -1
}

func (c *concatVertical) NumVar() int {
	return len(c.Names())
}

func (c *concatVertical) Names() []string {
	return c.streams[0].Names()
}

func (c *concatVertical) Get(name string) interface{} {
	return c.GetPos(c.namepos[name])
}

func (c *concatVertical) Next() bool {

	// Advance within current stream
	if c.streams[c.pos].Next() {
		return true
	}

	// Try to advance to next stream
	c.nobs += c.streams[c.pos].NumObs()
	c.pos++
	if c.pos >= len(c.streams) {
		// Done with all streams
		c.nobsKnown = true
		return false
	}

	// Advance to next stream
	c.streams[c.pos].Next()

	// Check that the names are the same
	a := c.streams[0].Names()
	b := c.streams[c.pos].Names()
	msg := "Streams to be concatenated have different column names"
	if len(a) != len(b) {
		panic(msg)
	}
	for i := range a {
		if a[i] != b[i] {
			panic(msg)
		}
	}

	return true
}

func (c *concatVertical) Reset() {

	c.pos = 0
	c.nobs = 0
	c.nobsKnown = false

	for _, s := range c.streams {
		s.Reset()
	}
}
