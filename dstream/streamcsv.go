package dstream

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// csvReader supports reading a Dstream from an io.Reader.
type csvReader struct {
	rdr io.Reader
	cvr *csv.Reader

	bdata []interface{}

	// Used to hold the first row of data if we needed to read it
	// to get the number of columns.
	stashrec []string

	chunkSize int
	nvar      int
	nobs      int
	hasheader bool
	doneinit  bool
	done      bool

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
}

// FromCSV returns a Dstream that reads from a CSV source.  Call at
// least one SetXX method to define variables to be retrieved.  For
// further configuration, chain calls to other SetXXX methods, and
// finally call Done to produce the Dstream.
func FromCSV(r io.Reader) *csvReader {

	dr := &csvReader{
		rdr: r,
	}

	return dr
}

// Done is called when all configuration is complete to obtain a Dstream.
func (cs *csvReader) Done() Dstream {
	return cs
}

// Close does nothing, the caller should explicitly close the
// io.Reader passed to FromCSV if needed.
func (cs *csvReader) Close() {
}

// HasHeader indicates that the first row of the data file contains
// column names.  The default behavior is that there is no header.
func (cs *csvReader) HasHeader() *csvReader {
	if cs.doneinit {
		msg := "FromCSV: can't call HasHeader after beginning data read"
		panic(msg)
	}
	cs.hasheader = true
	return cs
}

func (cs *csvReader) init() {

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
	if cs.hasheader {
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

	cs.doneinit = true
}

// SetChunkSize sets the size of chunks for this Dstream, it can only
// be called before reading begins.
func (cs *csvReader) SetChunkSize(c int) *csvReader {
	cs.chunkSize = c
	return cs
}

// SetFloatVars sets the names of the variables to be converted to
// float64 type.  Refer to the columns by V1, V2, etc. if there is no
// header row.
func (cs *csvReader) SetFloatVars(x []string) *csvReader {
	cs.floatVars = x
	return cs
}

// SetStringVars sets the names of the variables to be stored as
// string type values.  Refer to the columns by V1, V2, etc. if there
// is no header row.
func (cs *csvReader) SetStringVars(x []string) *csvReader {
	cs.stringVars = x
	return cs
}

// Names returns the names of the variables in the dstream.
func (cs *csvReader) Names() []string {
	if !cs.doneinit {
		cs.init()
	}
	return cs.names
}

// NumVar returns the number of variables in the dstream.
func (cs *csvReader) NumVar() int {
	if !cs.doneinit {
		cs.init()
	}
	return cs.nvar
}

// NumObs returns the number of observations in the dstream.  If the
// dstream has not been fully read, returns -1.
func (cs *csvReader) NumObs() int {
	if cs.done {
		return cs.nobs
	}
	return -1
}

// GetPos returns a chunk of a data column by column position.
func (cs *csvReader) GetPos(j int) interface{} {
	if !cs.doneinit {
		cs.init()
	}
	return cs.bdata[j]
}

// Get returns a chunk of a data column by name.
func (cs *csvReader) Get(na string) interface{} {
	if !cs.doneinit {
		cs.init()
	}
	pos, ok := cs.namepos[na]
	if !ok {
		msg := fmt.Sprintf("Variable '%s' not found", na)
		panic(msg)
	}
	return cs.bdata[pos]
}

// Reset attempts to reset the Dstream that is reading from an
// io.Reader.  This is only possible if the underlying reader is
// seekable, so reset panics if the seek cannot be performed.
func (cs *csvReader) Reset() {
	r, ok := cs.rdr.(io.ReadSeeker)
	if !ok {
		panic("cannot reset")
	}
	_, err := r.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}
	cs.done = false
	cs.doneinit = false
	cs.nobs = 0
}

// Next advances to the next chunk.
func (cs *csvReader) Next() bool {

	if !cs.doneinit {
		cs.init()
	}
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

// csvWriter supports writing a Dstream to an io.Writer in csv format.
type csvWriter struct {

	// The Dstream to be written.
	stream Dstream

	// Format for float type value
	floatFmt string

	// A slice of format types, stored per-variable.
	fmts []string

	wtr io.Writer
}

// ToCSV supports writing a DStream in CSV format.  Call SetWriter
// or Filename on the returned value to configure the underlying
// writer, then call additional methods for customization as desired,
// and finally call Done to complete the writing.
func ToCSV(d Dstream) *csvWriter {
	c := &csvWriter{
		stream: d,
	}
	return c
}

// FloatFmt sets the format string to be used when writing float
// values.  This value is ignored for columns specified in a call to
// the Formats method.
func (dw *csvWriter) FloatFmt(fmt string) *csvWriter {

	dw.floatFmt = fmt
	return dw
}

// Formats sets format strings to be used when writing the Dstream.
// The provided argument is a map from variable names to variable
// formats.
func (dw *csvWriter) Formats(fmts map[string]string) *csvWriter {

	vp := VarPos(dw.stream)

	if dw.fmts == nil {
		nvar := dw.stream.NumVar()
		dw.fmts = make([]string, nvar)
	}
	for v, f := range fmts {
		pos, ok := vp[v]
		if !ok {
			msg := fmt.Sprintf("ToCSV: column %s not found", v)
			panic(msg)
		}
		dw.fmts[pos] = f
	}

	return dw
}

// Filename configures the CSVWriter to write to the given named file.
func (dw *csvWriter) Filename(name string) *csvWriter {

	var err error
	dw.wtr, err = os.Create(name)
	if err != nil {
		panic(err)
	}

	return dw
}

// SetWriter configures the CSVWriter to write to the given io stream.
func (dw *csvWriter) SetWriter(w io.Writer) *csvWriter {

	dw.wtr = w
	return dw
}

// getFmt is a utility for getting the format string for a given
// column.
func (dw *csvWriter) getFmt(t string, col int) string {

	if dw.fmts != nil && dw.fmts[col] != "" {
		return dw.fmts[col]
	}

	switch t {
	case "float":
		if dw.floatFmt == "" {
			return "%.8f"
		} else {
			return dw.floatFmt
		}
	case "int":
		return "%d"
	default:
		panic("unkown type")
	}
}

// Done completes writing a Dstream to a specified io.Writer in csv
// format.
func (dw *csvWriter) Done() error {

	if dw.wtr == nil {
		return errors.New("ToCSV: writer must be set before calling Done")
	}

	csw := csv.NewWriter(dw.wtr)

	err := csw.Write(dw.stream.Names())
	if err != nil {
		return err
	}

	nvar := dw.stream.NumVar()
	rec := make([]string, nvar)
	fmts := make([]string, nvar)

	firstrow := true
	for dw.stream.Next() {
		n := ilen(dw.stream.GetPos(0))

		for i := 0; i < n; i++ {
			for j := 0; j < nvar; j++ {
				// TODO: better support for types
				switch x := dw.stream.GetPos(j).(type) {
				case []float64:
					if firstrow {
						fmts[j] = dw.getFmt("float", j)
					}
					rec[j] = fmt.Sprintf(fmts[j], x[i])
				case []string:
					rec[j] = x[i]
				default:
					rec[j] = ""
				}
			}
			if err := csw.Write(rec); err != nil {
				return err
			}
			firstrow = false
		}
	}

	csw.Flush()

	return nil
}
