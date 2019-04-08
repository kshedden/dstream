package dstream

import "testing"

func TestLinapply1(t *testing.T) {

	x := [][]interface{}{
		{
			[]float64{0, 0, 0},
			[]float64{1, 1, 1},
			[]float64{2, 2, 3},
		},
		{
			[]string{"a", "b", "c"},
			[]string{"d", "e", "f"},
			[]string{"g", "h", "i"},
		},
		{
			[]float64{1, 2, 3},
			[]float64{4, 5, 6},
			[]float64{7, 8, 9},
		},
	}
	na := []string{"x1", "x2", "x3"}
	da := NewFromArrays(x, na)

	x = append(x,
		[][]interface{}{
			{
				[]float64{2, 4, 6},
				[]float64{9, 11, 13},
				[]float64{16, 18, 21},
			},
			{
				[]float64{1, 2, 3},
				[]float64{3, 4, 5},
				[]float64{5, 6, 6},
			},
		}...)

	na = []string{"x1", "x2", "x3", "z0", "z1"}
	ex := NewFromArrays(x, na)

	coeffs := [][]float64{
		{1, 0, 2},
		{-1, 1, 1},
	}
	db := Linapply(da, coeffs, "z")

	if !EqualReport(db, ex, true) {
		t.Fail()
	}
}
