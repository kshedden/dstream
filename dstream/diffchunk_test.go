package dstream

import "testing"

func datad1() (Dstream, Dstream) {
	x1 := []interface{}{
		[]float64{0, 0, 0, 0, 0, 0},
		[]float64{1, 1, 1, 2, 3, 4},
		[]float64{2, 4, 6, 8, 10, 12},
	}
	x2 := []interface{}{
		[]float64{0, 0, 1, 2, 3, 4},
		[]float64{1, 2, 2, 2, 2, 1},
		[]float64{3, 3, 3, 3, 3, 3},
	}
	x3 := []interface{}{
		[]float64{1, 2, 3, 4, 5, 6},
		[]float64{4, 5, 6, 7, 8, 9},
		[]float64{7, 8, 9, 10, 11, 12},
	}
	x4 := []interface{}{
		[]string{"a", "b", "c", "d", "e", "f"},
		[]string{"d", "e", "f", "g", "h", "i"},
		[]string{"g", "h", "i", "j", "k", "l"},
	}
	dat := [][]interface{}{x1, x2, x3, x4}
	na := []string{"x1", "x2", "x3", "x4"}
	da := NewFromArrays(dat, na)

	x1 = []interface{}{
		[]float64{0, 0, 0, 0},
		[]float64{1, 2, 3, 4},
		[]float64{6, 8, 10, 12},
	}
	x1_1 := []interface{}{
		[]float64{0, 0, 0, 0},
		[]float64{0, 1, 1, 1},
		[]float64{2, 2, 2, 2},
	}
	x2 = []interface{}{
		[]float64{1, 2, 3, 4},
		[]float64{2, 2, 2, 1},
		[]float64{3, 3, 3, 3},
	}
	x2_2 := []interface{}{
		[]float64{1, 0, 0, 0},
		[]float64{-1, 0, 0, -1},
		[]float64{0, 0, 0, 0},
	}
	x3 = []interface{}{
		[]float64{3, 4, 5, 6},
		[]float64{6, 7, 8, 9},
		[]float64{9, 10, 11, 12},
	}
	x4 = []interface{}{
		[]string{"c", "d", "e", "f"},
		[]string{"f", "g", "h", "i"},
		[]string{"i", "j", "k", "l"},
	}
	dat = [][]interface{}{x1, x1_1, x2, x2_2, x3, x4}
	na = []string{"x1", "x1$d1", "x2", "x2$d2", "x3", "x4"}
	db := NewFromArrays(dat, na)

	return da, db
}

func TestDiff1(t *testing.T) {

	da, db := datad1()
	dx := DiffChunk(da, map[string]int{"x1": 1, "x2": 2})

	if !EqualReport(db, dx, true) {
		t.Fail()
	}
}
