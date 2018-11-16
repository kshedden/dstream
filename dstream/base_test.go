package dstream

import "testing"

func TestNewFromContig(t *testing.T) {

	chunked := [][]interface{}{
		{[]float64{1, 2, 3, 4, 5}},
		{[]float64{3, 2, 1, 0, 0}},
	}

	contig := []interface{}{
		[]float64{1, 2, 3, 4, 5},
		[]float64{3, 2, 1, 0, 0},
	}

	names := []string{"x", "y"}

	dchunked := NewFromArrays(chunked, names)
	dcontig := NewFromFlat(contig, names)

	if !EqualReport(dchunked, dcontig, true) {
		t.Fail()
	}
}
