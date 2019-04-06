// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
	"time"
)

func (fc *filterCol) Next() bool {

	if !fc.source.Next() {
		fc.nobsKnown = true
		return false
	}

	n := ilen(fc.source.GetPos(0))
	if n == 0 {
		return true
	}

	fc.keep = resizeBool(fc.keep, n)
	for i := range fc.keep {
		fc.keep[i] = true
	}

	vm := VarMap(fc.source)
	fc.filter(vm, fc.keep)

	fc.keeppos = fc.keeppos[0:0]
	for j := range fc.keep {
		if fc.keep[j] {
			fc.keeppos = append(fc.keeppos, j)
		}
	}
	fc.nobs += len(fc.keeppos)

	q := len(fc.keeppos)
	for k, na := range fc.source.Names() {
		v := fc.source.Get(na)
		switch x := v.(type) {
		case []string:
			var u []string
			if fc.bdata[k] != nil {
				u = fc.bdata[k].([]string)
			}
			u = resizeString(u, q)
			u = u[0:q]
			for i, j := range fc.keeppos {
				u[i] = x[j]
			}
			fc.bdata[k] = u
		case []time.Time:
			var u []time.Time
			if fc.bdata[k] != nil {
				u = fc.bdata[k].([]time.Time)
			}
			u = resizeTime(u, q)
			u = u[0:q]
			for i, j := range fc.keeppos {
				u[i] = x[j]
			}
			fc.bdata[k] = u
		case []uint8:
			var u []uint8
			if fc.bdata[k] != nil {
				u = fc.bdata[k].([]uint8)
			}
			u = resizeUint8(u, q)
			u = u[0:q]
			for i, j := range fc.keeppos {
				u[i] = x[j]
			}
			fc.bdata[k] = u
		case []uint16:
			var u []uint16
			if fc.bdata[k] != nil {
				u = fc.bdata[k].([]uint16)
			}
			u = resizeUint16(u, q)
			u = u[0:q]
			for i, j := range fc.keeppos {
				u[i] = x[j]
			}
			fc.bdata[k] = u
		case []uint32:
			var u []uint32
			if fc.bdata[k] != nil {
				u = fc.bdata[k].([]uint32)
			}
			u = resizeUint32(u, q)
			u = u[0:q]
			for i, j := range fc.keeppos {
				u[i] = x[j]
			}
			fc.bdata[k] = u
		case []uint64:
			var u []uint64
			if fc.bdata[k] != nil {
				u = fc.bdata[k].([]uint64)
			}
			u = resizeUint64(u, q)
			u = u[0:q]
			for i, j := range fc.keeppos {
				u[i] = x[j]
			}
			fc.bdata[k] = u
		case []int8:
			var u []int8
			if fc.bdata[k] != nil {
				u = fc.bdata[k].([]int8)
			}
			u = resizeInt8(u, q)
			u = u[0:q]
			for i, j := range fc.keeppos {
				u[i] = x[j]
			}
			fc.bdata[k] = u
		case []int16:
			var u []int16
			if fc.bdata[k] != nil {
				u = fc.bdata[k].([]int16)
			}
			u = resizeInt16(u, q)
			u = u[0:q]
			for i, j := range fc.keeppos {
				u[i] = x[j]
			}
			fc.bdata[k] = u
		case []int32:
			var u []int32
			if fc.bdata[k] != nil {
				u = fc.bdata[k].([]int32)
			}
			u = resizeInt32(u, q)
			u = u[0:q]
			for i, j := range fc.keeppos {
				u[i] = x[j]
			}
			fc.bdata[k] = u
		case []int64:
			var u []int64
			if fc.bdata[k] != nil {
				u = fc.bdata[k].([]int64)
			}
			u = resizeInt64(u, q)
			u = u[0:q]
			for i, j := range fc.keeppos {
				u[i] = x[j]
			}
			fc.bdata[k] = u
		case []float32:
			var u []float32
			if fc.bdata[k] != nil {
				u = fc.bdata[k].([]float32)
			}
			u = resizeFloat32(u, q)
			u = u[0:q]
			for i, j := range fc.keeppos {
				u[i] = x[j]
			}
			fc.bdata[k] = u
		case []float64:
			var u []float64
			if fc.bdata[k] != nil {
				u = fc.bdata[k].([]float64)
			}
			u = resizeFloat64(u, q)
			u = u[0:q]
			for i, j := range fc.keeppos {
				u[i] = x[j]
			}
			fc.bdata[k] = u
		default:
			msg := fmt.Sprintf("Unkown data type '%T'\n", v)
			panic(msg)
		}
	}

	return true
}
