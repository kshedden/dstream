package dstream

import "testing"

func TestConcatHoriz1(t *testing.T) {

	x1 := []interface{}{
		[]float64{0, 1, 1},
		[]float64{0, 0, 1, 0},
	}
	x2 := []interface{}{
		[]float64{1, 1, 1},
		[]float64{1, 1, 1, 1},
	}
	da := NewFromArrays([][]interface{}{x1, x2}, []string{"x1", "x2"})

	x3 := []interface{}{
		[]float64{2, 3, 4},
		[]float64{5, 6, 7, 8},
	}
	x4 := []interface{}{
		[]float64{9, 8, 7},
		[]float64{6, 5, 4, 3},
	}
	db := NewFromArrays([][]interface{}{x3, x4}, []string{"x3", "x4"})

	de := NewFromArrays([][]interface{}{x1, x2, x3, x4}, []string{"x1", "x2", "x3", "x4"})

	dq := ConcatHorizontal(da, db)

	if !EqualReport(dq, de, true) {
		t.Fail()
	}
}
