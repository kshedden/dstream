package dstream

import (
	"fmt"
	"testing"
)

func dataf1() (Dstream, Dstream) {
	x := [][]interface{}{
		{
			[]float64{0, 0, 0, 0, 0, 0},
			[]float64{1, 0, 1, 2, 3, 0},
			[]float64{2, 4, 6, 0, 10, 12},
		},
		{
			[]float64{0, 0, 1, 2, 3, 4},
			[]float64{1, 2, 2, 2, 2, 1},
			[]float64{3, 3, 3, 3, 3, 3},
		},
		{
			[]float64{1, 2, 3, 4, 5, 6},
			[]float64{4, 5, 6, 7, 8, 9},
			[]float64{7, 8, 9, 10, 11, 12},
		},
		{
			[]string{"a", "b", "c", "d", "e", "f"},
			[]string{"d", "e", "f", "g", "h", "i"},
			[]string{"g", "h", "i", "j", "k", "l"},
		},
	}
	na := []string{"x1", "x2", "x3", "x4"}
	da := NewFromArrays(x, na)

	x = [][]interface{}{
		{
			[]float64{},
			[]float64{1, 1, 2, 3},
			[]float64{2, 4, 6, 10, 12},
		},
		{
			[]float64{},
			[]float64{1, 2, 2, 2},
			[]float64{3, 3, 3, 3, 3},
		},
		{
			[]float64{},
			[]float64{4, 6, 7, 8},
			[]float64{7, 8, 9, 11, 12},
		},
		{
			[]string{},
			[]string{"d", "f", "g", "h"},
			[]string{"g", "h", "i", "k", "l"},
		},
	}
	na = []string{"x1", "x2", "x3", "x4"}
	db := NewFromArrays(x, na)

	return da, db
}

func TestFilter1(t *testing.T) {

	da, db := dataf1()

	// Keep if x1 is not zero.
	f := func(v map[string]interface{}, keep []bool) {
		x := v["x1"].([]float64)
		for i, z := range x {
			keep[i] = keep[i] && (z != 0)
		}
	}

	dx := Filter(da, f)

	// Process twice to make sure that Reset works
	for k := 0; k < 2; k++ {

		// Check the data values
		if !EqualReport(dx, db, true) {
			t.Fail()
		}

		// Check the number of observations
		if dx.NumObs() != db.NumObs() {
			fmt.Printf("NumObs mismatch: %d != %d\n", dx.NumObs(), db.NumObs())
			t.Fail()
		}

		dx.Reset()
	}
}

func TestFilter2(t *testing.T) {

	da, _ := dataf1()

	f := func(v map[string]interface{}, keep []bool) {
		x := v["x1"].([]float64)
		for i, z := range x {
			keep[i] = keep[i] && (z != 99)
		}
	}

	dx := Filter(da, f)

	if !EqualReport(dx, MemCopy(da), true) {
		t.Fail()
	}
}
