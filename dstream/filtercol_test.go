package dstream

import (
	"fmt"
	"testing"
)

func dataf1() (Dstream, Dstream) {
	x1 := []interface{}{
		[]float64{0, 0, 0, 0, 0, 0},
		[]float64{1, 0, 1, 2, 3, 0},
		[]float64{2, 4, 6, 0, 10, 12},
	}
	x2 := []interface{}{
		[]float64{0, 0, 1, 2, 3, 4},
		[]float64{1, 2, 2, 2, 2, 1},
		[]float64{3, 3, 3, 3, 3, 3},
	}
	x3 := []interface{}{
		[]float64{1, 2, 3, 4, 5, 6},
		[]float64{4, 5, 6, 7, 8, 9},
		[]float64{7, 8, 9, 10, 11, 12},
	}
	x4 := []interface{}{
		[]string{"a", "b", "c", "d", "e", "f"},
		[]string{"d", "e", "f", "g", "h", "i"},
		[]string{"g", "h", "i", "j", "k", "l"},
	}
	dat := [][]interface{}{x1, x2, x3, x4}
	na := []string{"x1", "x2", "x3", "x4"}
	da := NewFromArrays(dat, na)

	x1 = []interface{}{
		[]float64{},
		[]float64{1, 1, 2, 3},
		[]float64{2, 4, 6, 10, 12},
	}
	x2 = []interface{}{
		[]float64{},
		[]float64{1, 2, 2, 2},
		[]float64{3, 3, 3, 3, 3},
	}
	x3 = []interface{}{
		[]float64{},
		[]float64{4, 6, 7, 8},
		[]float64{7, 8, 9, 11, 12},
	}
	x4 = []interface{}{
		[]string{},
		[]string{"d", "f", "g", "h"},
		[]string{"g", "h", "i", "k", "l"},
	}
	dat = [][]interface{}{x1, x2, x3, x4}
	na = []string{"x1", "x2", "x3", "x4"}
	db := NewFromArrays(dat, na)

	return da, db
}

func TestFilterCol1(t *testing.T) {

	da, db := dataf1()

	// Keep if nonzero
	f1 := func(x interface{}, keep []bool) bool {
		y := x.([]float64)
		any := false
		for i, z := range y {
			keep[i] = true
			if z == 0 {
				keep[i] = false
				any = true
			}
		}
		return any
	}

	dx := FilterCol(da, map[string]FilterColFunc{"x1": f1})

	for k := 0; k < 2; k++ {
		if !EqualReport(dx, db, true) {
			t.Fail()
		}
		if dx.NumObs() != db.NumObs() {
			fmt.Printf("%d != %d\n", dx.NumObs(), db.NumObs())
			t.Fail()
		}
		dx.Reset()
	}
}

func TestFilterCol2(t *testing.T) {

	da, _ := dataf1()

	f1 := func(x interface{}, keep []bool) bool {
		any := false
		y := x.([]float64)
		for i, z := range y {
			if z != 99 {
				keep[i] = true
				any = true
			}
		}
		return any
	}

	dx := FilterCol(da, map[string]FilterColFunc{"x1": f1})

	if !EqualReport(dx, MemCopy(da), true) {
		t.Fail()
	}
}
