package dstream

import "testing"

func TestConcatVertical1(t *testing.T) {

	x1 := []interface{}{
		[]float64{0, 1, 1},
		[]float64{0, 0, 1, 0},
	}
	x2 := []interface{}{
		[]float64{1, 1, 1},
		[]float64{1, 1, 1, 1},
	}
	da := NewFromArrays([][]interface{}{x1, x2}, []string{"x1", "x2"})

	x1 = []interface{}{
		[]float64{0, 1, 1},
		[]float64{0, 0, 1, 0},
		[]float64{0, 1, 1},
		[]float64{0, 0, 1, 0},
	}
	x2 = []interface{}{
		[]float64{1, 1, 1},
		[]float64{1, 1, 1, 1},
		[]float64{1, 1, 1},
		[]float64{1, 1, 1, 1},
	}
	de := NewFromArrays([][]interface{}{x1, x2}, []string{"x1", "x2"})

	db := MemCopy(da, true)
	da.Reset()

	dq := ConcatVertical(da, db)

	if !EqualReport(dq, de, true) {
		t.Fail()
	}
}
