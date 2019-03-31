// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
	"os"
	"time"
)

// EqualReport compares two Dstream values.  If they are not equal,
// further information is written to the standard error stream.  Equality
// here implies that the data values, types, order, and chunk
// boundaries are all identical.
func EqualReport(x, y Dstream, report bool) bool {

	x.Reset()
	y.Reset()

	// Check variable names
	if !aequalString(x.Names(), y.Names()) {
		if report {
			msg := fmt.Sprintf("Names are not equal:\nx: %v\ny: %v\n",
				x.Names(), y.Names())
			os.Stderr.WriteString(msg)
		}
		return false
	}

	if x.NumVar() != y.NumVar() {
		if report {
			msg := fmt.Sprintf("Number of variables are not equal:\nx: %d\ny: %d\n",
				x.NumVar(), y.NumVar())
			os.Stderr.WriteString(msg)
		}
		return false
	}

	for chunk := 0; x.Next(); chunk++ {
		if !y.Next() {
			if report {
				msg := fmt.Sprintf("unequal numbers of chunks (y has fewer chunks than x)\n")
				print(msg)
			}
			return false
		}
		for j := 0; j < x.NumVar(); j++ {
			switch v := x.GetPos(j).(type) {

			case []string:
				u, ok := y.GetPos(j).([]string)
				if !ok || !aequalString(v, u) {
					if report {
						fmt.Printf("Chunk %d, %s\n", chunk, x.Names()[j])
						fmt.Printf("  Unequal floats:\n    (1) %v\n    (2) %v\n", v, u)
					}
					return false
				}

			case []time.Time:
				u, ok := y.GetPos(j).([]time.Time)
				if !ok || !aequalTime(v, u) {
					if report {
						fmt.Printf("Chunk %d, %s\n", chunk, x.Names()[j])
						fmt.Printf("  Unequal floats:\n    (1) %v\n    (2) %v\n", v, u)
					}
					return false
				}

			case []uint8:
				u, ok := y.GetPos(j).([]uint8)
				if !ok || !aequalUint8(v, u) {
					if report {
						fmt.Printf("Chunk %d, %s\n", chunk, x.Names()[j])
						fmt.Printf("  Unequal floats:\n    (1) %v\n    (2) %v\n", v, u)
					}
					return false
				}

			case []uint16:
				u, ok := y.GetPos(j).([]uint16)
				if !ok || !aequalUint16(v, u) {
					if report {
						fmt.Printf("Chunk %d, %s\n", chunk, x.Names()[j])
						fmt.Printf("  Unequal floats:\n    (1) %v\n    (2) %v\n", v, u)
					}
					return false
				}

			case []uint32:
				u, ok := y.GetPos(j).([]uint32)
				if !ok || !aequalUint32(v, u) {
					if report {
						fmt.Printf("Chunk %d, %s\n", chunk, x.Names()[j])
						fmt.Printf("  Unequal floats:\n    (1) %v\n    (2) %v\n", v, u)
					}
					return false
				}

			case []uint64:
				u, ok := y.GetPos(j).([]uint64)
				if !ok || !aequalUint64(v, u) {
					if report {
						fmt.Printf("Chunk %d, %s\n", chunk, x.Names()[j])
						fmt.Printf("  Unequal floats:\n    (1) %v\n    (2) %v\n", v, u)
					}
					return false
				}

			case []int8:
				u, ok := y.GetPos(j).([]int8)
				if !ok || !aequalInt8(v, u) {
					if report {
						fmt.Printf("Chunk %d, %s\n", chunk, x.Names()[j])
						fmt.Printf("  Unequal floats:\n    (1) %v\n    (2) %v\n", v, u)
					}
					return false
				}

			case []int16:
				u, ok := y.GetPos(j).([]int16)
				if !ok || !aequalInt16(v, u) {
					if report {
						fmt.Printf("Chunk %d, %s\n", chunk, x.Names()[j])
						fmt.Printf("  Unequal floats:\n    (1) %v\n    (2) %v\n", v, u)
					}
					return false
				}

			case []int32:
				u, ok := y.GetPos(j).([]int32)
				if !ok || !aequalInt32(v, u) {
					if report {
						fmt.Printf("Chunk %d, %s\n", chunk, x.Names()[j])
						fmt.Printf("  Unequal floats:\n    (1) %v\n    (2) %v\n", v, u)
					}
					return false
				}

			case []int64:
				u, ok := y.GetPos(j).([]int64)
				if !ok || !aequalInt64(v, u) {
					if report {
						fmt.Printf("Chunk %d, %s\n", chunk, x.Names()[j])
						fmt.Printf("  Unequal floats:\n    (1) %v\n    (2) %v\n", v, u)
					}
					return false
				}

			case []float32:
				u, ok := y.GetPos(j).([]float32)
				if !ok || !aequalFloat32(v, u) {
					if report {
						fmt.Printf("Chunk %d, %s\n", chunk, x.Names()[j])
						fmt.Printf("  Unequal floats:\n    (1) %v\n    (2) %v\n", v, u)
					}
					return false
				}

			case []float64:
				u, ok := y.GetPos(j).([]float64)
				if !ok || !aequalFloat64(v, u) {
					if report {
						fmt.Printf("Chunk %d, %s\n", chunk, x.Names()[j])
						fmt.Printf("  Unequal floats:\n    (1) %v\n    (2) %v\n", v, u)
					}
					return false
				}

			default:
				if report {
					print("mismatched types")
				}
				return false
			}
		}
	}

	if y.Next() {
		if report {
			msg := fmt.Sprintf("unequal numbers of chunks (x has fewer chunks than y)\n")
			print(msg)
		}
		return false
	}

	return true
}

func aequalString(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func aequalTime(x, y []time.Time) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func aequalUint8(x, y []uint8) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func aequalUint16(x, y []uint16) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func aequalUint32(x, y []uint32) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func aequalUint64(x, y []uint64) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func aequalInt8(x, y []int8) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func aequalInt16(x, y []int16) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func aequalInt32(x, y []int32) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func aequalInt64(x, y []int64) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func aequalFloat32(x, y []float32) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func aequalFloat64(x, y []float64) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}
