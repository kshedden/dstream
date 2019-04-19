package dstream

import (
	"bytes"
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

	var buf bytes.Buffer

	Save(data, &buf)
	da := NewLoad(&buf)
	db := MemCopy(da, false)

	if db.NumVar() != 3 {
		t.Fail()
	}

	if !EqualReport(data, db, true) {
		t.Fail()
	}

	if db.NumObs() != 8 {
		t.Fail()
	}

	db.Reset()
	for db.Next() {
		for j := 0; j < db.NumVar(); j++ {
			u := db.GetPos(j)
			v := db.Get(db.Names()[j])
			if !equalSlice(u, v) {
				t.Fail()
			}
		}
	}
}
