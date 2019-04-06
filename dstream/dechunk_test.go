package dstream

import "testing"

func TestDechunk1(t *testing.T) {

	x := [][]interface{}{
		{
			[]float64{0, 1, 2},
			[]float64{3, 4, 5, 6},
			[]float64{7},
		},
		{
			[]string{"a", "b", "c"},
			[]string{"d", "e", "f", "g"},
			[]string{"h"},
		},
		{
			[]int32{1, 2, 3},
			[]int32{4, 5, 6, 7},
			[]int32{8},
		},
	}
	na := []string{"x1", "x2", "x3"}
	da := NewFromArrays(x, na)

	x = [][]interface{}{
		{
			[]float64{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]string{"a", "b", "c", "d", "e", "f", "g", "h"},
		},
		{
			[]int32{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	na = []string{"x1", "x2", "x3"}
	db := NewFromArrays(x, na)

	dc := Dechunk(da)

	if !EqualReport(dc, db, true) {
		t.Fail()
	}
}
