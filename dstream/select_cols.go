package dstream

import "fmt"

type selectCols struct {
	xform

	keepVars []string
	keepPos  []int
}

// SelectCols retains only the given variables in a Dstream.
func SelectCols(data Dstream, keepvars ...string) Dstream {
	sl := &selectCols{
		xform: xform{
			source: data,
		},
		keepVars: keepvars,
	}

	vn := make(map[string]int)
	for k, na := range data.Names() {
		vn[na] = k
	}

	for _, na := range keepvars {
		pos, ok := vn[na]
		if !ok {
			panic(fmt.Sprintf("Variable '%s' not found in dstream.", na))
		}
		sl.keepPos = append(sl.keepPos, pos)
	}

	sl.names = keepvars
	sl.bdata = make([]interface{}, len(sl.keepVars))

	return sl
}

func (sl *selectCols) Next() bool {

	if !sl.source.Next() {
		return false
	}

	for j, k := range sl.keepPos {
		sl.bdata[j] = sl.source.GetPos(k)
	}

	return true
}
