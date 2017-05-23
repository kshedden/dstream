package dstream

import "testing"

func TestLinapply1(t *testing.T) {

	x1 := []interface{}{
		[]float64{0, 0, 0},
		[]float64{1, 1, 1},
		[]float64{2, 2, 3},
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
		[]float64{2, 4, 6},
		[]float64{9, 11, 13},
		[]float64{16, 18, 21},
	}

	z2 := []interface{}{
		[]float64{1, 2, 3},
		[]float64{3, 4, 5},
		[]float64{5, 6, 6},
	}

	dat = [][]interface{}{x1, x2, x3, z1, z2}
	na = []string{"x1", "x2", "x3", "z0", "z1"}
	ex := NewFromArrays(dat, na)

	coeffs := [][]float64{
		[]float64{1, 0, 2},
		[]float64{-1, 1, 1},
	}
	db := Linapply(da, coeffs, "z")

	if !EqualReport(db, ex, true) {
		t.Fail()
	}
}
