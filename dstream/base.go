package dstream

import (
	"fmt"
	"io"
	"os"
)

// Dstream streams chunks of data to a consumer.
type Dstream interface {

	// Next attempts to advance to the next chunk and returns true
	// if successful.
	Next() bool

	// Names returns the variable names.
	Names() []string

	// Get returns the values for one variable in the current
	// chunk, referring to the variable by name
	Get(string) interface{}

	// Get returns the values for one variable for the current
	// chunk, referring to the variable by position.
	GetPos(int) interface{}

	// NumVar returns the number of variables in the data set.
	NumVar() int

	// NumObs returns the number of rows in the data set, it may
	// return -1 if not known.
	NumObs() int

	// Reset sets the provider so that the data are read from the
	// beginning of the dataset.
	Reset()

	// Close frees any resources suh as file handles used by the dstream.
	Close()
}

// Dtype represents a data type
type Dtype uint8

// Uint8, etc. are constants defining data types
const (
	Uint8 Dtype = iota
	Uint16
	Uint32
	Uint64
	Int8
	Int16
	Int32
	Int64
	Float32
	Float64
	Time
	String
)

// DataFrame is an implementation of Dstream based on sharded arrays.
type DataFrame struct {
	xform // bdata is not used

	// data is the underlying data to be passed to the
	// consumer.
	data [][]interface{}

	chunk int // 1-based
	done  bool
	nobs  int
}

// Next advances to the next chunk and returns true if successful. If there are
// no more chunks, it returns false.
func (da *DataFrame) Next() bool {

	da.chunk++
	if da.chunk <= len(da.data[0]) {
		return true
	}
	da.done = true
	return false
}

// NumObs returns the number of observations in the DataFrame, if known. If the number
// of observations is not known, it returns -1.
func (da *DataFrame) NumObs() int {

	if da.nobs > 0 {
		return da.nobs
	}

	if da.data == nil || len(da.data) == 0 {
		// Not yet known
		return -1
	}

	var nobs int
	for _, v := range da.data[0] {
		nobs += ilen(v)
	}
	da.nobs = nobs

	return nobs
}

func (da *DataFrame) init() {
	da.setNamePos() // TODO should get rid of this
}

// Names returns the variable (column) names of the dstream.
func (da *DataFrame) Names() []string {
	return da.names
}

// Reset resets the dstream so that after the next call to Next, the
// dstream is at chunk zero.
func (da *DataFrame) Reset() {
	da.chunk = 0
	da.nobs = 0
	da.done = false
}

// GetPos returns the data slice for the variable at the given position.
func (da *DataFrame) GetPos(j int) interface{} {
	if da.done {
		return nil
	}

	return da.data[j][da.chunk-1]
}

// Get returns the data slice for the variable with the given name.
func (da *DataFrame) Get(na string) interface{} {

	pos := -1
	for j, nm := range da.Names() {
		if nm == na {
			pos = j
			break
		}
	}

	if pos == -1 {
		msg := fmt.Sprintf("Get: variable '%s' not found", na)
		panic(msg)
	}

	return da.GetPos(pos)
}

// NumVar returns the number of variables in the dstream.
func (da *DataFrame) NumVar() int {
	return len(da.data)
}

// NewFromArrays creates a Dstream from raw data stored as slices;
// data[i][j] is the data for the i^th variable in the j^th chunk.
func NewFromArrays(data [][]interface{}, names []string) Dstream {

	if len(data) != len(names) {
		msg := fmt.Sprintf("NewFromArrays: len(data) = %d != len(names) = %d",
			len(data), len(names))
		panic(msg)
	}

	da := &DataFrame{
		data: data,
		xform: xform{
			names: names,
		},
	}

	da.init()

	return da
}

// CheckValid runs thhrough the chunks and confirms that the lenghts of the slices within
// the chunks are the same.  If CheckValid returns false, the dstream is in a corrupted
// state.  On completion, the dstream is in its initial state.
func CheckValid(data Dstream) bool {

	data.Reset()
	names := data.Names()

	for c := 0; data.Next(); c++ {
		n0 := ilen(data.GetPos(0))
		for j := 1; j < len(names); j++ {
			n1 := ilen(data.GetPos(j))
			if n1 != n0 {
				msg := fmt.Sprintf("Length mismatch in chunk %d: len(%s) = %d, len(%s) = %d\n",
					c, names[0], n0, names[j], n1)
				io.WriteString(os.Stderr, msg)
				return false
			}
		}
	}

	data.Reset()

	return true
}

// NewFromFlat creates a Dstream from raw data stored as contiguous
// (flat) arrays.  data[i] is the data for the i^th variable.
func NewFromFlat(data []interface{}, names []string) Dstream {

	if len(data) != len(names) {
		msg := fmt.Sprintf("NewFromFlat: len(data) = %d != len(names) = %d",
			len(data), len(names))
		panic(msg)
	}

	var ar [][]interface{}
	for _, v := range data {
		ar = append(ar, []interface{}{v})
	}

	da := &DataFrame{
		data: ar,
		xform: xform{
			names: names,
		},
	}

	da.init()

	return da
}

// Shallow attempts to make a shallow copy of the data stream.
// Currently, only memory-backed data streams can be shallow copied,
// otherwise a deep copy is returned.  Shallow copies of the same data
// can be read in parallel.
func Shallow(data Dstream) Dstream {

	data.Reset()
	switch data := data.(type) {
	case *DataFrame:
		var dy DataFrame = *data
		return &dy
	default:
		return MemCopy(data)
	}
}
