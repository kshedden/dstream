package dstream

import (
	"math"
	"testing"
)

func describeData() Dstream {

	x1 := []interface{}{
		[]float64{0, 0, 0},
		[]float64{1, 1, 1},
		[]float64{2, 2, 2},
	}
	x2 := []interface{}{
		[]float64{1, 3, 5},
		[]float64{1, 3, 5},
		[]float64{1, 3, 5},
	}
	x3 := []interface{}{
		[]float64{5, 5, 5},
		[]float64{5, math.NaN(), 5},
		[]float64{5, math.Inf(-1), math.Inf(1)},
	}

	dat := [][]interface{}{x1, x2, x3}
	na := []string{"x1", "x2", "x3"}

	return NewFromArrays(dat, na)
}

func compareStats(a, b Stats) bool {

	if math.Abs(a.Mean-b.Mean) > 1e-6 {
		return false
	}

	if math.Abs(a.Min-b.Min) > 1e-6 {
		return false
	}

	if math.Abs(a.Max-b.Max) > 1e-6 {
		return false
	}

	if math.Abs(a.SD-b.SD) > 1e-6 {
		return false
	}

	return true
}

func TestDescribe1(t *testing.T) {

	da := describeData()
	st := Describe(da)

	// Correct values
	e := map[string]Stats{
		"x1": Stats{
			Min:  0,
			Max:  2,
			Mean: 1,
			SD:   math.Sqrt(6.0 / 9.0),
			N:    9,
			NaN:  0,
			Inf:  0,
		},
		"x2": Stats{
			Min:  1,
			Max:  5,
			Mean: 3,
			SD:   math.Sqrt(24.0 / 9.0),
			N:    9,
			NaN:  0,
			Inf:  0,
		},
		"x3": Stats{
			Min:  5,
			Max:  5,
			Mean: 5,
			SD:   0,
			N:    6,
			Inf:  2,
			NaN:  1,
		},
	}

	for k, v := range e {
		if !compareStats(v, st[k]) {
			t.Fail()
		}
	}
}
