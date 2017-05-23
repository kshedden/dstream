package dstream

import "fmt"

// FocusedReg embeds a Reg, and restricts it to the variable in column
// Col.  The other variables are linearly combined with the provided
// parameters (Params), and treated as an offset (added to the
// original offset if present).  The array Params should have length
// equal to NumCov() in the base DataProvider, and Params[Col] is not
// used.
type FocusedReg struct {
	Reg
	Col    int
	Params []float64
}

func (da *FocusedReg) NumCov() int {
	return 1
}

func (da *FocusedReg) XData(j int) []float64 {
	if j != 0 {
		msg := fmt.Sprintf("Invalid column (%d) in FocusedDataProvider", j)
		panic(msg)
	}
	return da.Reg.XData(da.Col)
}

func (da *FocusedReg) Offset() []float64 {

	n := len(da.YData())
	off := make([]float64, n)

	ofx := da.Reg.Offset()
	if ofx != nil {
		copy(off, ofx)
	}

	for j := 0; j < da.Reg.NumCov(); j++ {
		if j != da.Col {
			for i, x := range da.Reg.XData(j) {
				off[i] += da.Params[j] * x
			}
		}
	}

	return off
}
