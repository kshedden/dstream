package dstream

import "testing"

func TestVarTypes(t *testing.T) {

	x := [][]interface{}{
		{
			[]float64{0, 1, 2},
		},
		{
			[]float32{0, 1, 2},
		},
		{
			[]uint64{0, 1, 2},
		},
		{
			[]uint32{0, 1, 2},
		},
		{
			[]uint16{0, 1, 2},
		},
		{
			[]uint8{0, 1, 2},
		},
		{
			[]int64{0, 1, 2},
		},
		{
			[]int32{0, 1, 2},
		},
		{
			[]int16{0, 1, 2},
		},
		{
			[]int8{0, 1, 2},
		},
	}

	names := []string{"float64", "float32", "uint64", "uint32", "uint16", "uint8",
		"int64", "int32", "int16", "int8"}
	da := NewFromArrays(x, names)

	da.Reset()
	da.Next()
	vt := VarTypes(da)

	for k, v := range vt {
		if k != v {
			t.Fail()
		}
	}
}
