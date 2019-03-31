package dstream

import "testing"

func TestConcatHoriz1(t *testing.T) {

	x := [][]interface{}{
		{
			[]float64{0, 1, 1},
			[]float64{0, 0, 1, 0},
		},
		{
			[]float64{1, 1, 1},
			[]float64{1, 1, 1, 1},
		},
	}
	da := NewFromArrays(x, []string{"x1", "x2"})

	y := [][]interface{}{
		{
			[]float64{2, 3, 4},
			[]float64{5, 6, 7, 8},
		},
		{
			[]float64{9, 8, 7},
			[]float64{6, 5, 4, 3},
		},
	}
	db := NewFromArrays(y, []string{"x3", "x4"})

	z := [][]interface{}{x[0], x[1], y[0], y[1]}
	de := NewFromArrays(z, []string{"x1", "x2", "x3", "x4"})

	dq := ConcatHorizontal(da, db)

	if !EqualReport(dq, de, true) {
		t.Fail()
	}
}
