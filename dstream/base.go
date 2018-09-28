package dstream

import "fmt"

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

// dataArrays is an implementation of Dstream based on sharded arrays.
type dataArrays struct {
	xform // bdata is not used

	// arrays is the underlying data to be passed to the
	// consumer.
	arrays [][]interface{}

	chunk int // 1-based
	done  bool
	nobs  int
}

func (da *dataArrays) Next() bool {

	da.chunk++
	if da.chunk <= len(da.arrays[0]) {
		return true
	}
	da.done = true
	return false
}

func (da *dataArrays) NumObs() int {

	if da.nobs > 0 {
		return da.nobs
	}

	if da.arrays == nil || len(da.arrays) == 0 {
		// Not yet known
		return -1
	}

	var nobs int
	for _, v := range da.arrays[0] {
		nobs += ilen(v)
	}
	da.nobs = nobs

	return nobs
}

func (da *dataArrays) init() {
	da.setNamePos() // TODO should get rid of this
}

func (da *dataArrays) Names() []string {
	return da.names
}

func (da *dataArrays) Reset() {
	da.chunk = 0
	da.nobs = 0
	da.done = false
}

func (da *dataArrays) GetPos(j int) interface{} {
	if da.done {
		return nil
	}

	return da.arrays[j][da.chunk-1]
}

func (da *dataArrays) Get(na string) interface{} {

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

func (da *dataArrays) NumVar() int {
	return len(da.arrays)
}

// NewFromArrays creates a Dstream from raw data stored as slices;
// data[i][j] is the data for the i^th variable in the j^th chunk.
func NewFromArrays(data [][]interface{}, names []string) Dstream {

	if len(data) != len(names) {
		msg := fmt.Sprintf("NewFromArrays: len(data) = %d != len(names) = %d",
			len(data), len(names))
		panic(msg)
	}

	da := &dataArrays{
		arrays: data,
		xform: xform{
			names: names,
		},
	}

	da.init()

	return da
}

// NewFromContigArrays creates a Dstream from raw data stored as contiguous
// arrays.  data[i] is the data for the i^th variable.
func NewFromContigArrays(data []interface{}, names []string) Dstream {

	if len(data) != len(names) {
		msg := fmt.Sprintf("NewFromContig: len(data) = %d != len(names) = %d",
			len(data), len(names))
		panic(msg)
	}

	var ar [][]interface{}
	for _, v := range data {
		ar = append(ar, []interface{}{v})
	}

	da := &dataArrays{
		arrays: ar,
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
	case *dataArrays:
		var dy dataArrays = *data
		return &dy
	default:
		return MemCopy(data)
	}
}

//go:generate go run gen.go -template=memcopy.template
