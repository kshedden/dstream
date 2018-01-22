package dstream

import "testing"

func TestConcat1(t *testing.T) {

	x1 := []interface{}{
		[]float64{0, 1, 1},
		[]float64{0, 0, 1, 0},
	}
	x2 := []interface{}{
		[]float64{1, 1, 1},
		[]float64{1, 1, 1, 1},
	}
	dat := [][]interface{}{x1, x2}
	na := []string{"x1", "x2"}
	da := NewFromArrays(dat, na)

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
	dat = [][]interface{}{x1, x2}
	de := NewFromArrays(dat, na)

	db := MemCopy(da)
	da.Reset()

	dq := Concat([]Dstream{da, db})

	if !EqualReport(dq, de, true) {
		t.Fail()
	}
}
