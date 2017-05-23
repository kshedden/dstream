package dstream

import (
	"fmt"
	"math"
	"testing"
)

func datam1() (Dstream, Dstream) {
	x1 := []interface{}{
		[]float64{0, 1, 1},
		[]float64{0, math.NaN(), 1, 0},
	}
	x2 := []interface{}{
		[]float64{1, 1, math.NaN()},
		[]float64{1, 1, 1, 1},
	}
	x3 := []interface{}{
		[]float64{4, 1, -1},
		[]float64{3, 5, -5, 3},
	}
	x4 := []interface{}{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f", "g"},
	}
	dat := [][]interface{}{x1, x2, x3, x4}
	na := []string{"x1", "x2", "x3", "x4"}
	da := NewFromArrays(dat, na)

	x1 = []interface{}{
		[]float64{0, 1},
		[]float64{0, 1, 0},
	}
	x2 = []interface{}{
		[]float64{1, 1},
		[]float64{1, 1, 1},
	}
	x3 = []interface{}{
		[]float64{4, 1},
		[]float64{3, -5, 3},
	}
	x4 = []interface{}{
		[]string{"a", "b"},
		[]string{"d", "f", "g"},
	}
	dat = [][]interface{}{x1, x2, x3, x4}
	na = []string{"x1", "x2", "x3", "x4"}
	dm := NewFromArrays(dat, na)

	return da, dm
}

// Check that we can get variables by name and by position.
func checkPosName(da Dstream) (bool, string) {

	da.Reset()
	da.Next()

	// Make sure we can get variables by name and by position
	for k, na := range da.Names() {

		a := da.GetPos(k)
		b := da.Get(na)

		// TODO other types
		switch u := a.(type) {

		case []float64:
			v := b.([]float64)
			if len(u) != len(v) {
				msg := fmt.Sprintf("Variable '%s', position %d:\n", na, k)
				msg += fmt.Sprintf("Unequal lengths: %d != %d\n", len(u), len(v))
				return false, msg
			}
			for i, x := range u {
				if math.IsNaN(x) && math.IsNaN(v[i]) {
					continue
				}
				if x != v[i] {
					msg := fmt.Sprintf("Variable '%s', position %d:\n", na, k)
					return false, msg
				}
			}
		case []string:
			v := b.([]string)
			if len(u) != len(v) {
				msg := fmt.Sprintf("Variable '%s', position %d:\n", na, k)
				msg += fmt.Sprintf("Unequal lengths: %d != %d\n", len(u), len(v))
				return false, msg
			}
			for i, x := range u {
				if x != v[i] {
					return false, ""
				}
			}
		}
	}

	return true, ""
}

func TestDropNA1(t *testing.T) {

	da, de := datam1()
	dm := DropNA(da)

	dx := MemCopy(dm)

	if !EqualReport(de, dm, true) {
		t.Fail()
	}

	if !EqualReport(dm, dx, true) {
		t.Fail()
	}

	f, msg := checkPosName(da)
	if !f {
		print(msg)
		t.Fail()
	}
}
