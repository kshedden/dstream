package dstream

import (
	"fmt"
	"sort"
)

//go:generate go run gen.go -template=regroup.template

type argsort struct {
	s    []uint64
	inds []int
}

func (a argsort) Len() int {
	return len(a.s)
}

func (a argsort) Less(i, j int) bool {
	return a.s[i] < a.s[j]
}

func (a argsort) Swap(i, j int) {
	a.s[i], a.s[j] = a.s[j], a.s[i]
	a.inds[i], a.inds[j] = a.inds[j], a.inds[i]
}

// Regroup creates a new Dstream from the provided Dstream having
// identical rows, but with the chunks defined by the values of a
// provided id variable.  The resulting Dstream will have a chunk for
// each distinct level of the id variable, containing all the rows of
// the input Dstream with the given id value.  The id variable must
// have uint64 type.
func Regroup(ds Dstream, groupvar string, sortchunks bool) Dstream {

	// Find the variable's position
	idpos := -1
	for j, n := range ds.Names() {
		if n == groupvar {
			idpos = j
			break
		}
	}
	if idpos == -1 {
		msg := fmt.Sprintf("Regroup: variable '%s' not found", groupvar)
		panic(msg)
	}

	r := doRegroup(ds, idpos)

	if sortchunks {
		var x []uint64
		for _, v := range r.data[idpos] {
			y := v.([]uint64)
			x = append(x, y[0])
		}
		ii := make([]int, len(x))
		for i := range ii {
			ii[i] = i
		}
		a := argsort{x, ii}
		sort.Sort(a)

		nchunk := len(r.data[0])
		newar := make([][]interface{}, len(r.data))
		for j := 0; j < ds.NumVar(); j++ {
			newar[j] = make([]interface{}, nchunk)
			for k, i := range ii {
				newar[j][k] = r.data[j][i]
			}
		}
		r.data = newar
	}

	return r
}
