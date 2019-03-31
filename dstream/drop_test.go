package dstream

import "testing"

func TestDrop1(t *testing.T) {

	x := [][]interface{}{
		{
			[]float64{0, 1, 1},
			[]float64{0, 0, 1, 0},
		},
		{
			[]float64{1, 1, 1},
			[]float64{1, 1, 1, 1},
		},
		{
			[]float64{4, 1, -1},
			[]float64{3, 5, -5, 3},
		},
		{
			[]float64{1, 1, 1},
			[]float64{2, 2, 1, 1},
		},
		{
			[]float64{1, 1, 1},
			[]float64{2, 2, 2, 3},
		},
	}
	na := []string{"x1", "x2", "x3", "x4", "x5"}
	da := NewFromArrays(x, na)

	x = [][]interface{}{x[0], x[2], x[4]}
	na = []string{"x1", "x3", "x5"}
	de := NewFromArrays(x, na)

	db := DropCols(da, "x2", "x4")

	if !Equal(db, de) {
		t.Fail()
	}

	_, ok := db.(Dstream)
	if !ok {
		t.Fail()
	}
}
