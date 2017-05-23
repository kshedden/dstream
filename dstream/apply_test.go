package dstream

import "testing"

func TestApply1(t *testing.T) {

	x1 := []interface{}{
		[]float64{0, 0, 0},
		[]float64{1, 1, 1},
		[]float64{2, 2, 3},
	}
	x2 := []interface{}{
		[]float64{0, 0, 1},
		[]float64{1, 2, 2},
		[]float64{3, 3, 3},
	}
	x3 := []interface{}{
		[]float64{1, 2, 3},
		[]float64{4, 5, 6},
		[]float64{7, 8, 9},
	}
	x4 := []interface{}{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
	}
	dat := [][]interface{}{x1, x2, x3, x4}
	na := []string{"x1", "x2", "x3", "x4"}
	da := NewFromArrays(dat, na)

	x5 := []interface{}{
		[]float64{0, 0, 0},
		[]float64{1, 2, 2},
		[]float64{6, 6, 9},
	}
	dat = [][]interface{}{x1, x2, x3, x4, x5}
	na = []string{"x1", "x2", "x3", "x4", "x5"}
	ex := NewFromArrays(dat, na)

	f := func(v map[string]interface{}, x interface{}) {
		z := x.([]float64)

		x1 := v["x1"].([]float64)
		x2 := v["x2"].([]float64)

		for i, _ := range x1 {
			z[i] = x1[i] * x2[i]
		}
	}

	db := Apply(da, "x5", f, "float64")

	if !EqualReport(db, ex, true) {
		t.Fail()
	}
}
