package dstream

import "testing"

func TestMutate1(t *testing.T) {

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

	x3t := []interface{}{
		[]float64{3, 5, 7},
		[]float64{9, 11, 13},
		[]float64{15, 17, 19},
	}
	dat = [][]interface{}{x1, x2, x3t}
	na = []string{"x1", "x2", "x3"}
	de := NewFromArrays(dat, na)

	f := func(x interface{}) {
		z := x.([]float64)
		for i, y := range z {
			z[i] = 2*y + 1
		}
	}

	dm := Mutate(da, "x3", f)

	if !EqualReport(dm, de, true) {
		t.Fail()
	}
}
