package dstream

import "testing"

func TestMaxChunkSize(t *testing.T) {

	x1 := []interface{}{
		[]float64{0, 0, 0, 0, 0, 1, 1, 2, 3},
	}
	x2 := []interface{}{
		[]float64{1, 1, 1, 1, 2, 2, 4, 5, 6},
	}
	x3 := []interface{}{
		[]float64{2, 2, 3, 3, 3, 3, 7, 8, 9},
	}
	x4 := []interface{}{
		[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
	}
	dat := [][]interface{}{x1, x2, x3, x4}
	na := []string{"x1", "x2", "x3", "x4"}
	da := NewFromArrays(dat, na)

	x1 = []interface{}{
		[]float64{0, 0, 0, 0},
		[]float64{0, 1, 1, 2},
		[]float64{3},
	}
	x2 = []interface{}{
		[]float64{1, 1, 1, 1},
		[]float64{2, 2, 4, 5},
		[]float64{6},
	}
	x3 = []interface{}{
		[]float64{2, 2, 3, 3},
		[]float64{3, 3, 7, 8},
		[]float64{9},
	}
	x4 = []interface{}{
		[]string{"a", "b", "c", "d"},
		[]string{"e", "f", "g", "h"},
		[]string{"i"},
	}
	dat = [][]interface{}{x1, x2, x3, x4}
	na = []string{"x1", "x2", "x3", "x4"}
	de := NewFromArrays(dat, na)

	dx := MaxChunkSize(da, 4)

	if !EqualReport(dx, de, true) {
		t.Fail()
	}
}
