package dstream

import (
	"testing"

	"github.com/gonum/floats"
)

func TestGetCol1(t *testing.T) {

	da, _ := datal1()

	z := []float64{0, 0, 1, 2, 3, 4,
		1, 2, 2, 2, 2, 1,
		3, 3, 3, 3, 3, 3}

	x2 := GetCol(da, "x2").([]float64)

	if !floats.Equal(z, x2) {
		t.Fail()
	}
}
