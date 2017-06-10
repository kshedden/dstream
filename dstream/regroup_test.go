package dstream

import "testing"

func TestRegroup1(t *testing.T) {

	x1 := []interface{}{
		[]uint64{0, 1, 3},
		[]uint64{1, 2, 1},
		[]uint64{1, 2, 0},
	}
	x2 := []interface{}{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
	}
	x3 := []interface{}{
		[]float64{1, 2, 3},
		[]float64{4, 5, 6},
		[]float64{7, 8, 9},
	}
	dat := [][]interface{}{x1, x2, x3}
	na := []string{"x1", "x2", "x3"}
	da := NewFromArrays(dat, na)

	z1 := []interface{}{
		[]uint64{0, 0},
		[]uint64{1, 1, 1, 1},
		[]uint64{2, 2},
		[]uint64{3},
	}
	z2 := []interface{}{
		[]string{"a", "i"},
		[]string{"b", "d", "f", "g"},
		[]string{"e", "h"},
		[]string{"c"},
	}
	z3 := []interface{}{
		[]float64{1, 9},
		[]float64{2, 4, 6, 7},
		[]float64{5, 8},
		[]float64{3},
	}
	dat = [][]interface{}{z1, z2, z3}
	na = []string{"x1", "x2", "x3"}
	db := NewFromArrays(dat, na)

	dr := Regroup(da, "x1", true)

	if !EqualReport(dr, db, true) {
		t.Fail()
	}
}
