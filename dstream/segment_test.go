package dstream

import (
	"os"
	"testing"
)

// All segments are within the original chunks.
func datas1() (Dstream, Dstream) {

	x := [][]interface{}{
		{
			[]float64{0, 0, 0},
			[]float64{1, 1, 1},
			[]float64{2, 2, 3},
		},
		{
			[]float64{0, 0, 1},
			[]float64{1, 2, 2},
			[]float64{3, 3, 3},
		},
		{
			[]float64{1, 2, 3},
			[]float64{4, 5, 6},
			[]float64{7, 8, 9},
		},
		{
			[]string{"a", "b", "c"},
			[]string{"d", "e", "f"},
			[]string{"g", "h", "i"},
		},
	}

	na := []string{"x1", "x2", "x3", "x4"}
	da := NewFromArrays(x, na)

	x = [][]interface{}{
		{
			[]float64{0, 0},
			[]float64{0},
			[]float64{1},
			[]float64{1, 1},
			[]float64{2, 2},
			[]float64{3},
		},
		{
			[]float64{0, 0},
			[]float64{1},
			[]float64{1},
			[]float64{2, 2},
			[]float64{3, 3},
			[]float64{3},
		},
		{
			[]float64{1, 2},
			[]float64{3},
			[]float64{4},
			[]float64{5, 6},
			[]float64{7, 8},
			[]float64{9},
		},
		{
			[]string{"a", "b"},
			[]string{"c"},
			[]string{"d"},
			[]string{"e", "f"},
			[]string{"g", "h"},
			[]string{"i"},
		},
	}
	na = []string{"x1", "x2", "x3", "x4"}
	dm := NewFromArrays(x, na)

	return da, dm
}

// Segments span the original chunks.
func datas2() (Dstream, Dstream) {

	x := [][]interface{}{
		{
			[]float64{0, 0, 1},
			[]float64{1, 2, 2},
			[]float64{2, 3, 3},
		},
		{
			[]float64{1, 1, 2},
			[]float64{2, 3, 3},
			[]float64{3, 4, 4},
		},
		{
			[]float64{1, 2, 3},
			[]float64{4, 5, 6},
			[]float64{7, 8, 9},
		},
		{
			[]string{"a", "b", "c"},
			[]string{"d", "e", "f"},
			[]string{"g", "h", "i"},
		},
	}
	na := []string{"x1", "x2", "x3", "x4"}
	da := NewFromArrays(x, na)

	x = [][]interface{}{
		{
			[]float64{0, 0},
			[]float64{1, 1},
			[]float64{2, 2, 2},
			[]float64{3, 3},
		},
		{
			[]float64{1, 1},
			[]float64{2, 2},
			[]float64{3, 3, 3},
			[]float64{4, 4},
		},
		{
			[]float64{1, 2},
			[]float64{3, 4},
			[]float64{5, 6, 7},
			[]float64{8, 9},
		},
		{
			[]string{"a", "b"},
			[]string{"c", "d"},
			[]string{"e", "f", "g"},
			[]string{"h", "i"},
		},
	}
	na = []string{"x1", "x2", "x3", "x4"}
	dm := NewFromArrays(x, na)

	return da, dm
}

func TestSegment1(t *testing.T) {

	da, dm := datas1()
	dx := Segment(da, "x1", "x2")

	if !EqualReport(dm, dx, true) {
		t.Fail()
	}
}

func TestSegment2(t *testing.T) {

	da, dm := datas2()
	dx := Segment(da, "x1", "x2")

	if !EqualReport(dm, dx, true) {
		t.Fail()
	}
}

func gensegdat(a, b, n int) Dstream {

	var chunks []interface{}
	var x1 []uint16

	for k := 0; k < n; k++ {

		x1 = append(x1, uint16(k/(b+1)))

		if k%a == 0 {
			chunks = append(chunks, x1)
			x1 = make([]uint16, 0, 0)
		}
	}

	na := []string{"x1"}
	return NewFromArrays([][]interface{}{chunks}, na)
}

func TestSegment3(t *testing.T) {

	for _, q := range []struct {
		a int
		b int
		n int
	}{
		{3, 2, 10},
		{2, 3, 10},
		{2, 2, 10},
		{10, 2, 10},
	} {
		da := gensegdat(q.a, q.b, q.n)
		dx := Segment(da, "x1")

		dx.Reset()
		for dx.Next() {
			x := dx.GetPos(0).([]uint16)
			for _, y := range x {
				if y != x[0] {
					t.Fail()
				}
			}
		}

		f, msg := checkPosName(da)
		if !f {
			os.Stderr.WriteString(msg)
			t.Fail()
		}
	}
}
