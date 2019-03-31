package dstream

import "testing"

func TestAddcol1(t *testing.T) {

	x := [][]interface{}{
		{
			[]float64{0, 0, 0},
			[]float64{1, 1},
			[]float64{2, 2, 2, 2},
		},
		{
			[]string{"a", "a", "a"},
			[]string{"b", "b"},
			[]string{"c", "c", "c", "c"},
		},
	}
	na := []string{"x1", "x2"}
	da := NewFromArrays(x, na)

	x = append(x,
		[]interface{}{
			[]float64{1, 2, 3},
			[]float64{4, 5},
			[]float64{6, 7, 8, 9},
		})
	na = append(na, "x3")
	db := NewFromArrays(x, na)

	z3 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	dm := AddCol(da, z3, "x3")

	if !EqualReport(dm, db, true) {
		t.Fail()
	}
}
