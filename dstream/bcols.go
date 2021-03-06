package dstream

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"sort"

	"github.com/golang/snappy"
)

type bcols struct {
	bpath string

	dtypes map[string]string

	// Data types for all variables in the source directory, including
	// those not included in the Dstream when reading.
	dtypesAll map[string]string

	names []string

	bdata []interface{}

	rdrs []*bufio.Reader

	// io.Readers, etc. that need closing
	toclose []io.Closer

	chunksize int

	namepos map[string]int

	done     bool
	doneInit bool

	nobsKnown bool
	nobs      int

	// Names to include or exclude.  If include is empty it is ignored.
	include []string
	exclude []string
}

// NewBCols takes data stored in a column-wise compressed format under
// the given directory path, and returns it as a Dstream.  'path' is
// the directory containing the data, 'chunksize' is the number of
// values included in each chunk, 'include' and 'exclude' are lists of
// variable names to include (exclude) respectively.
//
// The underlying bcols format is simple.  Each column of data is
// stored in its own file, in binary native format, compressed using
// either gzip or snappy compression.  A file called 'dtypes.json'
// contains a dictionary mapping variable names to data types.
func NewBCols(path string, chunksize int) *bcols {

	b := &bcols{
		bpath:     path,
		chunksize: chunksize,
	}

	return b
}

// Include specifies variables that are included when reading the data from BCols
// storage into a Dstream.  If Include is not called, all variables not listed in
// a call to Exclude are read.
func (bc *bcols) Include(vars ...string) *bcols {
	bc.include = vars
	return bc
}

// Exclude specifies variables that are not read from the Bcols storage.
func (bc *bcols) Exclude(vars ...string) *bcols {
	bc.exclude = vars
	return bc
}

// Done is called to signal the conclusion of configuring the BCols reader or
// writer.
func (bc *bcols) Done() Dstream {
	bc.init()
	return bc
}

// usenames returns the variable names to include in the Dstream.
func (bc *bcols) usenames() []string {

	inc := make(map[string]bool)
	for _, v := range bc.include {
		inc[v] = true

		_, ok := bc.dtypesAll[v]
		if !ok {
			msg := fmt.Sprintf("Variable '%s' does not exist\n", v)
			panic(msg)
		}
	}

	// If no variables are included, default is to set include to
	// equal all variable names.
	if len(bc.include) == 0 {
		for k := range bc.dtypesAll {
			inc[k] = true
		}
	}

	exc := make(map[string]bool)
	for _, v := range bc.exclude {
		exc[v] = true

		_, ok := bc.dtypesAll[v]
		if !ok {
			msg := fmt.Sprintf("Variable '%s' does not exist\n", v)
			panic(msg)
		}
	}

	var use []string
	for k := range inc {
		if !exc[k] {
			use = append(use, k)
		}
	}

	sort.StringSlice(use).Sort()

	return use
}

func (bc *bcols) Names() []string {
	return bc.names
}

func (bc *bcols) Close() {
	for _, x := range bc.toclose {
		x.Close()
	}
}

func (bc *bcols) Reset() {
	bc.Close()
	bc.toclose = bc.toclose[0:0]
	bc.rdrs = bc.rdrs[0:0]
	bc.init()
	bc.nobsKnown = false
	bc.nobs = 0
	bc.done = false
}

func (bc *bcols) NumVar() int {
	return len(bc.names)
}

func (bc *bcols) NumObs() int {
	if !bc.nobsKnown {
		return -1
	}
	return bc.nobs
}

func (bc *bcols) GetPos(j int) interface{} {
	return bc.bdata[j]
}

func (bc *bcols) Get(na string) interface{} {
	pos, ok := bc.namepos[na]
	if !ok {
		msg := fmt.Sprintf("Variable '%s' not found", na)
		panic(msg)
	}
	return bc.bdata[pos]
}

// BColsWriter writes a dstream to disk in bcols format.
type BColsWriter struct {

	// The source data to be written
	stream Dstream

	// The directory where the results will be written
	path string

	// either "sz" (snappy), "gz" (gzip), or empty string
	cmpr string

	// A writer for each variable
	wtrs []io.Writer

	// All io values needing closing
	needsClosing []io.Closer
}

// NewBColsWriter creates a new BColsWriter that writes the given
// dstream.
func NewBColsWriter(d Dstream) *BColsWriter {

	return &BColsWriter{
		stream: d,
	}
}

// Path sets the location (a directory path) to which the data are
// written.
func (bw *BColsWriter) Path(p string) *BColsWriter {

	bw.path = p

	err := os.MkdirAll(p, 0770)
	if err != nil {
		panic(err)
	}

	return bw
}

func (bw *BColsWriter) init() {

	if bw.path == "" {
		msg := "Path value not set"
		panic(msg)
	}

	// Default compression type
	if bw.cmpr == "" {
		bw.cmpr = "gz"
	}

	names := bw.stream.Names()

	bw.writeDtypes()

	for _, na := range names {

		na += ".bin"
		if bw.cmpr == "gz" {
			na += ".gz"
		} else if bw.cmpr == "sz" {
			na += ".sz"
		} else {
			msg := fmt.Sprintf("Compression type %s not recognized", bw.cmpr)
			panic(msg)
		}

		// Create the file
		fn := path.Join(bw.path, na)
		f, err := os.Create(fn)
		if err != nil {
			panic(err)
		}
		bw.needsClosing = append(bw.needsClosing, f)

		// Wrap it in a compressor if needed
		switch bw.cmpr {
		case "gz":
			g := gzip.NewWriter(f)
			bw.needsClosing = append(bw.needsClosing, g)
			bw.wtrs = append(bw.wtrs, g)
		case "sz":
			g := snappy.NewBufferedWriter(f)
			bw.needsClosing = append(bw.needsClosing, g)
			bw.wtrs = append(bw.wtrs, g)
		default:
			bw.wtrs = append(bw.wtrs, f)
		}
	}
}

// Done flushes the data to disk.
func (bw *BColsWriter) Done() {

	bw.init()
	bw.write()

	// Need to process in reverse order so that nested writers are
	// closed inside-out.
	for j := len(bw.needsClosing) - 1; j >= 0; j-- {
		f := bw.needsClosing[j]
		f.Close()
	}
}
