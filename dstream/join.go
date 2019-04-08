package dstream

import (
	"fmt"
	"os"
)

// Join performs a streaming join on several Dstreams that have been
// segmented by id variables.  If join has type Join, then
// join.Data[i] is the current chunk of the i^th stream.  All streams
// being joined must have been segmented by an id variable whose
// values are ascending.  The id variable must have type uint64.
//
// A call to the Next method always advances the first stream
// (Data[0]) by one chunk.  The other elements of Data are advanced
// until their id variable is equal to (if possible) or greater than
// the id variable of Data[0].  If equality is achieved, the
// corresponding element of join.status is set to true.
// join.Status[0] is always false and has no meaning.
//
// The dstream values to be joined must be segmented so that the id
// variable is constant within chunks, and increases in numeric value
// with subsequent calls to the Next method.
type Join struct {

	// A sequence of segmented Dstreams to advance in unison.
	Data []Dstream

	// Status[j] means that the id variable for Data value j is
	// equal to the id variable for Data value 0.  Status[0] is
	// not used.
	Status []bool

	inames []string

	ipos []int

	id [][]uint64
}

// NewJoin creates a Join of the given Dstreams, using the variable
// names in names as ids. The Dstreams in data must be segmented by
// the inames variables before calling NewJoin.
func NewJoin(data []Dstream, names []string) *Join {

	if len(data) != len(names) {
		panic("NewJoin: data and names should have same length\n")
	}

	w := &Join{
		Data:   data,
		inames: names,
	}

	for k, da := range data {
		na := da.Names()
		f := false
		for j, na := range na {
			if na == names[k] {
				w.ipos = append(w.ipos, j)
				f = true
				break
			}
		}
		if !f {
			msg := fmt.Sprintf("Can't find index variable %s in %dth dataset", names[k], k)
			os.Stderr.WriteString(msg)
			os.Exit(1)
		}
	}

	w.Status = make([]bool, len(data))
	w.id = make([][]uint64, len(data))

	return w
}

func (w *Join) needsadvance(j int) bool {
	if w.id[j] == nil || len(w.id[j]) == 0 {
		return true
	}

	if w.id[j][0] < w.id[0][0] {
		return true
	}

	return false
}

func (w *Join) clearstatus() {
	for j := range w.Status {
		w.Status[j] = false
	}
}

// Next advances to the next chunk.  The first dstream, which is
// contained in join.Data[0], always advances to the next sequential
// value of its id variable.  The other dstreams (join.Data[j] for j >
// 0) advance until their id variables are equal to or greater than
// the id variable for the current chunk of join.Data[0].  The status
// field (join.status) indicates which dstreams in the join are
// currently on the same id value as the first dstream (join.Data[0]).
func (w *Join) Next() bool {

	// Advance the index stream
	f := w.Data[0].Next()
	if !f {
		return false
	}
	w.id[0] = w.Data[0].GetPos(w.ipos[0]).([]uint64)

	// Advance the other streams
	w.clearstatus()
	for j := 1; j < len(w.Data); j++ {

		// Keep advancing as long as the stream is behind the index stream.
		for w.needsadvance(j) {
			f := w.Data[j].Next()
			if !f {
				break
			}
			w.id[j] = w.Data[j].GetPos(w.ipos[j]).([]uint64)
		}

		w.Status[j] = w.id[j][0] == w.id[0][0]
	}

	return true
}
