// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
	"time"
)

// Resize returns a slice of the given type having size n,
// re-using the provided slice if it is big enough.
func resizeString(x []string, n int) []string {
	if cap(x) < n {
		x = make([]string, n)
	}
	return x[0:n]
}

// Resize returns a slice of the given type having size n,
// re-using the provided slice if it is big enough.
func resizeTime(x []time.Time, n int) []time.Time {
	if cap(x) < n {
		x = make([]time.Time, n)
	}
	return x[0:n]
}

// Resize returns a slice of the given type having size n,
// re-using the provided slice if it is big enough.
func resizeUint8(x []uint8, n int) []uint8 {
	if cap(x) < n {
		x = make([]uint8, n)
	}
	return x[0:n]
}

// Resize returns a slice of the given type having size n,
// re-using the provided slice if it is big enough.
func resizeUint16(x []uint16, n int) []uint16 {
	if cap(x) < n {
		x = make([]uint16, n)
	}
	return x[0:n]
}

// Resize returns a slice of the given type having size n,
// re-using the provided slice if it is big enough.
func resizeUint32(x []uint32, n int) []uint32 {
	if cap(x) < n {
		x = make([]uint32, n)
	}
	return x[0:n]
}

// Resize returns a slice of the given type having size n,
// re-using the provided slice if it is big enough.
func resizeUint64(x []uint64, n int) []uint64 {
	if cap(x) < n {
		x = make([]uint64, n)
	}
	return x[0:n]
}

// Resize returns a slice of the given type having size n,
// re-using the provided slice if it is big enough.
func resizeInt8(x []int8, n int) []int8 {
	if cap(x) < n {
		x = make([]int8, n)
	}
	return x[0:n]
}

// Resize returns a slice of the given type having size n,
// re-using the provided slice if it is big enough.
func resizeInt16(x []int16, n int) []int16 {
	if cap(x) < n {
		x = make([]int16, n)
	}
	return x[0:n]
}

// Resize returns a slice of the given type having size n,
// re-using the provided slice if it is big enough.
func resizeInt32(x []int32, n int) []int32 {
	if cap(x) < n {
		x = make([]int32, n)
	}
	return x[0:n]
}

// Resize returns a slice of the given type having size n,
// re-using the provided slice if it is big enough.
func resizeInt64(x []int64, n int) []int64 {
	if cap(x) < n {
		x = make([]int64, n)
	}
	return x[0:n]
}

// Resize returns a slice of the given type having size n,
// re-using the provided slice if it is big enough.
func resizeFloat32(x []float32, n int) []float32 {
	if cap(x) < n {
		x = make([]float32, n)
	}
	return x[0:n]
}

// Resize returns a slice of the given type having size n,
// re-using the provided slice if it is big enough.
func resizeFloat64(x []float64, n int) []float64 {
	if cap(x) < n {
		x = make([]float64, n)
	}
	return x[0:n]
}

// VarTypes returns a map relating each variable by name to its corresponding
// data type.  The dstream should be readable before calling (Reset and call
// Next if needed).
func VarTypes(d Dstream) map[string]string {
	types := make(map[string]string)
	for k, na := range d.Names() {
		v := d.GetPos(k)
		switch v.(type) {
		case []string:
			types[na] = "string"
		case []time.Time:
			types[na] = "time.Time"
		case []uint8:
			types[na] = "uint8"
		case []uint16:
			types[na] = "uint16"
		case []uint32:
			types[na] = "uint32"
		case []uint64:
			types[na] = "uint64"
		case []int8:
			types[na] = "int8"
		case []int16:
			types[na] = "int16"
		case []int32:
			types[na] = "int32"
		case []int64:
			types[na] = "int64"
		case []float32:
			types[na] = "float32"
		case []float64:
			types[na] = "float64"
		default:
			types[na] = "unknown type"
		}
	}
	return types
}

func ilen(x interface{}) int {
	switch x := x.(type) {
	case []string:
		return len(x)
	case []time.Time:
		return len(x)
	case []uint8:
		return len(x)
	case []uint16:
		return len(x)
	case []uint32:
		return len(x)
	case []uint64:
		return len(x)
	case []int8:
		return len(x)
	case []int16:
		return len(x)
	case []int32:
		return len(x)
	case []int64:
		return len(x)
	case []float32:
		return len(x)
	case []float64:
		return len(x)
	case nil:
		return 0
	default:
		msg := fmt.Sprintf("unknown type: %T", x)
		panic(msg)
	}
}

func truncate(z []interface{}) {
	for j, x := range z {
		if x != nil {
			switch x := x.(type) {
			case []string:
				z[j] = x[0:0]
			case []time.Time:
				z[j] = x[0:0]
			case []uint8:
				z[j] = x[0:0]
			case []uint16:
				z[j] = x[0:0]
			case []uint32:
				z[j] = x[0:0]
			case []uint64:
				z[j] = x[0:0]
			case []int8:
				z[j] = x[0:0]
			case []int16:
				z[j] = x[0:0]
			case []int32:
				z[j] = x[0:0]
			case []int64:
				z[j] = x[0:0]
			case []float32:
				z[j] = x[0:0]
			case []float64:
				z[j] = x[0:0]
			default:
				msg := fmt.Sprintf("unknown type %T", x)
				panic(msg)
			}
		}
	}
}

// GetCol returns a copy of the data for one variable.  The data
// are returned as a slice.  The column is returned starting with the
// current chunk, call Reset to ensure that the column is extracted
// from the first chunk.
func GetCol(da Dstream, na string) interface{} {

	vn := da.Names()
	for j, x := range vn {
		if na == x {
			return GetColPos(da, j)
		}
	}
	panic(fmt.Sprintf("GetCol: variable '%s' not found", na))
	return nil
}

// GetColPos returns a copy of the data for one variable.
// The data are returned as a slice, which is a coy of the
// underlying data.
func GetColPos(da Dstream, j int) interface{} {

	da.Reset()

	if !da.Next() {
		return nil
	}

	// Get the first chunk so that we have the data type.
	v := da.GetPos(j)

	switch v := v.(type) {
	case []string:
		var x []string
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]string)
			x = append(x, y...)
		}
		return x
	case []time.Time:
		var x []time.Time
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]time.Time)
			x = append(x, y...)
		}
		return x
	case []uint8:
		var x []uint8
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]uint8)
			x = append(x, y...)
		}
		return x
	case []uint16:
		var x []uint16
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]uint16)
			x = append(x, y...)
		}
		return x
	case []uint32:
		var x []uint32
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]uint32)
			x = append(x, y...)
		}
		return x
	case []uint64:
		var x []uint64
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]uint64)
			x = append(x, y...)
		}
		return x
	case []int8:
		var x []int8
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]int8)
			x = append(x, y...)
		}
		return x
	case []int16:
		var x []int16
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]int16)
			x = append(x, y...)
		}
		return x
	case []int32:
		var x []int32
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]int32)
			x = append(x, y...)
		}
		return x
	case []int64:
		var x []int64
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]int64)
			x = append(x, y...)
		}
		return x
	case []float32:
		var x []float32
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]float32)
			x = append(x, y...)
		}
		return x
	case []float64:
		var x []float64
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]float64)
			x = append(x, y...)
		}
		return x
	case nil:
		return nil
	default:
		msg := fmt.Sprintf("GetCol: unknown data type %T.\n", v)
		panic(msg)
	}
}

// equalSlice returns true if x and y are slices of the same type and
// hold equal values.
func equalSlice(x, y interface{}) bool {

	switch u := x.(type) {
	case []string:
		v, ok := y.([]string)
		if !ok {
			return false
		}
		if len(u) != len(v) {
			return false
		}
		for i := range u {
			if u[i] != v[i] {
				return false
			}
		}
	case []time.Time:
		v, ok := y.([]time.Time)
		if !ok {
			return false
		}
		if len(u) != len(v) {
			return false
		}
		for i := range u {
			if u[i] != v[i] {
				return false
			}
		}
	case []uint8:
		v, ok := y.([]uint8)
		if !ok {
			return false
		}
		if len(u) != len(v) {
			return false
		}
		for i := range u {
			if u[i] != v[i] {
				return false
			}
		}
	case []uint16:
		v, ok := y.([]uint16)
		if !ok {
			return false
		}
		if len(u) != len(v) {
			return false
		}
		for i := range u {
			if u[i] != v[i] {
				return false
			}
		}
	case []uint32:
		v, ok := y.([]uint32)
		if !ok {
			return false
		}
		if len(u) != len(v) {
			return false
		}
		for i := range u {
			if u[i] != v[i] {
				return false
			}
		}
	case []uint64:
		v, ok := y.([]uint64)
		if !ok {
			return false
		}
		if len(u) != len(v) {
			return false
		}
		for i := range u {
			if u[i] != v[i] {
				return false
			}
		}
	case []int8:
		v, ok := y.([]int8)
		if !ok {
			return false
		}
		if len(u) != len(v) {
			return false
		}
		for i := range u {
			if u[i] != v[i] {
				return false
			}
		}
	case []int16:
		v, ok := y.([]int16)
		if !ok {
			return false
		}
		if len(u) != len(v) {
			return false
		}
		for i := range u {
			if u[i] != v[i] {
				return false
			}
		}
	case []int32:
		v, ok := y.([]int32)
		if !ok {
			return false
		}
		if len(u) != len(v) {
			return false
		}
		for i := range u {
			if u[i] != v[i] {
				return false
			}
		}
	case []int64:
		v, ok := y.([]int64)
		if !ok {
			return false
		}
		if len(u) != len(v) {
			return false
		}
		for i := range u {
			if u[i] != v[i] {
				return false
			}
		}
	case []float32:
		v, ok := y.([]float32)
		if !ok {
			return false
		}
		if len(u) != len(v) {
			return false
		}
		for i := range u {
			if u[i] != v[i] {
				return false
			}
		}
	case []float64:
		v, ok := y.([]float64)
		if !ok {
			return false
		}
		if len(u) != len(v) {
			return false
		}
		for i := range u {
			if u[i] != v[i] {
				return false
			}
		}
	default:
		panic("unkown type")
	}

	return true
}
