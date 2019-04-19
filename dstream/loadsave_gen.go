// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"compress/gzip"
	"encoding/gob"
	"io"
	"os"
	"time"
)

// Save saves the dstream to a file.  It can subsequently be reloaded with Load.
func Save(ds Dstream, fname string) {

	ds.Reset()

	fid, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer fid.Close()

	gid := gzip.NewWriter(fid)
	defer gid.Close()

	enc := gob.NewEncoder(gid)

	err = enc.Encode(ds.Names())
	if err != nil {
		panic(err)
	}

	first := true
	for ds.Next() {

		if first {
			var dtypes []Dtype
			for j := 0; j < ds.NumVar(); j++ {
				switch ds.GetPos(j).(type) {
				case []string:
					dtypes = append(dtypes, String)
				case []time.Time:
					dtypes = append(dtypes, Time)
				case []uint8:
					dtypes = append(dtypes, Uint8)
				case []uint16:
					dtypes = append(dtypes, Uint16)
				case []uint32:
					dtypes = append(dtypes, Uint32)
				case []uint64:
					dtypes = append(dtypes, Uint64)
				case []int8:
					dtypes = append(dtypes, Int8)
				case []int16:
					dtypes = append(dtypes, Int16)
				case []int32:
					dtypes = append(dtypes, Int32)
				case []int64:
					dtypes = append(dtypes, Int64)
				case []float32:
					dtypes = append(dtypes, Float32)
				case []float64:
					dtypes = append(dtypes, Float64)
				}
			}

			err := enc.Encode(&dtypes)
			if err != nil {
				panic(err)
			}

			first = false
		}

		for j := 0; j < ds.NumVar(); j++ {
			switch x := ds.GetPos(j).(type) {
			case []string:
				err := enc.Encode(&x)
				if err != nil {
					panic(err)
				}
			case []time.Time:
				err := enc.Encode(&x)
				if err != nil {
					panic(err)
				}
			case []uint8:
				err := enc.Encode(&x)
				if err != nil {
					panic(err)
				}
			case []uint16:
				err := enc.Encode(&x)
				if err != nil {
					panic(err)
				}
			case []uint32:
				err := enc.Encode(&x)
				if err != nil {
					panic(err)
				}
			case []uint64:
				err := enc.Encode(&x)
				if err != nil {
					panic(err)
				}
			case []int8:
				err := enc.Encode(&x)
				if err != nil {
					panic(err)
				}
			case []int16:
				err := enc.Encode(&x)
				if err != nil {
					panic(err)
				}
			case []int32:
				err := enc.Encode(&x)
				if err != nil {
					panic(err)
				}
			case []int64:
				err := enc.Encode(&x)
				if err != nil {
					panic(err)
				}
			case []float32:
				err := enc.Encode(&x)
				if err != nil {
					panic(err)
				}
			case []float64:
				err := enc.Encode(&x)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

// load is a dstream that reads its data from a file that has been constructed using Save.
type load struct {

	// done becomes true when the whole data source has been read
	done bool

	// The name of the file containing the data
	filename string

	// The variable names
	names []string

	// The number of observations, not relevant until reading is complete.
	nobs int

	// The data types of the variables
	dtypes []Dtype

	// Use this to decode gob data from the file.
	dec *gob.Decoder

	// The file and gzip reader should be closed when done.
	needsClosing []io.Closer

	// namespos maps variable names to their column positions
	namespos map[string]int

	// This holds the data for the current chunk
	data []interface{}
}

// NewLoad returns a dstream that loads data from the given file.  The file must be created
// using the Save function.
func NewLoad(fname string) Dstream {

	var ld load
	ld.filename = fname
	ld.init()

	return &ld
}

func (ld *load) init() {

	fid, err := os.Open(ld.filename)
	if err != nil {
		panic(err)
	}
	ld.needsClosing = append(ld.needsClosing, fid)

	gid, err := gzip.NewReader(fid)
	if err != nil {
		panic(err)
	}
	ld.needsClosing = append(ld.needsClosing, gid)

	ld.dec = gob.NewDecoder(gid)

	err = ld.dec.Decode(&ld.names)
	if err != nil {
		panic(err)
	}

	ld.namespos = make(map[string]int)
	for k, na := range ld.names {
		ld.namespos[na] = k
	}

	err = ld.dec.Decode(&ld.dtypes)
	if err != nil {
		panic(err)
	}

	ld.data = make([]interface{}, len(ld.names))
}

// Reset resets the loader so that the data can be re-read from the beginning.
func (ld *load) Reset() {
	ld.Close()
	ld.done = false
	ld.nobs = 0
	ld.init()
}

// GetPos returns the data for the variable at the given position in the current chunk.
func (ld *load) GetPos(j int) interface{} {
	return ld.data[j]
}

// Names returns the variable names.
func (ld *load) Names() []string {
	return ld.names
}

// NumObs returns the number of observations (data rows) in the dstream.  A -1 is
// returned if the source has not been completely read.
func (ld *load) NumObs() int {
	if !ld.done {
		return -1
	}
	return ld.nobs
}

// Get returns the data for the given named variable in the current chunk.
func (ld *load) Get(na string) interface{} {
	return ld.GetPos(ld.namespos[na])
}

// Close all io resources associated with the value.
func (ld *load) Close() {

	for _, f := range ld.needsClosing {
		f.Close()
	}

	ld.needsClosing = nil
}

// NumVar returns the number of variables in the dstream.
func (ld *load) NumVar() int {
	return len(ld.names)
}

// Next advances to the next chunk and returns true, or returns false if
// all chunks have been read.
func (ld *load) Next() bool {

	// Loop over variables within chunks
	for j := 0; j < ld.NumVar(); j++ {
		switch ld.dtypes[j] {
		case String:
			var x []string
			err := ld.dec.Decode(&x)
			if err == io.EOF {
				if j != 0 {
					panic("File is corrupt")
				}
				ld.done = true
				return false
			} else if err != nil {
				panic(err)
			}
			ld.data[j] = x
			if j == 0 {
				ld.nobs += len(x)
			}
		case Time:
			var x []time.Time
			err := ld.dec.Decode(&x)
			if err == io.EOF {
				if j != 0 {
					panic("File is corrupt")
				}
				ld.done = true
				return false
			} else if err != nil {
				panic(err)
			}
			ld.data[j] = x
			if j == 0 {
				ld.nobs += len(x)
			}
		case Uint8:
			var x []uint8
			err := ld.dec.Decode(&x)
			if err == io.EOF {
				if j != 0 {
					panic("File is corrupt")
				}
				ld.done = true
				return false
			} else if err != nil {
				panic(err)
			}
			ld.data[j] = x
			if j == 0 {
				ld.nobs += len(x)
			}
		case Uint16:
			var x []uint16
			err := ld.dec.Decode(&x)
			if err == io.EOF {
				if j != 0 {
					panic("File is corrupt")
				}
				ld.done = true
				return false
			} else if err != nil {
				panic(err)
			}
			ld.data[j] = x
			if j == 0 {
				ld.nobs += len(x)
			}
		case Uint32:
			var x []uint32
			err := ld.dec.Decode(&x)
			if err == io.EOF {
				if j != 0 {
					panic("File is corrupt")
				}
				ld.done = true
				return false
			} else if err != nil {
				panic(err)
			}
			ld.data[j] = x
			if j == 0 {
				ld.nobs += len(x)
			}
		case Uint64:
			var x []uint64
			err := ld.dec.Decode(&x)
			if err == io.EOF {
				if j != 0 {
					panic("File is corrupt")
				}
				ld.done = true
				return false
			} else if err != nil {
				panic(err)
			}
			ld.data[j] = x
			if j == 0 {
				ld.nobs += len(x)
			}
		case Int8:
			var x []int8
			err := ld.dec.Decode(&x)
			if err == io.EOF {
				if j != 0 {
					panic("File is corrupt")
				}
				ld.done = true
				return false
			} else if err != nil {
				panic(err)
			}
			ld.data[j] = x
			if j == 0 {
				ld.nobs += len(x)
			}
		case Int16:
			var x []int16
			err := ld.dec.Decode(&x)
			if err == io.EOF {
				if j != 0 {
					panic("File is corrupt")
				}
				ld.done = true
				return false
			} else if err != nil {
				panic(err)
			}
			ld.data[j] = x
			if j == 0 {
				ld.nobs += len(x)
			}
		case Int32:
			var x []int32
			err := ld.dec.Decode(&x)
			if err == io.EOF {
				if j != 0 {
					panic("File is corrupt")
				}
				ld.done = true
				return false
			} else if err != nil {
				panic(err)
			}
			ld.data[j] = x
			if j == 0 {
				ld.nobs += len(x)
			}
		case Int64:
			var x []int64
			err := ld.dec.Decode(&x)
			if err == io.EOF {
				if j != 0 {
					panic("File is corrupt")
				}
				ld.done = true
				return false
			} else if err != nil {
				panic(err)
			}
			ld.data[j] = x
			if j == 0 {
				ld.nobs += len(x)
			}
		case Float32:
			var x []float32
			err := ld.dec.Decode(&x)
			if err == io.EOF {
				if j != 0 {
					panic("File is corrupt")
				}
				ld.done = true
				return false
			} else if err != nil {
				panic(err)
			}
			ld.data[j] = x
			if j == 0 {
				ld.nobs += len(x)
			}
		case Float64:
			var x []float64
			err := ld.dec.Decode(&x)
			if err == io.EOF {
				if j != 0 {
					panic("File is corrupt")
				}
				ld.done = true
				return false
			} else if err != nil {
				panic(err)
			}
			ld.data[j] = x
			if j == 0 {
				ld.nobs += len(x)
			}
		}
	}

	return true
}
