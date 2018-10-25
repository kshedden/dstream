package dstream

import "fmt"

//go:generate go run gen.go -template=segment.template

type segmentedData struct {
	xform

	// Variables whose changing values define the segments
	vars []string

	// Positions of the variables in vars
	vpos []int

	stash []interface{}

	// Beginning/end of the current segment
	pos int

	// true if there is no more data to read from the source; this
	// does not necessarily mean that there is no more data to
	// yield.
	done bool
}

// Segment restructures the chunks of a Dstream so that chunk
// boundaries are determined by any change in the consecutive values
// of a specified set of variables.
func Segment(data Dstream, vars ...string) Dstream {
	s := &segmentedData{
		xform: xform{
			source: data,
		},
		vars: vars,
	}
	s.init()

	return s
}

func (sd *segmentedData) init() {

	// Get the positions of the variables that define the
	// segments.
	sd.names = sd.source.Names()
	mp := make(map[string]int)
	for k, v := range sd.names {
		mp[v] = k
	}
	var vpos []int
	for _, v := range sd.vars {
		vpos = append(vpos, mp[v])
	}
	sd.vpos = vpos

	nvar := sd.source.NumVar()
	sd.bdata = make([]interface{}, nvar)
	sd.stash = make([]interface{}, nvar)

	sd.source.Next()
	sd.setb()
}

// setb sets the pointers in bdata to point to the current source data.
func (sd *segmentedData) setb() {
	nvar := sd.source.NumVar()
	for j := 0; j < nvar; j++ {
		sd.bdata[j] = sd.source.GetPos(j)
	}
}

func (sd *segmentedData) Get(na string) interface{} {

	if sd.namepos == nil {
		sd.setNamePos()
	}

	pos, ok := sd.namepos[na]
	if !ok {
		msg := fmt.Sprintf("Variable '%s' not found", na)
		panic(msg)
	}
	return sd.GetPos(pos)
}

func (sd *segmentedData) Reset() {
	sd.xform.Reset()
	sd.pos = 0
	sd.source.Next()
	sd.setb()
	sd.done = false
}

func (sd *segmentedData) Next() bool {

	// Stash contains at most one group.
	truncate(sd.stash)

	if ilen(sd.bdata[0]) == 0 && sd.done {
		return false
	} else if sd.done {
		truncate(sd.bdata)
	}

	// Cut off the previous group, and find the end of the current
	// group.
	sd.leftsliceb(sd.pos)
	sd.pos = sd.findSegment(0)

	// Found a complete group in the current chunk.
	if sd.pos != -1 {
		return true
	}

	sd.pos = 0
	sd.setstash()

	// Get a complete group in the stash, keep reading until we
	// have somthing to return.
	for {
		f := sd.source.Next()

		if !f {
			sd.done = true

			// No more data, whatever is left in bdata is
			// the last segment, or we are done if adata
			// is empty.
			if ilen(sd.bdata[0]) == 0 {
				return false
			}
			return true
		}
		sd.setb()

		fd := sd.fixstash()

		if fd {
			return true
		}
	}
}
