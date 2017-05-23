package dstream

// Reg is an interface for accessing a dataset for regression analysis
// by chunks.
type Reg interface {
	Dstream

	// YData returns the dependent variable data for the current
	// chunk.
	YData() []float64

	// XData(j) returns the dependent variable data for variable j
	// in the current chunk.
	XData(int) []float64

	// Weights returns the frequency weights for the current
	// chunk.  If there are no weights, a nil slice is returned.
	Weights() []float64

	// Offset returns offsets for the current chunk.  If there is
	// no offset, a nil slice is returned.
	Offset() []float64

	// NumCov returns the number of covariates in the dataset.
	NumCov() int

	// XNames returns the covariate names
	XNames() []string
}

// regArrays implements the Reg interface.
type regArrays struct {
	Dstream

	// Positions within data array
	y   int
	x   []int
	wgt int
	off int
}

func (dp *regArrays) XNames() []string {
	names := dp.Dstream.Names()
	xnames := make([]string, len(dp.x))
	for j, i := range dp.x {
		xnames[j] = names[i]
	}
	return xnames
}

func (dp *regArrays) NumCov() int {
	return len(dp.x)
}

func (dp *regArrays) YData() []float64 {
	return dp.Dstream.GetPos(dp.y).([]float64)
}

func (dp *regArrays) Offset() []float64 {
	if dp.off == -1 {
		return nil
	}
	return dp.Dstream.GetPos(dp.off).([]float64)
}

func (dp *regArrays) Weights() []float64 {
	if dp.wgt == -1 {
		return nil
	}
	return dp.Dstream.GetPos(dp.wgt).([]float64)
}

func (dp *regArrays) XData(j int) []float64 {
	return dp.Dstream.GetPos(dp.x[j]).([]float64)
}

func find(x string, y []string) int {
	for i, v := range y {
		if x == v {
			return i
		}
	}
	return -1
}

// NewReg creates a Reg from the given Dstream.  The arguments y, wgt,
// offand x provide the names of the variables that will become the
// corresponding components of Reg.  The offset and weight may be
// passed as zero-length strings if not present.  If x is passed as
// nil, all variables other than y, wgt, and off are included in x.
func NewReg(data Dstream, y string, x []string, wgt, off string) Reg {

	names := data.Names()

	if x == nil {
		for _, na := range names {
			if na != y && na != wgt && na != off {
				x = append(x, na)
			}
		}
	}

	yi := find(y, names)
	wgti := find(wgt, names)
	offi := find(off, names)
	xi := make([]int, len(x))
	for j, v := range x {
		xi[j] = find(v, names)
	}

	return &regArrays{
		Dstream: data,
		y:       yi,
		x:       xi,
		wgt:     wgti,
		off:     offi,
	}
}
