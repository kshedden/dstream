// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
)

func resizefloat64(x []float64, n int) []float64 {
	if cap(x) < n {
		x = make([]float64, n)
	}
	return x[0:n]
}

func resizefloat32(x []float32, n int) []float32 {
	if cap(x) < n {
		x = make([]float32, n)
	}
	return x[0:n]
}

func resizeuint64(x []uint64, n int) []uint64 {
	if cap(x) < n {
		x = make([]uint64, n)
	}
	return x[0:n]
}

func resizeuint32(x []uint32, n int) []uint32 {
	if cap(x) < n {
		x = make([]uint32, n)
	}
	return x[0:n]
}

func resizeuint16(x []uint16, n int) []uint16 {
	if cap(x) < n {
		x = make([]uint16, n)
	}
	return x[0:n]
}

func resizeuint8(x []uint8, n int) []uint8 {
	if cap(x) < n {
		x = make([]uint8, n)
	}
	return x[0:n]
}

func resizeint64(x []int64, n int) []int64 {
	if cap(x) < n {
		x = make([]int64, n)
	}
	return x[0:n]
}

func resizeint32(x []int32, n int) []int32 {
	if cap(x) < n {
		x = make([]int32, n)
	}
	return x[0:n]
}

func resizeint16(x []int16, n int) []int16 {
	if cap(x) < n {
		x = make([]int16, n)
	}
	return x[0:n]
}

func resizeint8(x []int8, n int) []int8 {
	if cap(x) < n {
		x = make([]int8, n)
	}
	return x[0:n]
}

func resizeint(x []int, n int) []int {
	if cap(x) < n {
		x = make([]int, n)
	}
	return x[0:n]
}

func resizestring(x []string, n int) []string {
	if cap(x) < n {
		x = make([]string, n)
	}
	return x[0:n]
}

func ilen(x interface{}) int {
	switch x := x.(type) {
	case []float64:
		return len(x)
	case []float32:
		return len(x)
	case []uint64:
		return len(x)
	case []uint32:
		return len(x)
	case []uint16:
		return len(x)
	case []uint8:
		return len(x)
	case []int64:
		return len(x)
	case []int32:
		return len(x)
	case []int16:
		return len(x)
	case []int8:
		return len(x)
	case []int:
		return len(x)
	case []string:
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

			case []float64:
				z[j] = x[0:0]

			case []float32:
				z[j] = x[0:0]

			case []uint64:
				z[j] = x[0:0]

			case []uint32:
				z[j] = x[0:0]

			case []uint16:
				z[j] = x[0:0]

			case []uint8:
				z[j] = x[0:0]

			case []int64:
				z[j] = x[0:0]

			case []int32:
				z[j] = x[0:0]

			case []int16:
				z[j] = x[0:0]

			case []int8:
				z[j] = x[0:0]

			case []int:
				z[j] = x[0:0]

			case []string:
				z[j] = x[0:0]

			default:
				msg := fmt.Sprintf("unknown type %T", x)
				panic(msg)
			}
		}
	}
}

// GetCol returns a copy of the data for one variable.  The data
// are returned as a slice.
func GetCol(da Dstream, na string) interface{} {

	vn := da.Names()
	for j, x := range vn {
		if na == x {
			return GetColPos(da, j)
		}
	}
	panic(fmt.Sprintf("Variable '%s' not found", na))
	return nil
}

// GetColPos returns a copy of the data for one variable.
// The data are returned as a slice.
func GetColPos(da Dstream, j int) interface{} {

	da.Reset()
	da.Next()
	v := da.GetPos(j)

	switch v := v.(type) {
	case []float64:
		var x []float64
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]float64)
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
	case []uint64:
		var x []uint64
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]uint64)
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
	case []uint16:
		var x []uint16
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]uint16)
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
	case []int64:
		var x []int64
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]int64)
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
	case []int16:
		var x []int16
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]int16)
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
	case []int:
		var x []int
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]int)
			x = append(x, y...)
		}
		return x
	case []string:
		var x []string
		x = append(x, v...)
		for da.Next() {
			y := da.GetPos(j).([]string)
			x = append(x, y...)
		}
		return x
	}

	panic("GetColPos: unknown type")
	return nil
}
