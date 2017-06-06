package dstream

import (
	"bytes"
	"testing"
)

func TestStreamCSV1(t *testing.T) {

	x1 := []interface{}{
		[]float64{1, 2, 3},
		[]float64{4, 5},
	}
	x2 := []interface{}{
		[]float64{34, 19, 45},
		[]float64{46, 44},
	}
	x3 := []interface{}{
		[]string{"Argentina", "Canada", "India"},
		[]string{"Mexico", "Egypt"},
	}
	dat := [][]interface{}{x1, x2, x3}
	na := []string{"Id", "Age", "Country"}
	ex := NewFromArrays(dat, na)

	var bbuf bytes.Buffer
	bbuf.Write([]byte("Id,Country,Pop,Age\n"))
	bbuf.Write([]byte("1,Argentina,29,34\n"))
	bbuf.Write([]byte("2,Canada,17,19\n"))
	bbuf.Write([]byte("3,India,234,45\n"))
	bbuf.Write([]byte("4,Mexico,94,46\n"))
	bbuf.Write([]byte("5,Egypt,89,44\n"))

	rdr := bytes.NewReader(bbuf.Bytes())

	da := FromCSV(rdr).SetStringVars([]string{"Country"}).SetFloatVars([]string{"Id", "Age"}).SetChunkSize(3).HasHeader()

	// Check first read
	if !EqualReport(ex, da, true) {
		t.Fail()
	}

	// Make sure reset/reread works
	da.Reset()
	if !EqualReport(ex, da, true) {
		t.Fail()
	}

	if da.NumObs() != 5 {
		t.Fail()
	}
	if da.NumVar() != 3 {
		t.Fail()
	}

	// Make sure we can get variables by name and by position.
	f, msg := checkPosName(da)
	if !f {
		print(msg)
		t.Fail()
	}
}
