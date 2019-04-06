// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"compress/gzip"
	"encoding/gob"
	"os"
	"time"
)

// Save saves the DataFrame to a file.
func (df *DataFrame) Save(fname string) {

	fid, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer fid.Close()

	gid := gzip.NewWriter(fid)
	defer gid.Close()

	enc := gob.NewEncoder(gid)

	nchunk := uint64(len(df.data[0]))
	err = enc.Encode(&nchunk)
	if err != nil {
		panic(err)
	}

	err = enc.Encode(df.names)
	if err != nil {
		panic(err)
	}

	var dtypes []Dtype
	for j := range df.data {
		switch df.data[j][0].(type) {
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
		default:
			panic("unknown dtype")
		}
	}

	err = enc.Encode(dtypes)
	if err != nil {
		panic(err)
	}

	for j := range df.data {
		for i := range df.data[j] {
			switch x := df.data[j][i].(type) {
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

// Load loads the DataFrame from a file.
func (df *DataFrame) Load(fname string) {

	fid, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer fid.Close()

	gid, err := gzip.NewReader(fid)
	if err != nil {
		panic(err)
	}
	defer gid.Close()

	enc := gob.NewDecoder(gid)

	var nc uint64
	err = enc.Decode(&nc)
	if err != nil {
		panic(err)
	}
	nchunk := int(nc)

	var names []string
	err = enc.Decode(&names)
	if err != nil {
		panic(err)
	}
	nvar := len(names)

	var dtypes []Dtype
	err = enc.Decode(&dtypes)
	if err != nil {
		panic(err)
	}

	data := make([][]interface{}, nvar)
	for j := 0; j < nvar; j++ {
		data[j] = make([]interface{}, nchunk)
	}

	for j := 0; j < nvar; j++ {
		for i := 0; i < nchunk; i++ {
			switch dtypes[j] {
			case String:
				var x []string
				err := enc.Decode(&x)
				if err != nil {
					panic(err)
				}
				data[j][i] = x
			case Time:
				var x []time.Time
				err := enc.Decode(&x)
				if err != nil {
					panic(err)
				}
				data[j][i] = x
			case Uint8:
				var x []uint8
				err := enc.Decode(&x)
				if err != nil {
					panic(err)
				}
				data[j][i] = x
			case Uint16:
				var x []uint16
				err := enc.Decode(&x)
				if err != nil {
					panic(err)
				}
				data[j][i] = x
			case Uint32:
				var x []uint32
				err := enc.Decode(&x)
				if err != nil {
					panic(err)
				}
				data[j][i] = x
			case Uint64:
				var x []uint64
				err := enc.Decode(&x)
				if err != nil {
					panic(err)
				}
				data[j][i] = x
			case Int8:
				var x []int8
				err := enc.Decode(&x)
				if err != nil {
					panic(err)
				}
				data[j][i] = x
			case Int16:
				var x []int16
				err := enc.Decode(&x)
				if err != nil {
					panic(err)
				}
				data[j][i] = x
			case Int32:
				var x []int32
				err := enc.Decode(&x)
				if err != nil {
					panic(err)
				}
				data[j][i] = x
			case Int64:
				var x []int64
				err := enc.Decode(&x)
				if err != nil {
					panic(err)
				}
				data[j][i] = x
			case Float32:
				var x []float32
				err := enc.Decode(&x)
				if err != nil {
					panic(err)
				}
				data[j][i] = x
			case Float64:
				var x []float64
				err := enc.Decode(&x)
				if err != nil {
					panic(err)
				}
				data[j][i] = x
			}
		}
	}

	df.names = names
	df.data = data
}
