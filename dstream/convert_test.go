package dstream

import "testing"

func TestConvert(t *testing.T) {

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
		[]int32{0, 0, 0},
		[]int32{1, 1, 1},
		[]int32{2, 2, 3},
	}
	z2 := []interface{}{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
	}
	z3 := []interface{}{
		[]float32{1, 2, 3},
		[]float32{4, 5, 6},
		[]float32{7, 8, 9},
	}
	dat = [][]interface{}{z1, z2, z3}
	db := NewFromArrays(dat, na)

	dx := Convert(da, "x1", "int32")
	dx = Convert(dx, "x3", "float32")

	if !EqualReport(dx, db, true) {
		t.Fail()
	}

}
