// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
)

func (c *convertType) GetPos(j int) interface{} {

	if j != c.vpos {
		return c.source.GetPos(j)
	}

	if c.conv[j] {
		// The result is already available so just return it.
		//print("HI\n")
		return c.bdata[j]
	}

	// Initialize the backing array on the first call
	to := c.bdata[j]
	if to == nil {
		switch c.dtype {
		case "float64":
			to = make([]float64, 0, 100)
		case "float32":
			to = make([]float32, 0, 100)
		case "uint64":
			to = make([]uint64, 0, 100)
		case "uint32":
			to = make([]uint32, 0, 100)
		case "uint16":
			to = make([]uint16, 0, 100)
		case "uint8":
			to = make([]uint8, 0, 100)
		case "int64":
			to = make([]int64, 0, 100)
		case "int32":
			to = make([]int32, 0, 100)
		case "int16":
			to = make([]int16, 0, 100)
		case "int8":
			to = make([]int8, 0, 100)
		case "int":
			to = make([]int, 0, 100)
		default:
			msg := fmt.Sprintf("Convert: unknown type %s\n", c.dtype)
			panic(msg)
		}
	}

	// Need this to do nested switches.

	from := c.source.GetPos(j)
	switch to := to.(type) {
	case []float64:
		switch from := from.(type) {
		case []float64:
			c.bdata[j] = from
		case []float32:
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[j] = to
		case []uint64:
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[j] = to
		case []uint32:
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[j] = to
		case []uint16:
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[j] = to
		case []uint8:
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[j] = to
		case []int64:
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[j] = to
		case []int32:
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[j] = to
		case []int16:
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[j] = to
		case []int8:
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[j] = to
		case []int:
			for _, x := range from {
				to = append(to, float64(x))
			}
			c.bdata[j] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []float32:
		switch from := from.(type) {
		case []float64:
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[j] = to
		case []float32:
			c.bdata[j] = from
		case []uint64:
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[j] = to
		case []uint32:
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[j] = to
		case []uint16:
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[j] = to
		case []uint8:
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[j] = to
		case []int64:
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[j] = to
		case []int32:
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[j] = to
		case []int16:
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[j] = to
		case []int8:
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[j] = to
		case []int:
			for _, x := range from {
				to = append(to, float32(x))
			}
			c.bdata[j] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []uint64:
		switch from := from.(type) {
		case []float64:
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[j] = to
		case []float32:
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[j] = to
		case []uint64:
			c.bdata[j] = from
		case []uint32:
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[j] = to
		case []uint16:
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[j] = to
		case []uint8:
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[j] = to
		case []int64:
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[j] = to
		case []int32:
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[j] = to
		case []int16:
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[j] = to
		case []int8:
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[j] = to
		case []int:
			for _, x := range from {
				to = append(to, uint64(x))
			}
			c.bdata[j] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []uint32:
		switch from := from.(type) {
		case []float64:
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[j] = to
		case []float32:
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[j] = to
		case []uint64:
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[j] = to
		case []uint32:
			c.bdata[j] = from
		case []uint16:
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[j] = to
		case []uint8:
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[j] = to
		case []int64:
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[j] = to
		case []int32:
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[j] = to
		case []int16:
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[j] = to
		case []int8:
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[j] = to
		case []int:
			for _, x := range from {
				to = append(to, uint32(x))
			}
			c.bdata[j] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []uint16:
		switch from := from.(type) {
		case []float64:
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[j] = to
		case []float32:
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[j] = to
		case []uint64:
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[j] = to
		case []uint32:
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[j] = to
		case []uint16:
			c.bdata[j] = from
		case []uint8:
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[j] = to
		case []int64:
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[j] = to
		case []int32:
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[j] = to
		case []int16:
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[j] = to
		case []int8:
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[j] = to
		case []int:
			for _, x := range from {
				to = append(to, uint16(x))
			}
			c.bdata[j] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []uint8:
		switch from := from.(type) {
		case []float64:
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[j] = to
		case []float32:
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[j] = to
		case []uint64:
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[j] = to
		case []uint32:
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[j] = to
		case []uint16:
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[j] = to
		case []uint8:
			c.bdata[j] = from
		case []int64:
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[j] = to
		case []int32:
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[j] = to
		case []int16:
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[j] = to
		case []int8:
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[j] = to
		case []int:
			for _, x := range from {
				to = append(to, uint8(x))
			}
			c.bdata[j] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []int64:
		switch from := from.(type) {
		case []float64:
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[j] = to
		case []float32:
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[j] = to
		case []uint64:
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[j] = to
		case []uint32:
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[j] = to
		case []uint16:
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[j] = to
		case []uint8:
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[j] = to
		case []int64:
			c.bdata[j] = from
		case []int32:
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[j] = to
		case []int16:
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[j] = to
		case []int8:
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[j] = to
		case []int:
			for _, x := range from {
				to = append(to, int64(x))
			}
			c.bdata[j] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []int32:
		switch from := from.(type) {
		case []float64:
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[j] = to
		case []float32:
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[j] = to
		case []uint64:
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[j] = to
		case []uint32:
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[j] = to
		case []uint16:
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[j] = to
		case []uint8:
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[j] = to
		case []int64:
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[j] = to
		case []int32:
			c.bdata[j] = from
		case []int16:
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[j] = to
		case []int8:
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[j] = to
		case []int:
			for _, x := range from {
				to = append(to, int32(x))
			}
			c.bdata[j] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []int16:
		switch from := from.(type) {
		case []float64:
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[j] = to
		case []float32:
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[j] = to
		case []uint64:
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[j] = to
		case []uint32:
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[j] = to
		case []uint16:
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[j] = to
		case []uint8:
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[j] = to
		case []int64:
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[j] = to
		case []int32:
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[j] = to
		case []int16:
			c.bdata[j] = from
		case []int8:
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[j] = to
		case []int:
			for _, x := range from {
				to = append(to, int16(x))
			}
			c.bdata[j] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []int8:
		switch from := from.(type) {
		case []float64:
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[j] = to
		case []float32:
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[j] = to
		case []uint64:
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[j] = to
		case []uint32:
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[j] = to
		case []uint16:
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[j] = to
		case []uint8:
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[j] = to
		case []int64:
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[j] = to
		case []int32:
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[j] = to
		case []int16:
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[j] = to
		case []int8:
			c.bdata[j] = from
		case []int:
			for _, x := range from {
				to = append(to, int8(x))
			}
			c.bdata[j] = to
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	case []int:
		switch from := from.(type) {
		case []float64:
			for _, x := range from {
				to = append(to, int(x))
			}
			c.bdata[j] = to
		case []float32:
			for _, x := range from {
				to = append(to, int(x))
			}
			c.bdata[j] = to
		case []uint64:
			for _, x := range from {
				to = append(to, int(x))
			}
			c.bdata[j] = to
		case []uint32:
			for _, x := range from {
				to = append(to, int(x))
			}
			c.bdata[j] = to
		case []uint16:
			for _, x := range from {
				to = append(to, int(x))
			}
			c.bdata[j] = to
		case []uint8:
			for _, x := range from {
				to = append(to, int(x))
			}
			c.bdata[j] = to
		case []int64:
			for _, x := range from {
				to = append(to, int(x))
			}
			c.bdata[j] = to
		case []int32:
			for _, x := range from {
				to = append(to, int(x))
			}
			c.bdata[j] = to
		case []int16:
			for _, x := range from {
				to = append(to, int(x))
			}
			c.bdata[j] = to
		case []int8:
			for _, x := range from {
				to = append(to, int(x))
			}
			c.bdata[j] = to
		case []int:
			c.bdata[j] = from
		default:
			msg := fmt.Sprintf("Convert: unknown origin type %T\n", from)
			panic(msg)
		}
	default:
		msg := fmt.Sprintf("Convert: unkown destination type %T\n", to)
		panic(msg)
	}

	c.conv[j] = true
	return c.bdata[j]
}
