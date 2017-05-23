package dstream

import "testing"

// Test long chunks being broken into smaller chunks
func TestSizeChunk1(t *testing.T) {

	x1 := []interface{}{
		[]float64{0, 0, 0, 0, 0, 1, 1, 2, 3},
	}
	x2 := []interface{}{
		[]float64{1, 1, 1, 1, 2, 2, 4, 5, 6},
	}
	x3 := []interface{}{
		[]float64{2, 2, 3, 3, 3, 3, 7, 8, 9},
	}
	x4 := []interface{}{
		[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
	}
	dat := [][]interface{}{x1, x2, x3, x4}
	na := []string{"x1", "x2", "x3", "x4"}
	da := NewFromArrays(dat, na)

	x1 = []interface{}{
		[]float64{0, 0, 0, 0},
		[]float64{0, 1, 1, 2},
		[]float64{3},
	}
	x2 = []interface{}{
		[]float64{1, 1, 1, 1},
		[]float64{2, 2, 4, 5},
		[]float64{6},
	}
	x3 = []interface{}{
		[]float64{2, 2, 3, 3},
		[]float64{3, 3, 7, 8},
		[]float64{9},
	}
	x4 = []interface{}{
		[]string{"a", "b", "c", "d"},
		[]string{"e", "f", "g", "h"},
		[]string{"i"},
	}
	dat = [][]interface{}{x1, x2, x3, x4}
	na = []string{"x1", "x2", "x3", "x4"}
	de := NewFromArrays(dat, na)

	dx := SizeChunk(da, 4)

	if !EqualReport(dx, de, true) {
		t.Fail()
	}
}

// Test short chunks being combined into longer chunks
func TestSizeChunk2(t *testing.T) {

	x1 := []interface{}{
		[]float64{0, 0},
		[]float64{0, 0},
		[]float64{0, 1},
		[]float64{1, 2, 3},
	}
	x2 := []interface{}{
		[]float64{1, 1},
		[]float64{1, 1},
		[]float64{2, 2},
		[]float64{4, 5, 6},
	}
	x3 := []interface{}{
		[]float64{2, 2},
		[]float64{3, 3},
		[]float64{3, 3},
		[]float64{7, 8, 9},
	}
	x4 := []interface{}{
		[]string{"a", "b"},
		[]string{"c", "d"},
		[]string{"e", "f"},
		[]string{"g", "h", "i"},
	}
	dat := [][]interface{}{x1, x2, x3, x4}
	na := []string{"x1", "x2", "x3", "x4"}
	da := NewFromArrays(dat, na)

	x1 = []interface{}{
		[]float64{0, 0, 0, 0},
		[]float64{0, 1, 1, 2},
		[]float64{3},
	}
	x2 = []interface{}{
		[]float64{1, 1, 1, 1},
		[]float64{2, 2, 4, 5},
		[]float64{6},
	}
	x3 = []interface{}{
		[]float64{2, 2, 3, 3},
		[]float64{3, 3, 7, 8},
		[]float64{9},
	}
	x4 = []interface{}{
		[]string{"a", "b", "c", "d"},
		[]string{"e", "f", "g", "h"},
		[]string{"i"},
	}
	dat = [][]interface{}{x1, x2, x3, x4}
	na = []string{"x1", "x2", "x3", "x4"}
	de := NewFromArrays(dat, na)

	dx := SizeChunk(da, 4)

	if !EqualReport(dx, de, true) {
		t.Fail()
	}
}
