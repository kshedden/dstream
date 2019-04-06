// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
	"time"
)

func (gen *generate) init() {

	for _, na := range gen.source.Names() {
		if gen.newvarname == na {
			msg := fmt.Sprintf("Generate: variable '%s' already exists.\n", gen.newvarname)
			panic(msg)
		}
	}

	gen.names = append(gen.source.Names(), gen.newvarname)
	gen.bdata = make([]interface{}, len(gen.names))

	switch gen.dtype {
	case String:
		gen.bdata[len(gen.bdata)-1] = make([]string, 0)

	case Time:
		gen.bdata[len(gen.bdata)-1] = make([]time.Time, 0)

	case Uint8:
		gen.bdata[len(gen.bdata)-1] = make([]uint8, 0)

	case Uint16:
		gen.bdata[len(gen.bdata)-1] = make([]uint16, 0)

	case Uint32:
		gen.bdata[len(gen.bdata)-1] = make([]uint32, 0)

	case Uint64:
		gen.bdata[len(gen.bdata)-1] = make([]uint64, 0)

	case Int8:
		gen.bdata[len(gen.bdata)-1] = make([]int8, 0)

	case Int16:
		gen.bdata[len(gen.bdata)-1] = make([]int16, 0)

	case Int32:
		gen.bdata[len(gen.bdata)-1] = make([]int32, 0)

	case Int64:
		gen.bdata[len(gen.bdata)-1] = make([]int64, 0)

	case Float32:
		gen.bdata[len(gen.bdata)-1] = make([]float32, 0)

	case Float64:
		gen.bdata[len(gen.bdata)-1] = make([]float64, 0)

	default:
		panic("Generate: unknown dtype")
	}
}

func (gen *generate) Next() bool {

	if !gen.source.Next() {
		return false
	}

	// All but new variable
	for j := 0; j < gen.source.NumVar(); j++ {
		gen.bdata[j] = gen.source.GetPos(j)
	}

	n := ilen(gen.GetPos(0))

	// The new variable goes in the last position
	q := len(gen.names) - 1

	switch x := gen.bdata[q].(type) {
	case []string:
		gen.bdata[q] = resizeString(x, n)

	case []time.Time:
		gen.bdata[q] = resizeTime(x, n)

	case []uint8:
		gen.bdata[q] = resizeUint8(x, n)

	case []uint16:
		gen.bdata[q] = resizeUint16(x, n)

	case []uint32:
		gen.bdata[q] = resizeUint32(x, n)

	case []uint64:
		gen.bdata[q] = resizeUint64(x, n)

	case []int8:
		gen.bdata[q] = resizeInt8(x, n)

	case []int16:
		gen.bdata[q] = resizeInt16(x, n)

	case []int32:
		gen.bdata[q] = resizeInt32(x, n)

	case []int64:
		gen.bdata[q] = resizeInt64(x, n)

	case []float32:
		gen.bdata[q] = resizeFloat32(x, n)

	case []float64:
		gen.bdata[q] = resizeFloat64(x, n)

	default:
		panic("unknown type")
	}

	mp := make(map[string]interface{})
	for k, na := range gen.names {
		if na != gen.newvarname {
			mp[na] = gen.bdata[k]
		}
	}

	gen.gfunc(mp, gen.bdata[q])
	return true
}
