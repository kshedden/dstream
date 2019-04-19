package dstream

import "testing"

func TestMemCopy1(t *testing.T) {

	x := [][]interface{}{
		{
			[]float64{0, 1, 1},
			[]float64{0, 0, 1, 0},
		},
		{
			[]float64{1, 1, 1},
			[]float64{1, 1, 1, 1},
		},
		{
			[]float64{4, 1, -1},
			[]float64{3, 5, -5, 3},
		},
		{
			[]float64{1, 1, 1},
			[]float64{2, 2, 1, 1},
		},
		{
			[]float64{1, 1, 1},
			[]float64{2, 2, 2, 3},
		},
	}

	na := []string{"x1", "x2", "x3", "x4", "x5"}
	da := NewFromArrays(x, na)

	db := MemCopy(da, true)

	if !EqualReport(da, db, true) {
		t.Fail()
	}
}
