package dstream

import "fmt"

type replace struct {

	// The data stream in which we are replacing a column
	source Dstream

	// Name of the column that we are replacing
	name string

	// Column position of the column that we are replacing
	colpos int

	// First row index of the current chunk
	rowpos int

	// Size of the current chunk.
	csize int

	// The data being used to replace the column
	coldata interface{}

	// Map from column names to position
	vpos map[string]int
}

// Replace returns a new Dstream in which the column with the given
// name is replaced with the given data.  The col value must be an
// array type of a valid primitive type (e.g. int, float64, string),
// and must have length equal to the number of rows of data.
func Replace(data Dstream, name string, coldata interface{}) Dstream {

	vpos := make(map[string]int)
	for q, na := range data.Names() {
		vpos[na] = q
	}

	colpos := -1
	for j, na := range data.Names() {
		if na == name {
			colpos = j
			break
		}
	}
	if colpos == -1 {
		msg := fmt.Sprintf("Replace: column '%s' not found.\n", name)
		panic(msg)
	}

	r := &replace{
		source:  data,
		name:    name,
		colpos:  colpos,
		coldata: coldata,
		vpos:    vpos,
	}

	return r
}

func (r *replace) Next() bool {

	if !r.source.Next() {
		return false
	}

	r.rowpos += r.csize
	r.csize = ilen(r.source.GetPos(0))

	return true
}

func (r *replace) Names() []string {
	return r.source.Names()
}

func (r *replace) NumVar() int {
	return r.source.NumVar()
}

func (r *replace) NumObs() int {
	return r.source.NumObs()
}

func (r *replace) Reset() {
	r.rowpos = 0
	r.csize = 0
	r.source.Reset()
}

func (r *replace) Close() {
	r.source.Close()
}

func (r *replace) Get(name string) interface{} {

	j, ok := r.vpos[name]
	if !ok {
		msg := fmt.Sprintf("Replace: variable '%s' not found.\n", name)
		panic(msg)
	}

	return r.GetPos(j)
}
