package dstream

import (
	"fmt"
	"testing"
)

func TestAddcol1(t *testing.T) {

	x1 := []interface{}{
		[]float64{0, 0, 0},
		[]float64{1, 1},
		[]float64{2, 2, 2, 2},
	}
	x2 := []interface{}{
		[]string{"a", "a", "a"},
		[]string{"b", "b"},
		[]string{"c", "c", "c", "c"},
	}
	dat := [][]interface{}{x1, x2}
	na := []string{"x1", "x2"}
	da := NewFromArrays(dat, na)

	x3 := []interface{}{
		[]float64{1, 2, 3},
		[]float64{4, 5},
		[]float64{6, 7, 8, 9},
	}
	dat = [][]interface{}{x1, x2, x3}
	na = []string{"x1", "x2", "x3"}
	db := NewFromArrays(dat, na)

	z3 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	dm := Addcol(da, z3, "x3")

	dm.Reset()
	for dm.Next() {
		x1 := dm.GetPos(0).([]float64)
		x2 := dm.GetPos(1).([]string)
		x3 := dm.GetPos(2).([]float64)
		fmt.Printf("%v\n", x1)
		fmt.Printf("%v\n", x2)
		fmt.Printf("%v\n\n", x3)
	}

	if !EqualReport(dm, db, true) {
		t.Fail()
	}
}
