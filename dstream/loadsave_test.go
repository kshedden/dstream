package dstream

import (
	"testing"
)

func TestLoadSave(t *testing.T) {

	x := [][]interface{}{
		{
			[]float64{1, 2, 3, 4, 5},
			[]float64{2, 9, 6},
		},
		{
			[]uint64{1, 0, 2, 5, 4},
			[]uint64{7, 4, 9},
		},
		{
			[]string{"a", "b", "c", "d", "e"},
			[]string{"f", "g", "h"},
		},
	}

	names := []string{"x", "y", "z"}
	data := NewFromArrays(x, names)

	Save(data, "tmp.gz")

	da := NewLoad("tmp.gz")

	if da.NumVar() != 3 {
		t.Fail()
	}

	for k := 0; k < 2; k++ {
		if !EqualReport(data, da, true) {
			t.Fail()
		}

		if da.NumObs() != 8 {
			t.Fail()
		}

		da.Reset()
	}

	da.Reset()

	for da.Next() {
		for j := 0; j < da.NumObs(); j++ {
			u := da.GetPos(j)
			v := da.Get(da.Names()[j])
			if !equalSlice(u, v) {
				t.Fail()
			}
		}
	}
}
