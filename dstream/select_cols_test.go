package dstream

import (
	"testing"
)

func TestSelect1(t *testing.T) {

	da, _ := dataf1()
	db, _ := dataf1()

	da = SelectCols(da, "x2", "x3")
	db = DropCols(db, "x1", "x4")

	if !EqualReport(da, db, true) {
		t.Fail()
	}
}
