// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
	"time"
)

// Dechunk combines all chunks into a single chunk.
func Dechunk(source Dstream) Dstream {

	// Use the first chunk to get started
	if !source.Next() {
		msg := "Can't read from stream"
		panic(msg)
	}

	names := source.Names()
	nvar := source.NumVar()
	data := make([][]interface{}, nvar)
	for j := 0; j < nvar; j++ {
		// The result has only one chunk
		data[j] = make([]interface{}, 1)
	}

	for {
		for j := 0; j < nvar; j++ {
			x := source.GetPos(j)
			switch x := x.(type) {
			case []string:
				var z []string
				if data[j][0] != nil {
					z = data[j][0].([]string)
				}
				data[j][0] = append(z, x...)
			case []time.Time:
				var z []time.Time
				if data[j][0] != nil {
					z = data[j][0].([]time.Time)
				}
				data[j][0] = append(z, x...)
			case []uint8:
				var z []uint8
				if data[j][0] != nil {
					z = data[j][0].([]uint8)
				}
				data[j][0] = append(z, x...)
			case []uint16:
				var z []uint16
				if data[j][0] != nil {
					z = data[j][0].([]uint16)
				}
				data[j][0] = append(z, x...)
			case []uint32:
				var z []uint32
				if data[j][0] != nil {
					z = data[j][0].([]uint32)
				}
				data[j][0] = append(z, x...)
			case []uint64:
				var z []uint64
				if data[j][0] != nil {
					z = data[j][0].([]uint64)
				}
				data[j][0] = append(z, x...)
			case []int8:
				var z []int8
				if data[j][0] != nil {
					z = data[j][0].([]int8)
				}
				data[j][0] = append(z, x...)
			case []int16:
				var z []int16
				if data[j][0] != nil {
					z = data[j][0].([]int16)
				}
				data[j][0] = append(z, x...)
			case []int32:
				var z []int32
				if data[j][0] != nil {
					z = data[j][0].([]int32)
				}
				data[j][0] = append(z, x...)
			case []int64:
				var z []int64
				if data[j][0] != nil {
					z = data[j][0].([]int64)
				}
				data[j][0] = append(z, x...)
			case []float32:
				var z []float32
				if data[j][0] != nil {
					z = data[j][0].([]float32)
				}
				data[j][0] = append(z, x...)
			case []float64:
				var z []float64
				if data[j][0] != nil {
					z = data[j][0].([]float64)
				}
				data[j][0] = append(z, x...)
			default:
				panic(fmt.Sprintf("Type %T is not known\n", x))
			}
		}
		if !source.Next() {
			break
		}
	}

	return NewFromArrays(data, names)
}
