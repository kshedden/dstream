package dstream

//go:generate go run gen.go -template=bcols.template -numeric

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

// BCols takes data stored in a column-wise compressed format under
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

func (bc *bcols) Include(vars []string) *bcols {
	bc.include = vars
	return bc
}

func (bc *bcols) Exclude(vars []string) *bcols {
	bc.exclude = vars
	return bc
}

func (bc *bcols) Done() *bcols {
	bc.init()
	return bc
}

// usenames returns the variable names to include in the Dstream.
func (b *bcols) usenames() []string {

	inc := make(map[string]bool)
	exc := make(map[string]bool)

	for _, v := range b.include {
		inc[v] = true
	}
	for _, v := range b.exclude {
		exc[v] = true
	}

	var use []string

	for v, _ := range b.dtypes {

		if inc[v] && exc[v] {
			msg := fmt.Sprintf("%s is in both the 'include' and 'exclude' lists",
				v)
			panic(msg)
		}

		if exc[v] {
			continue
		}

		if len(exc) > 0 && exc[v] {
			continue
		}

		use = append(use, v)
	}

	sort.StringSlice(use).Sort()

	return use
}

func (b *bcols) Names() []string {
	return b.names
}

func (b *bcols) Close() {
	for _, x := range b.toclose {
		x.Close()
	}
}

func (b *bcols) Reset() {
	b.Close()
	b.toclose = b.toclose[0:0]
	b.rdrs = b.rdrs[0:0]
	b.init()
	b.nobsKnown = false
	b.nobs = 0
	b.done = false
}

func (b *bcols) NumVar() int {
	return len(b.names)
}

func (b *bcols) NumObs() int {
	if !b.nobsKnown {
		return -1
	}
	return b.nobs
}

func (b *bcols) GetPos(j int) interface{} {
	return b.bdata[j]
}

func (b *bcols) Get(na string) interface{} {
	pos, ok := b.namepos[na]
	if !ok {
		msg := fmt.Sprintf("Variable '%s' not found", na)
		panic(msg)
	}
	return b.bdata[pos]
}

type toBCols struct {

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

func ToBCols(d Dstream) *toBCols {

	return &toBCols{
		stream: d,
	}
}

func (tb *toBCols) Path(p string) *toBCols {

	tb.path = p

	err := os.MkdirAll(p, 0770)
	if err != nil {
		panic(err)
	}

	return tb
}

func (tb *toBCols) init() {

	if tb.path == "" {
		msg := "Path value not set"
		panic(msg)
	}

	// Default compression type
	if tb.cmpr == "" {
		tb.cmpr = "gz"
	}

	names := tb.stream.Names()

	tb.writeDtypes()

	for _, na := range names {

		na += ".bin"
		if tb.cmpr == "gz" {
			na += ".gz"
		} else if tb.cmpr == "sz" {
			na += ".sz"
		} else {
			msg := fmt.Sprintf("Compression type %s not recognized", tb.cmpr)
			panic(msg)
		}

		// Create the file
		fn := path.Join(tb.path, na)
		f, err := os.Create(fn)
		if err != nil {
			panic(err)
		}
		tb.needsClosing = append(tb.needsClosing, f)

		// Wrap it in a compressor if needed
		switch tb.cmpr {
		case "gz":
			g := gzip.NewWriter(f)
			tb.needsClosing = append(tb.needsClosing, g)
			tb.wtrs = append(tb.wtrs, g)
		case "sz":
			g := snappy.NewWriter(f)
			tb.needsClosing = append(tb.needsClosing, g)
			tb.wtrs = append(tb.wtrs, g)
		default:
			tb.wtrs = append(tb.wtrs, f)
		}
	}
}

func (tb *toBCols) Done() {

	tb.init()
	tb.write()

	// Need to process in reverse order so that nested writers are
	// closed inside-out.
	for j := len(tb.needsClosing) - 1; j >= 0; j-- {
		f := tb.needsClosing[j]
		f.Close()
	}
}
