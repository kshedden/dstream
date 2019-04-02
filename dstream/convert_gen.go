// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
)

func (c *convert) Next() bool {

	if !c.source.Next() {
		return false
	}

	// The non-converted variables are just pointer assignments.
	for j := 0; j < c.source.NumVar(); j++ {
		if j != c.vpos {
			c.bdata[j] = c.source.GetPos(j)
		}
	}

	// Initialize the backing array if needed
	to := c.bdata[c.vpos]
	if to == nil {
		switch c.dtype {
		case Uint8:
			to = make([]uint8, 0, 100)
		case Uint16:
			to = make([]uint16, 0, 100)
		case Uint32:
			to = make([]uint32, 0, 100)
		case Uint64:
			to = make([]uint64, 0, 100)
		case Int8:
			to = make([]int8, 0, 100)
		case Int16:
			to = make([]int16, 0, 100)
		case Int32:
			to = make([]int32, 0, 100)
		case Int64:
			to = make([]int64, 0, 100)
		case Float32:
			to = make([]float32, 0, 100)
		case Float64:
			to = make([]float64, 0, 100)
		default:
			msg := fmt.Sprintf("Convert: unknown type %v\n", c.dtype)
			panic(msg)
		}
	}

	// Need this to do nested switches.

	// Unconverted data
	from := c.source.GetPos(c.vpos)

	switch to := to.(type) {
	case []uint8:
		switch from := from.(type) {
		case []uint8:
			// Same types, nothing to do
			c.bdata[c.vpos] = from
		case []uint16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[c.vpos] = to
		case []uint32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[c.vpos] = to
		case []uint64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[c.vpos] = to
		case []int8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[c.vpos] = to
		case []int16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[c.vpos] = to
		case []int32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[c.vpos] = to
		case []int64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[c.vpos] = to
		case []float32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[c.vpos] = to
		case []float64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[c.vpos] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []uint16:
		switch from := from.(type) {
		case []uint8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[c.vpos] = to
		case []uint16:
			// Same types, nothing to do
			c.bdata[c.vpos] = from
		case []uint32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[c.vpos] = to
		case []uint64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[c.vpos] = to
		case []int8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[c.vpos] = to
		case []int16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[c.vpos] = to
		case []int32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[c.vpos] = to
		case []int64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[c.vpos] = to
		case []float32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[c.vpos] = to
		case []float64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[c.vpos] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []uint32:
		switch from := from.(type) {
		case []uint8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[c.vpos] = to
		case []uint16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[c.vpos] = to
		case []uint32:
			// Same types, nothing to do
			c.bdata[c.vpos] = from
		case []uint64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[c.vpos] = to
		case []int8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[c.vpos] = to
		case []int16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[c.vpos] = to
		case []int32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[c.vpos] = to
		case []int64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[c.vpos] = to
		case []float32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[c.vpos] = to
		case []float64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[c.vpos] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []uint64:
		switch from := from.(type) {
		case []uint8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[c.vpos] = to
		case []uint16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[c.vpos] = to
		case []uint32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[c.vpos] = to
		case []uint64:
			// Same types, nothing to do
			c.bdata[c.vpos] = from
		case []int8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[c.vpos] = to
		case []int16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[c.vpos] = to
		case []int32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[c.vpos] = to
		case []int64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[c.vpos] = to
		case []float32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[c.vpos] = to
		case []float64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[c.vpos] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []int8:
		switch from := from.(type) {
		case []uint8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[c.vpos] = to
		case []uint16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[c.vpos] = to
		case []uint32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[c.vpos] = to
		case []uint64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[c.vpos] = to
		case []int8:
			// Same types, nothing to do
			c.bdata[c.vpos] = from
		case []int16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[c.vpos] = to
		case []int32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[c.vpos] = to
		case []int64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[c.vpos] = to
		case []float32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[c.vpos] = to
		case []float64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[c.vpos] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []int16:
		switch from := from.(type) {
		case []uint8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[c.vpos] = to
		case []uint16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[c.vpos] = to
		case []uint32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[c.vpos] = to
		case []uint64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[c.vpos] = to
		case []int8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[c.vpos] = to
		case []int16:
			// Same types, nothing to do
			c.bdata[c.vpos] = from
		case []int32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[c.vpos] = to
		case []int64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[c.vpos] = to
		case []float32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[c.vpos] = to
		case []float64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[c.vpos] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []int32:
		switch from := from.(type) {
		case []uint8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[c.vpos] = to
		case []uint16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[c.vpos] = to
		case []uint32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[c.vpos] = to
		case []uint64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[c.vpos] = to
		case []int8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[c.vpos] = to
		case []int16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[c.vpos] = to
		case []int32:
			// Same types, nothing to do
			c.bdata[c.vpos] = from
		case []int64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[c.vpos] = to
		case []float32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[c.vpos] = to
		case []float64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[c.vpos] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []int64:
		switch from := from.(type) {
		case []uint8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[c.vpos] = to
		case []uint16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[c.vpos] = to
		case []uint32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[c.vpos] = to
		case []uint64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[c.vpos] = to
		case []int8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[c.vpos] = to
		case []int16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[c.vpos] = to
		case []int32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[c.vpos] = to
		case []int64:
			// Same types, nothing to do
			c.bdata[c.vpos] = from
		case []float32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[c.vpos] = to
		case []float64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[c.vpos] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []float32:
		switch from := from.(type) {
		case []uint8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[c.vpos] = to
		case []uint16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[c.vpos] = to
		case []uint32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[c.vpos] = to
		case []uint64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[c.vpos] = to
		case []int8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[c.vpos] = to
		case []int16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[c.vpos] = to
		case []int32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[c.vpos] = to
		case []int64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[c.vpos] = to
		case []float32:
			// Same types, nothing to do
			c.bdata[c.vpos] = from
		case []float64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[c.vpos] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []float64:
		switch from := from.(type) {
		case []uint8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[c.vpos] = to
		case []uint16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[c.vpos] = to
		case []uint32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[c.vpos] = to
		case []uint64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[c.vpos] = to
		case []int8:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[c.vpos] = to
		case []int16:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[c.vpos] = to
		case []int32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[c.vpos] = to
		case []int64:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[c.vpos] = to
		case []float32:
			to = to[0:0]
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[c.vpos] = to
		case []float64:
			// Same types, nothing to do
			c.bdata[c.vpos] = from
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	default:
		msg := fmt.Sprintf("Convert: unkown destination type %T\n", to)
		panic(msg)
	}

	return true
}
