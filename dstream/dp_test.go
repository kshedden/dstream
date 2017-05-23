package dstream

import (
	"math"
	"testing"
)

func datar1() Dstream {
	x1 := []interface{}{
		[]float64{0, 1, 1},
		[]float64{0, 0, 1, 0},
	}
	x2 := []interface{}{
		[]float64{1, 1, 1},
		[]float64{1, 1, 1, 1},
	}
	x3 := []interface{}{
		[]float64{4, 1, -1},
		[]float64{3, 5, -5, 3},
	}
	x4 := []interface{}{
		[]float64{1, 1, 1},
		[]float64{2, 2, 1, 1},
	}
	x5 := []interface{}{
		[]float64{1, 1, 1},
		[]float64{2, 2, 2, 3},
	}
	dat := [][]interface{}{x1, x2, x3, x4, x5}
	na := []string{"x1", "x2", "x3", "x4", "x5"}
	return NewFromArrays(dat, na)
}

func datap1() Reg {
	y := []interface{}{
		[]float64{0, 1, 1},
		[]float64{0, 0, 1, 0},
	}
	x1 := []interface{}{
		[]float64{1, 1, 1},
		[]float64{1, 1, 1, 1},
	}
	x2 := []interface{}{
		[]float64{4, 1, -1},
		[]float64{3, 5, -5, 3},
	}
	wgt := []interface{}{
		[]float64{1, 1, 1},
		[]float64{2, 2, 1, 1},
	}
	offset := []interface{}{
		[]float64{1, 1, 1},
		[]float64{2, 2, 2, 3},
	}
	dx := [][]interface{}{y, x1, x2, wgt, offset}
	na := []string{"y", "x1", "x2", "wgt", "offset"}
	da := NewFromArrays(dx, na)
	return NewReg(da, "y", []string{"x1", "x2"}, "wgt", "offset")
}

func datap2() Reg {
	y := []interface{}{
		[]float64{0, 1, 1},
		[]float64{0, 0, 1, 0},
	}
	x1 := []interface{}{
		[]float64{1, 1, 1},
		[]float64{1, 1, 1, 1},
	}
	x2 := []interface{}{
		[]float64{4, 1, -1},
		[]float64{3, 5, -5, 3},
	}
	x3 := []interface{}{
		[]float64{0, 2, -2},
		[]float64{1, 3, -2, 2},
	}
	wgt := []interface{}{
		[]float64{1, 1, 1},
		[]float64{2, 2, 1, 1},
	}
	offset := []interface{}{
		[]float64{1, 1, 1},
		[]float64{2, 2, 2, 3},
	}
	dx := [][]interface{}{y, x1, x2, x3, wgt, offset}
	na := []string{"y", "x1", "x2", "x3", "wgt", "offset"}
	da := NewFromArrays(dx, na)
	return NewReg(da, "y", []string{"x1", "x2", "x3"}, "wgt", "offset")
}

func datap3() Reg {
	y := []interface{}{
		[]float64{0, math.NaN(), 1},
		[]float64{0, 0, 1, 0},
	}
	x1 := []interface{}{
		[]float64{1, 1, 1},
		[]float64{1, 1, 1, 1},
	}
	x2 := []interface{}{
		[]float64{4, 1, -1},
		[]float64{3, 5, -5, 3},
	}
	x3 := []interface{}{
		[]float64{0, 2, -2},
		[]float64{1, 3, -2, math.NaN()},
	}
	wgt := []interface{}{
		[]float64{1, 1, 1},
		[]float64{2, 2, 1, 1},
	}
	offset := []interface{}{
		[]float64{1, 1, 1},
		[]float64{2, 2, 2, 3},
	}
	dx := [][]interface{}{y, x1, x2, x3, wgt, offset}
	na := []string{"y", "x1", "x2", "x3", "wgt", "offset"}
	da := NewFromArrays(dx, na)
	return NewReg(da, "y", []string{"x1", "x2", "x3"}, "wgt", "offset")
}

func TestDProv(t *testing.T) {

	dp := datap1()

	ye := [][]float64{
		[]float64{0, 1, 1},
		[]float64{0, 0, 1, 0},
	}
	x2e := [][]float64{
		[]float64{4, 1, -1},
		[]float64{3, 5, -5, 3},
	}
	wgte := [][]float64{
		[]float64{1, 1, 1},
		[]float64{2, 2, 1, 1},
	}

	for k := 0; k < 2; k++ {
		for i := 0; dp.Next(); i++ {
			if !aequalfloat64(dp.YData(), ye[i]) {
				t.Fail()
			}
			if !aequalfloat64(dp.XData(1), x2e[i]) {
				t.Fail()
			}
			if !aequalfloat64(dp.Weights(), wgte[i]) {
				t.Fail()
			}
		}
		dp.Reset()
	}
}

func TestFocus(t *testing.T) {

	ye := [][]float64{
		[]float64{0, 1, 1},
		[]float64{0, 0, 1, 0},
	}
	xe := [][]float64{
		[]float64{4, 1, -1},
		[]float64{3, 5, -5, 3},
	}
	ofe := [][]float64{
		[]float64{2, 8, -4},
		[]float64{6, 12, -3, 10},
	}
	wgte := [][]float64{
		[]float64{1, 1, 1},
		[]float64{2, 2, 1, 1},
	}

	dp := datap2()
	fp := FocusedReg{
		Reg:    dp,
		Col:    1,
		Params: []float64{1, 2, 3},
	}

	for k := 0; k < 2; k++ {
		for i := 0; dp.Next(); i++ {
			if !aequalfloat64(fp.YData(), ye[i]) {
				t.Fail()
			}
			if !aequalfloat64(fp.XData(0), xe[i]) {
				t.Fail()
			}
			if !aequalfloat64(fp.Offset(), ofe[i]) {
				t.Fail()
			}
			if !aequalfloat64(fp.Weights(), wgte[i]) {
				t.Fail()
			}
		}
		fp.Reset()
	}
}
