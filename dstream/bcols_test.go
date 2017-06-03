package dstream

import (
	"math"
	"testing"
)

func TestBcols1(t *testing.T) {

	da := NewBCols("testdata/bcols1", 5, nil, nil)

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
