package dstream

import "testing"

func TestCenter1(t *testing.T) {

	x0 := []interface{}{
		[]float64{5, 8, 1},
		[]float64{2, 1, 3, 1},
	}
	x1 := []interface{}{
		[]float64{5, 8, 1},
		[]float64{2, 1, 3, 1},
	}
	x2 := []interface{}{
		[]float64{2, 3, 4},
		[]float64{1, 2, 1, 1},
	}
	x3 := []interface{}{
		[]float64{4, 1, 2},
		[]float64{0, 1, 1, 0},
	}
	da := NewFromArrays([][]interface{}{x0, x1, x2, x3}, []string{"x0", "x1", "x2", "x3"})

	dx := Center(da, "x1", "x2")

	x1 = []interface{}{
		[]float64{2, 5, -2},
		[]float64{-1, -2, 0, -2},
	}
	x2 = []interface{}{
		[]float64{0, 1, 2},
		[]float64{-1, 0, -1, -1},
	}
	db := NewFromArrays([][]interface{}{x0, x1, x2, x3}, []string{"x0", "x1", "x2", "x3"})

	for j := 0; j < 2; j++ {
		if !EqualReport(dx, db, true) {
			t.Fail()
		}
		dx.Reset()
		dx = MemCopy(dx)
	}
}
