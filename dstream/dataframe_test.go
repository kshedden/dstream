package dstream

import "testing"

func TestDataFrameLoadSave(t *testing.T) {

	x := [][]interface{}{
		{
			[]float64{1, 2, 3, 4, 5},
			[]float64{2, 9, 6},
		},
		{
			[]uint64{1, 0, 2, 5, 4},
			[]uint64{7, 4, 9},
		},
		{
			[]string{"a", "b", "c", "d", "e"},
			[]string{"f", "g", "h"},
		},
	}

	names := []string{"x", "y", "z"}
	data := NewFromArrays(x, names)

	data.(*DataFrame).Save("tmp.gz")

	da := &DataFrame{}
	da.Load("tmp.gz")

	if !EqualReport(data, da, true) {
		t.Fail()
	}
}
