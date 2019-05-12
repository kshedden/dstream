package dstream

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// CSVReader supports reading a Dstream from an io.Reader.
type CSVReader struct {

	// The reader for the file
	rdr io.Reader

	// The csv reader that parses the raw text
	csvrdr *csv.Reader

	// bdata holds the data
	bdata []interface{}

	// We need to read the first row to get the number of columns.
	// or the header.  We can then store it here so that it is
	// included in the first data chunk.
	firstrow []string

	// If true, skip records with unparseable CSV data, otherwise
	// panic on them.
	skipErrors bool

	// The record separator, passed to csv.Reader
	comma rune

	// The number of records to read at once
	chunkSize int

	// The number of variables
	nvar int

	// The number of observations, not known until reading is complete
	nobs int

	// If true, the first row of the data file is a header (contains column names)
	hasheader bool

	// doneInit becomes true after init is called
	doneinit bool

	// done becomes true after configuration is complete
	done bool

	// Index of current chunk
	chunknum int

	// If limitchunk > 0, read only this many chunks, otherwise read all chunks.
	limitchunk int

	// Map from variable names to column positions in the dstream
	namepos map[string]int

	// Map from variable names to column positions in the file
	filenamepos map[string]int

	// The positions in the file of the variables in the dstream, in dstream order
	filepos []int

	// The names of the dstream variables
	names []string

	// The variable name/type structs, provided by the caller
	types []VarType

	// The variable types
	dtypes []Dtype

	// A function for parsing time values
	parseTime func(string) time.Time
}

// VarType defines the name and type of one column of the CSV file.
type VarType struct {

	// The variable name
	Name string

	// The data type of the variable
	Type Dtype
}

// FromCSV returns a Dstream that reads from a CSV source.  Call at
// least one SetXX method to define variables to be retrieved.  For
// further configuration, chain calls to other SetXXX methods, and
// finally call Done to produce the Dstream.
func FromCSV(r io.Reader) *CSVReader {

	dr := &CSVReader{
		rdr: r,
	}

	return dr
}

// ParseTime sets the time parsing function./
func (cs *CSVReader) ParseTime(f func(string) time.Time) *CSVReader {
	cs.parseTime = f
	return cs
}

// Done is called when all configuration is complete.  After calling
// Done, the DStream can be used.
func (cs *CSVReader) Done() Dstream {
	cs.init()
	return cs
}

// SkipErrors results in lines with unpareable CSV content being
// skipped (the csv.ParseError is printed to stdio).
func (cs *CSVReader) SkipErrors() *CSVReader {
	cs.skipErrors = true
	return cs
}

// SetTypes species the types of the variables.  If the CSV file has a header,
// these values may appear in any order, and variables not included in types
// are omitted. If the CSV file does not have a header, then
// types must match the columns in the file, in the correct order.
func (cs *CSVReader) SetTypes(types []VarType) *CSVReader {
	cs.types = types
	return cs
}

// LimitChunk sets the number of chunks to read.
func (cs *CSVReader) LimitChunk(n int) *CSVReader {
	cs.limitchunk = n
	return cs
}

// Close does nothing and is implemented to satisfy the Dstream interface.
// If any io.Reader values passed to FromCSV need closing, they should be
// closed by the caller.
func (cs *CSVReader) Close() {
}

// HasHeader indicates that the first row of the data file contains
// column names.  The default behavior is that there is no header.
func (cs *CSVReader) HasHeader() *CSVReader {
	if cs.doneinit {
		msg := "FromCSV: can't call HasHeader after beginning data read"
		panic(msg)
	}
	cs.hasheader = true
	return cs
}

// Comma sets the delimiter (comma rune) for the CSVReader.  By default,
// the comma rune is a comma.
func (cs *CSVReader) Comma(c rune) *CSVReader {
	cs.comma = c
	return cs
}

// Consistency checks for arguments.
func (cs *CSVReader) checkArgs() {

}

func (cs *CSVReader) init() {

	cs.checkArgs()

	if cs.chunkSize == 0 {
		cs.chunkSize = 10000
	}

	cs.csvrdr = csv.NewReader(cs.rdr)
	if cs.comma != 0 {
		cs.csvrdr.Comma = cs.comma
	}

	// Read the first row (may or may not be column header)
	var firstrow []string
	var err error
	firstrow, err = cs.csvrdr.Read()
	if err != nil {
		panic(err)
	}

	if len(cs.types) == 0 {
		panic("SetTypes must be called.")
	}
	cs.nvar = len(cs.types)

	cs.filenamepos = make(map[string]int)
	if cs.hasheader {
		for pos, na := range firstrow {
			cs.filenamepos[na] = pos
		}
	} else {
		// Save the first row since it contains data
		cs.firstrow = firstrow

		if len(firstrow) != len(cs.types) {
			msg := fmt.Sprintf("File has %d columns and no header, but types has %d values.", len(firstrow), len(cs.types))
			panic(msg)
		}

		// If no header, types contains the names in the proper order
		for pos, vt := range cs.types {
			cs.filenamepos[vt.Name] = pos
		}
	}

	cs.setNames()
	cs.setbdata()

	cs.doneinit = true
}

func (cs *CSVReader) setNames() {

	// By the time this is called, dtypes is guaranteed to exist.
	cs.dtypes = make([]Dtype, len(cs.types))
	cs.names = make([]string, len(cs.types))
	cs.namepos = make(map[string]int)
	cs.filepos = make([]int, len(cs.types))
	for pos, vt := range cs.types {
		cs.names[pos] = vt.Name
		cs.dtypes[pos] = vt.Type
		cs.namepos[vt.Name] = pos
		cs.filepos[pos] = cs.filenamepos[vt.Name]
	}
}

// ChunkSize sets the size of chunks for this Dstream, it can only
// be called before reading begins.
func (cs *CSVReader) ChunkSize(c int) *CSVReader {
	cs.chunkSize = c
	return cs
}

// Names returns the names of the variables in the dstream.
func (cs *CSVReader) Names() []string {
	return cs.names
}

// NumVar returns the number of variables in the dstream.
func (cs *CSVReader) NumVar() int {
	return cs.nvar
}

// NumObs returns the number of observations in the dstream.  If the
// dstream has not been fully read, returns -1.
func (cs *CSVReader) NumObs() int {
	if cs.done {
		return cs.nobs
	}
	return -1
}

// GetPos returns a chunk of a data column by column position.
func (cs *CSVReader) GetPos(j int) interface{} {
	return cs.bdata[j]
}

// Get returns a chunk of a data column by name.
func (cs *CSVReader) Get(na string) interface{} {
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
func (cs *CSVReader) Reset() {
	if !cs.doneinit {
		panic("cannot reset, Dstream has not been fully constructed")
	}

	if cs.nobs == 0 {
		return
	}

	r, ok := cs.rdr.(io.ReadSeeker)
	if !ok {
		panic("cannot reset")
	}
	_, err := r.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}
	cs.nobs = 0
	cs.done = false
	cs.chunknum = 0
	cs.rdr = r                   // is this needed?
	cs.csvrdr = csv.NewReader(r) // is this needed?

	// Skip over the header if needed.
	if cs.hasheader {
		_, err := cs.csvrdr.Read()
		if err != nil {
			panic(err)
		}
	}
}

// CSVWriter supports writing a Dstream to an io.Writer in csv format.
type CSVWriter struct {

	// The Dstream to be written.
	stream Dstream

	// Format for float type value
	floatFmt string

	// A slice of format types, stored per-variable.
	fmts []string

	wtr io.Writer
}

// ToCSV writes a Dstream in CSV format.  Call SetWriter or Filename
// to configure the underlying writer, then call additional methods
// for customization as desired, and finally call Done to complete the
// writing.
func ToCSV(d Dstream) *CSVWriter {
	c := &CSVWriter{
		stream: d,
	}
	return c
}

// FloatFmt sets the format string to be used when writing float
// values.  This value is ignored for columns specified in a call to
// the Formats method.
func (dw *CSVWriter) FloatFmt(fmt string) *CSVWriter {

	dw.floatFmt = fmt
	return dw
}

// Formats sets format strings to be used when writing the Dstream.
// The provided argument is a map from variable names to variable
// formats.
func (dw *CSVWriter) Formats(fmts map[string]string) *CSVWriter {

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
func (dw *CSVWriter) Filename(name string) *CSVWriter {

	var err error
	dw.wtr, err = os.Create(name)
	if err != nil {
		panic(err)
	}

	return dw
}

// SetWriter configures the CSVWriter to write to the given io stream.
func (dw *CSVWriter) SetWriter(w io.Writer) *CSVWriter {

	dw.wtr = w
	return dw
}

// getFmt is a utility for getting the format string for a given
// column.
func (dw *CSVWriter) getFmt(t string, col int) string {

	if dw.fmts != nil && dw.fmts[col] != "" {
		return dw.fmts[col]
	}

	switch t {
	case "float":
		if dw.floatFmt == "" {
			return "%.8f"
		}
		return dw.floatFmt
	case "int":
		return "%d"
	default:
		panic("unknown type")
	}
}

// Done completes writing a Dstream to a specified io.Writer in csv
// format.
func (dw *CSVWriter) Done() error {

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
