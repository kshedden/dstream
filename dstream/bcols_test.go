package dstream

import (
	"math"
	"testing"
)

func TestBcols1(t *testing.T) {

	da := NewBCols("testdata/bcols1", 5).Done()

	for rep := 0; rep < 2; rep++ {
		var m, ii int
		for da.Next() {
			for j := 0; j < da.NumVar(); j++ {

				// Get the column by position and by name
				a := da.GetPos(j)
				b := da.Get(da.Names()[j])

				for _, v := range []interface{}{a, b} {
					switch v := v.(type) {
					case []uint8:
						m = len(v)
						for i, x := range v {
							if x != uint8((ii+i)*(ii+i)) {
								t.Fail()
							}
						}
					case []uint16:
						m = len(v)
						for i, x := range v {
							if x != uint16((ii+i)*(ii+i)) {
								t.Fail()
							}
						}
					case []uint32:
						m = len(v)
						for i, x := range v {
							if x != uint32((ii+i)*(ii+i)) {
								t.Fail()
							}
						}
					case []uint64:
						m = len(v)
						for i, x := range v {
							if x != uint64((ii+i)*(ii+i)) {
								t.Fail()
							}
						}
					case []float64:
						m = len(v)
						for i, x := range v {
							d := x - float64((ii+i)*(ii+i))
							if math.Abs(d) > 1e-2 {
								t.Fail()
							}
						}
					}
				}
			}
			ii += m
		}
		da.Reset()
	}

	da.Close()
}

func TestToBCols(t *testing.T) {

	x1 := []interface{}{
		[]float64{0, 1, 2, 3},
		[]float64{4, 5, 6},
	}
	x2 := []interface{}{
		[]float32{1, 2, 3, 4},
		[]float32{5, 6, 7},
	}
	x3 := []interface{}{
		[]string{"a", "b", "c", "d"},
		[]string{"e", "f", "g"},
	}
	dat := [][]interface{}{x1, x2, x3}
	na := []string{"x1", "x2", "x3"}
	da := NewFromArrays(dat, na)

	ToBCols(da).Path("testdata/tobcols").Done()
	db := NewBCols("testdata/tobcols", 4).Done()
	if !EqualReport(da, db, true) {
		t.Fail()
	}

	dax := DropCols(da, []string{"x2"})
	ToBCols(dax).Path("testdata/tobcols").Done()
	db = NewBCols("testdata/tobcols", 4).Done()
	if !EqualReport(dax, db, true) {
		t.Fail()
	}

	dax = DropCols(da, []string{"x2"})
	ToBCols(dax).Path("testdata/tobcols").Done()
	db = NewBCols("testdata/tobcols", 4).Include([]string{"x1", "x3"}).Done()
	if !EqualReport(dax, db, true) {
		t.Fail()
	}
}
