package dstream

import (
	"bytes"
	"fmt"
	"os"
	"strings"
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

	da := FromCSV(rdr).SetStringVars([]string{"Country"}).SetFloatVars([]string{"Id", "Age"}).SetChunkSize(3).HasHeader().Done()

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

func TestStreamCSV2(t *testing.T) {

	x1 := []interface{}{
		[]float64{1, 2, 3},
		[]float64{4, 5},
	}
	x2 := []interface{}{
		[]float64{34, 19, 45},
		[]float64{46, 44},
	}
	dat := [][]interface{}{x1, x2}
	na := []string{"Id", "Age"}
	ex := NewFromArrays(dat, na)

	var bbuf bytes.Buffer
	bbuf.Write([]byte("Id,Age\n"))
	bbuf.Write([]byte("1,34\n"))
	bbuf.Write([]byte("2,19\n"))
	bbuf.Write([]byte("3,45\n"))
	bbuf.Write([]byte("4,46\n"))
	bbuf.Write([]byte("5,44\n"))

	rdr := bytes.NewReader(bbuf.Bytes())

	da := FromCSV(rdr).AllFloat().SetChunkSize(3).HasHeader().Done()

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
		msg := fmt.Sprintf("Incorrect number of observations (found %d, expected 3)\n", da.NumObs())
		os.Stderr.WriteString(msg)
		t.Fail()
	}
	if da.NumVar() != 2 {
		msg := fmt.Sprintf("Incorrect number of variables (found %d, expected 3)\n", da.NumVar())
		os.Stderr.WriteString(msg)
		t.Fail()
	}

	// Make sure we can get variables by name and by position.
	f, msg := checkPosName(da)
	if !f {
		print(msg)
		t.Fail()
	}
}

func TestCSVWriter1(t *testing.T) {
	data1 := `id,v1,v2,v3
1,2,3,4
1,3,4,5
2,4,5,6
3,5,6,7
3,99,99,99
3,100,101,102
4,200,201,202
`

	r := strings.NewReader(data1)
	ds := FromCSV(r).SetFloatVars([]string{"id", "v1", "v2", "v3"}).SetChunkSize(2).HasHeader().Done()

	var buf bytes.Buffer
	fm := map[string]string{"v1": "%.1f"}
	err := ToCSV(ds).SetWriter(&buf).FloatFmt("%.0f").Formats(fm).Done()
	if err != nil {
		panic(err)
	}

	es := `id,v1,v2,v3
1,2.0,3,4
1,3.0,4,5
2,4.0,5,6
3,5.0,6,7
3,99.0,99,99
3,100.0,101,102
4,200.0,201,202
`

	if es != string(buf.Bytes()) {
		t.Fail()
	}
}
