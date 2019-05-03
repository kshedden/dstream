// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"time"
)

// Next advances to the next chunk.
func (cs *CSVReader) Next() bool {

	if cs.done {
		return false
	}

	if cs.limitchunk > 0 && cs.limitchunk <= cs.chunknum {
		cs.done = true
		return false
	}

	cs.chunknum++

	truncate(cs.bdata)

	for j := 0; j < cs.chunkSize; j++ {

		// Try to read a row, return false if done.
		var rec []string
		var err error
		if cs.firstrow != nil {
			rec = cs.firstrow
			cs.firstrow = nil
		} else {
			rec, err = cs.csvrdr.Read()
			if err == io.EOF {
				cs.done = true
				return ilen(cs.bdata[0]) > 0
			} else if err != nil {
				if cs.skipErrors {
					os.Stderr.WriteString(fmt.Sprintf("%v\n", err))
					continue
				}
				panic(err)
			}
		}
		cs.nobs++

		for pos, typ := range cs.types {
			fpos := cs.filepos[pos]
			switch typ.Type {
			case String:
				x := rec[fpos]
				u := cs.bdata[pos].([]string)
				cs.bdata[pos] = append(u, string(x))
			case Time:
				x := cs.parseTime(rec[fpos])
				u := cs.bdata[pos].([]time.Time)
				cs.bdata[pos] = append(u, time.Time(x))
			case Uint8:
				x, err := strconv.Atoi(rec[fpos])
				if err != nil {
					panic(err)
				}
				u := cs.bdata[pos].([]uint8)
				cs.bdata[pos] = append(u, uint8(x))
			case Uint16:
				x, err := strconv.Atoi(rec[fpos])
				if err != nil {
					panic(err)
				}
				u := cs.bdata[pos].([]uint16)
				cs.bdata[pos] = append(u, uint16(x))
			case Uint32:
				x, err := strconv.Atoi(rec[fpos])
				if err != nil {
					panic(err)
				}
				u := cs.bdata[pos].([]uint32)
				cs.bdata[pos] = append(u, uint32(x))
			case Uint64:
				x, err := strconv.Atoi(rec[fpos])
				if err != nil {
					panic(err)
				}
				u := cs.bdata[pos].([]uint64)
				cs.bdata[pos] = append(u, uint64(x))
			case Int8:
				x, err := strconv.Atoi(rec[fpos])
				if err != nil {
					panic(err)
				}
				u := cs.bdata[pos].([]int8)
				cs.bdata[pos] = append(u, int8(x))
			case Int16:
				x, err := strconv.Atoi(rec[fpos])
				if err != nil {
					panic(err)
				}
				u := cs.bdata[pos].([]int16)
				cs.bdata[pos] = append(u, int16(x))
			case Int32:
				x, err := strconv.Atoi(rec[fpos])
				if err != nil {
					panic(err)
				}
				u := cs.bdata[pos].([]int32)
				cs.bdata[pos] = append(u, int32(x))
			case Int64:
				x, err := strconv.Atoi(rec[fpos])
				if err != nil {
					panic(err)
				}
				u := cs.bdata[pos].([]int64)
				cs.bdata[pos] = append(u, int64(x))
			case Float32:
				x, err := strconv.ParseFloat(rec[fpos], 64)
				if err != nil {
					x = math.NaN()
				}
				u := cs.bdata[pos].([]float32)
				cs.bdata[pos] = append(u, float32(x))
			case Float64:
				x, err := strconv.ParseFloat(rec[fpos], 64)
				if err != nil {
					x = math.NaN()
				}
				u := cs.bdata[pos].([]float64)
				cs.bdata[pos] = append(u, float64(x))
			default:
				panic("unknown type")
			}
		}
	}

	return true
}

func (cs *CSVReader) setbdata() {

	cs.bdata = make([]interface{}, len(cs.names))

	for pos, dtype := range cs.dtypes {
		switch dtype {
		case String:
			cs.bdata[pos] = make([]string, 0)
		case Time:
			cs.bdata[pos] = make([]time.Time, 0)
		case Uint8:
			cs.bdata[pos] = make([]uint8, 0)
		case Uint16:
			cs.bdata[pos] = make([]uint16, 0)
		case Uint32:
			cs.bdata[pos] = make([]uint32, 0)
		case Uint64:
			cs.bdata[pos] = make([]uint64, 0)
		case Int8:
			cs.bdata[pos] = make([]int8, 0)
		case Int16:
			cs.bdata[pos] = make([]int16, 0)
		case Int32:
			cs.bdata[pos] = make([]int32, 0)
		case Int64:
			cs.bdata[pos] = make([]int64, 0)
		case Float32:
			cs.bdata[pos] = make([]float32, 0)
		case Float64:
			cs.bdata[pos] = make([]float64, 0)
		default:
			panic("Unknown type")
		}
	}
}
