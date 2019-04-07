// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
	"time"
)

func (r *replaceColumn) GetPos(j int) interface{} {

	if j != r.colpos {
		return r.source.GetPos(j)
	}

	switch z := r.coldata.(type) {
	case []string:
		return z[r.rowpos : r.rowpos+r.csize]
	case []time.Time:
		return z[r.rowpos : r.rowpos+r.csize]
	case []uint8:
		return z[r.rowpos : r.rowpos+r.csize]
	case []uint16:
		return z[r.rowpos : r.rowpos+r.csize]
	case []uint32:
		return z[r.rowpos : r.rowpos+r.csize]
	case []uint64:
		return z[r.rowpos : r.rowpos+r.csize]
	case []int8:
		return z[r.rowpos : r.rowpos+r.csize]
	case []int16:
		return z[r.rowpos : r.rowpos+r.csize]
	case []int32:
		return z[r.rowpos : r.rowpos+r.csize]
	case []int64:
		return z[r.rowpos : r.rowpos+r.csize]
	case []float32:
		return z[r.rowpos : r.rowpos+r.csize]
	case []float64:
		return z[r.rowpos : r.rowpos+r.csize]
	default:
		msg := fmt.Sprintf("unknown type %T\n", z)
		panic(msg)
	}
}
