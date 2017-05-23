package dstream

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"strconv"
)

type CSVStreamer struct {
	rdr io.Reader
	cvr *csv.Reader

	bdata []interface{}

	// Used to hold the first row of data if we needed to read it
	// to get the number of columns.
	stashrec []string

	chunkSize int
	nvar      int
	nobs      int
	nobsKnown bool
	hasheader bool

	namepos map[string]int
	names   []string

	// Names of variables to be converted to floats
	floatVars []string

	// Names of variables to be stored as strings
	stringVars []string

	// Positions of variables to be converted to floats
	floatVarsPos []int

	// Positions of variables to be stored as strings
	stringVarsPos []int

	done bool
}

// FromCSV returns a Dstream that reads from a CSV data file.  Call at
// least one SetXX method to define variables to be retrieved, then
// call Init before using.
func FromCSV(r io.Reader) *CSVStreamer {

	da := &CSVStreamer{
		rdr: r,
	}

	return da
}

// Init must be called before using a Dstream that reads a CSV file.
// If hasheader is true, the first row is assumed to be a row of
// variable names.  Otherwise the data starts in the first row, and
// the variables are named V1, V2, ....
func (cs *CSVStreamer) Init(hasheader bool) {

	cs.hasheader = hasheader
	if cs.chunkSize == 0 {
		cs.chunkSize = 10000
	}

	cs.cvr = csv.NewReader(cs.rdr)

	var row1 []string
	var err error
	row1, err = cs.cvr.Read()
	if err != nil {
		panic(err)
	}

	hdrmap := make(map[string]int)
	if hasheader {
		for k, v := range row1 {
			hdrmap[v] = k
		}
	} else {
		cs.stashrec = row1
		for k := range row1 {
			hdrmap[fmt.Sprintf("V%d", k+1)] = k
		}
	}

	cs.floatVarsPos = cs.floatVarsPos[0:0]
	for _, v := range cs.floatVars {
		pos, ok := hdrmap[v]
		if !ok {
			msg := fmt.Sprintf("Variable '%s' not found", v)
			panic(msg)
		}
		cs.floatVarsPos = append(cs.floatVarsPos, pos)
	}

	cs.stringVarsPos = cs.stringVarsPos[0:0]
	for _, v := range cs.stringVars {
		pos, ok := hdrmap[v]
		if !ok {
			msg := fmt.Sprintf("Variable '%s' not found", v)
			panic(msg)
		}
		cs.stringVarsPos = append(cs.stringVarsPos, pos)
	}

	cs.nvar = len(cs.floatVars) + len(cs.stringVars)
	for _, _ = range cs.floatVars {
		cs.bdata = append(cs.bdata, make([]float64, 0, 1000))
	}
	for _, _ = range cs.stringVars {
		cs.bdata = append(cs.bdata, make([]string, 0, 1000))
	}

	cs.names = append(cs.floatVars, cs.stringVars...)

	cs.namepos = make(map[string]int)
	for k, na := range cs.names {
		cs.namepos[na] = k
	}
}

func (cs *CSVStreamer) SetChunkSize(c int) {
	cs.chunkSize = c
}

// SetFloatVars sets the names of the variables to be converted to
// float64 type.  Refer to the columns by V1, V2, etc. if there is no
// header row.
func (cs *CSVStreamer) SetFloatVars(x []string) {
	cs.floatVars = x
}

// SetStringVars sets the names of the variables to be stored as
// string type values.  Refer to the columns by V1, V2, etc. if there
// is no header row.
func (cs *CSVStreamer) SetStringVars(x []string) {
	cs.stringVars = x
}

func (cs *CSVStreamer) Names() []string {
	return cs.names
}

func (cs *CSVStreamer) NumVar() int {
	return cs.nvar
}

func (cs *CSVStreamer) NumObs() int {
	if cs.nobsKnown {
		return cs.nobs
	}
	return -1
}

func (cs *CSVStreamer) GetPos(j int) interface{} {
	return cs.bdata[j]
}

func (cs *CSVStreamer) Get(na string) interface{} {
	pos, ok := cs.namepos[na]
	if !ok {
		msg := fmt.Sprintf("Variable '%s' not found", na)
		panic(msg)
	}
	return cs.bdata[pos]
}

// Reset attempts to reset the CSVStreamer, but it may be possible
// depending on the whether the underlying reader is seekable.
func (cs *CSVStreamer) Reset() {
	r, ok := cs.rdr.(io.ReadSeeker)
	if !ok {
		panic("cannot reset")
	}
	_, err := r.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}
	cs.cvr = csv.NewReader(cs.rdr)
	cs.done = false
	cs.nobs = 0
	if cs.hasheader {
		cs.cvr.Read()
	}
}

func (cs *CSVStreamer) Next() bool {

	if cs.done {
		return false
	}

	truncate(cs.bdata)

	for j := 0; j < cs.chunkSize; j++ {

		var rec []string
		var err error
		if cs.stashrec != nil {
			rec = cs.stashrec
			cs.stashrec = nil
		} else {
			rec, err = cs.cvr.Read()
			if err == io.EOF {
				cs.done = true
				cs.nobsKnown = true
				return ilen(cs.bdata[0]) > 0
			} else if err != nil {
				panic(err)
			}
		}
		cs.nobs++

		for k, pos := range cs.floatVarsPos {
			x, err := strconv.ParseFloat(rec[pos], 64)
			if err != nil {
				x = math.NaN()
			}
			u := cs.bdata[k].([]float64)
			cs.bdata[k] = append(u, x)
		}

		m := len(cs.floatVarsPos)
		for k, pos := range cs.stringVarsPos {
			u := cs.bdata[m+k].([]string)
			cs.bdata[m+k] = append(u, rec[pos])
		}
	}

	return true
}

func ToCSV(data Dstream, wtr io.Writer) error {

	data.Reset()

	csw := csv.NewWriter(wtr)
	err := csw.Write(data.Names())
	if err != nil {
		return err
	}

	nvar := data.NumVar()
	rec := make([]string, nvar)
	for data.Next() {
		n := ilen(data.GetPos(0))

		for i := 0; i < n; i++ {
			for j := 0; j < nvar; j++ {
				switch x := data.GetPos(j).(type) {
				case []float64:
					rec[j] = fmt.Sprintf("%.8f", x[i])
				case []string:
					rec[j] = x[i]
				}
			}
			err = csw.Write(rec)
			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}
