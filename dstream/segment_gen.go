// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
	"time"
)

func (sd *segmentedData) GetPos(j int) interface{} {
	var x interface{}
	var stash bool
	if ilen(sd.stash[j]) > 0 {
		x = sd.stash[j]
		stash = true
	} else {
		x = sd.bdata[j]
	}
	switch x := x.(type) {
	case []string:
		pos := sd.pos
		if stash {
			pos = len(x)
		}
		return x[0:pos]
	case []time.Time:
		pos := sd.pos
		if stash {
			pos = len(x)
		}
		return x[0:pos]
	case []uint8:
		pos := sd.pos
		if stash {
			pos = len(x)
		}
		return x[0:pos]
	case []uint16:
		pos := sd.pos
		if stash {
			pos = len(x)
		}
		return x[0:pos]
	case []uint32:
		pos := sd.pos
		if stash {
			pos = len(x)
		}
		return x[0:pos]
	case []uint64:
		pos := sd.pos
		if stash {
			pos = len(x)
		}
		return x[0:pos]
	case []int8:
		pos := sd.pos
		if stash {
			pos = len(x)
		}
		return x[0:pos]
	case []int16:
		pos := sd.pos
		if stash {
			pos = len(x)
		}
		return x[0:pos]
	case []int32:
		pos := sd.pos
		if stash {
			pos = len(x)
		}
		return x[0:pos]
	case []int64:
		pos := sd.pos
		if stash {
			pos = len(x)
		}
		return x[0:pos]
	case []float32:
		pos := sd.pos
		if stash {
			pos = len(x)
		}
		return x[0:pos]
	case []float64:
		pos := sd.pos
		if stash {
			pos = len(x)
		}
		return x[0:pos]
	default:
		msg := fmt.Sprintf("Segment: unknown type %T\n", x)
		panic(msg)
	}
}

// fixstash appends the matching initial segment of bdata to the
// end of the stash
func (sd *segmentedData) fixstash() bool {
	pos, fd := sd.findSegmentStash()
	if pos == 0 {
		return true
	}
	for j := 0; j < sd.source.NumVar(); j++ {
		x := sd.bdata[j]
		switch x := x.(type) {
		case []string:
			z := sd.stash[j].([]string)
			sd.stash[j] = append(z, x[0:pos]...)
			sd.bdata[j] = x[pos:len(x)]
		case []time.Time:
			z := sd.stash[j].([]time.Time)
			sd.stash[j] = append(z, x[0:pos]...)
			sd.bdata[j] = x[pos:len(x)]
		case []uint8:
			z := sd.stash[j].([]uint8)
			sd.stash[j] = append(z, x[0:pos]...)
			sd.bdata[j] = x[pos:len(x)]
		case []uint16:
			z := sd.stash[j].([]uint16)
			sd.stash[j] = append(z, x[0:pos]...)
			sd.bdata[j] = x[pos:len(x)]
		case []uint32:
			z := sd.stash[j].([]uint32)
			sd.stash[j] = append(z, x[0:pos]...)
			sd.bdata[j] = x[pos:len(x)]
		case []uint64:
			z := sd.stash[j].([]uint64)
			sd.stash[j] = append(z, x[0:pos]...)
			sd.bdata[j] = x[pos:len(x)]
		case []int8:
			z := sd.stash[j].([]int8)
			sd.stash[j] = append(z, x[0:pos]...)
			sd.bdata[j] = x[pos:len(x)]
		case []int16:
			z := sd.stash[j].([]int16)
			sd.stash[j] = append(z, x[0:pos]...)
			sd.bdata[j] = x[pos:len(x)]
		case []int32:
			z := sd.stash[j].([]int32)
			sd.stash[j] = append(z, x[0:pos]...)
			sd.bdata[j] = x[pos:len(x)]
		case []int64:
			z := sd.stash[j].([]int64)
			sd.stash[j] = append(z, x[0:pos]...)
			sd.bdata[j] = x[pos:len(x)]
		case []float32:
			z := sd.stash[j].([]float32)
			sd.stash[j] = append(z, x[0:pos]...)
			sd.bdata[j] = x[pos:len(x)]
		case []float64:
			z := sd.stash[j].([]float64)
			sd.stash[j] = append(z, x[0:pos]...)
			sd.bdata[j] = x[pos:len(x)]
		default:
			msg := fmt.Sprintf("Segment: unknown type %T\n", x)
			panic(msg)
		}
	}
	return fd
}

// setstash copies bdata into stash, replacing whatever was there.
func (sd *segmentedData) setstash() {
	sd.stash = make([]interface{}, sd.source.NumVar())
	for j := 0; j < sd.source.NumVar(); j++ {
		x := sd.bdata[j]
		switch x := x.(type) {
		case []string:
			var z []string
			if sd.stash[j] != nil {
				z = sd.stash[j].([]string)
			}
			z = resizeString(z, len(x))
			copy(z, x)
			sd.stash[j] = z
		case []time.Time:
			var z []time.Time
			if sd.stash[j] != nil {
				z = sd.stash[j].([]time.Time)
			}
			z = resizeTime(z, len(x))
			copy(z, x)
			sd.stash[j] = z
		case []uint8:
			var z []uint8
			if sd.stash[j] != nil {
				z = sd.stash[j].([]uint8)
			}
			z = resizeUint8(z, len(x))
			copy(z, x)
			sd.stash[j] = z
		case []uint16:
			var z []uint16
			if sd.stash[j] != nil {
				z = sd.stash[j].([]uint16)
			}
			z = resizeUint16(z, len(x))
			copy(z, x)
			sd.stash[j] = z
		case []uint32:
			var z []uint32
			if sd.stash[j] != nil {
				z = sd.stash[j].([]uint32)
			}
			z = resizeUint32(z, len(x))
			copy(z, x)
			sd.stash[j] = z
		case []uint64:
			var z []uint64
			if sd.stash[j] != nil {
				z = sd.stash[j].([]uint64)
			}
			z = resizeUint64(z, len(x))
			copy(z, x)
			sd.stash[j] = z
		case []int8:
			var z []int8
			if sd.stash[j] != nil {
				z = sd.stash[j].([]int8)
			}
			z = resizeInt8(z, len(x))
			copy(z, x)
			sd.stash[j] = z
		case []int16:
			var z []int16
			if sd.stash[j] != nil {
				z = sd.stash[j].([]int16)
			}
			z = resizeInt16(z, len(x))
			copy(z, x)
			sd.stash[j] = z
		case []int32:
			var z []int32
			if sd.stash[j] != nil {
				z = sd.stash[j].([]int32)
			}
			z = resizeInt32(z, len(x))
			copy(z, x)
			sd.stash[j] = z
		case []int64:
			var z []int64
			if sd.stash[j] != nil {
				z = sd.stash[j].([]int64)
			}
			z = resizeInt64(z, len(x))
			copy(z, x)
			sd.stash[j] = z
		case []float32:
			var z []float32
			if sd.stash[j] != nil {
				z = sd.stash[j].([]float32)
			}
			z = resizeFloat32(z, len(x))
			copy(z, x)
			sd.stash[j] = z
		case []float64:
			var z []float64
			if sd.stash[j] != nil {
				z = sd.stash[j].([]float64)
			}
			z = resizeFloat64(z, len(x))
			copy(z, x)
			sd.stash[j] = z
		default:
			msg := fmt.Sprintf("Segment: unknown type %T\n", x)
			panic(msg)
		}
	}
}

// leftsliceb reslices every element of bdata from position
// pos to the end of the slice.
func (sd *segmentedData) leftsliceb(pos int) {
	for j := 0; j < sd.source.NumVar(); j++ {
		x := sd.bdata[j]
		switch x := x.(type) {
		case []string:
			sd.bdata[j] = x[pos:len(x)]
		case []time.Time:
			sd.bdata[j] = x[pos:len(x)]
		case []uint8:
			sd.bdata[j] = x[pos:len(x)]
		case []uint16:
			sd.bdata[j] = x[pos:len(x)]
		case []uint32:
			sd.bdata[j] = x[pos:len(x)]
		case []uint64:
			sd.bdata[j] = x[pos:len(x)]
		case []int8:
			sd.bdata[j] = x[pos:len(x)]
		case []int16:
			sd.bdata[j] = x[pos:len(x)]
		case []int32:
			sd.bdata[j] = x[pos:len(x)]
		case []int64:
			sd.bdata[j] = x[pos:len(x)]
		case []float32:
			sd.bdata[j] = x[pos:len(x)]
		case []float64:
			sd.bdata[j] = x[pos:len(x)]
		default:
			msg := fmt.Sprintf("Segment: unknown type %T\n", x)
			panic(msg)
		}
	}
}

// findSegment finds the next segment boundary after start in the
// current backing slice.  If there is no boundary, returns -1.
func (sd *segmentedData) findSegment(start int) int {
	pos := -1
	for _, j := range sd.vpos {
		x := sd.bdata[j]
		switch x := x.(type) {
		case []string:
			for i := start + 1; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != x[i-1] {
					pos = i
					break
				}
			}
		case []time.Time:
			for i := start + 1; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != x[i-1] {
					pos = i
					break
				}
			}
		case []uint8:
			for i := start + 1; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != x[i-1] {
					pos = i
					break
				}
			}
		case []uint16:
			for i := start + 1; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != x[i-1] {
					pos = i
					break
				}
			}
		case []uint32:
			for i := start + 1; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != x[i-1] {
					pos = i
					break
				}
			}
		case []uint64:
			for i := start + 1; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != x[i-1] {
					pos = i
					break
				}
			}
		case []int8:
			for i := start + 1; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != x[i-1] {
					pos = i
					break
				}
			}
		case []int16:
			for i := start + 1; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != x[i-1] {
					pos = i
					break
				}
			}
		case []int32:
			for i := start + 1; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != x[i-1] {
					pos = i
					break
				}
			}
		case []int64:
			for i := start + 1; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != x[i-1] {
					pos = i
					break
				}
			}
		case []float32:
			for i := start + 1; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != x[i-1] {
					pos = i
					break
				}
			}
		case []float64:
			for i := start + 1; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != x[i-1] {
					pos = i
					break
				}
			}
		case nil:
			return -1
		default:
			msg := fmt.Sprintf("Segment: unknown type %T\n", x)
			panic(msg)
		}
	}

	return pos
}

// findSegmentStash finds the first segment boundary in bdata, viewing bstash
// as a continuation of stash.
func (sd *segmentedData) findSegmentStash() (int, bool) {
	pos := -1
	var m int
	for _, j := range sd.vpos {
		x := sd.bdata[j]
		switch x := x.(type) {
		case []string:
			m = len(x)
			y := sd.stash[j].([]string)
			v := y[len(y)-1]
			for i := 0; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != v {
					pos = i
					break
				}
			}
		case []time.Time:
			m = len(x)
			y := sd.stash[j].([]time.Time)
			v := y[len(y)-1]
			for i := 0; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != v {
					pos = i
					break
				}
			}
		case []uint8:
			m = len(x)
			y := sd.stash[j].([]uint8)
			v := y[len(y)-1]
			for i := 0; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != v {
					pos = i
					break
				}
			}
		case []uint16:
			m = len(x)
			y := sd.stash[j].([]uint16)
			v := y[len(y)-1]
			for i := 0; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != v {
					pos = i
					break
				}
			}
		case []uint32:
			m = len(x)
			y := sd.stash[j].([]uint32)
			v := y[len(y)-1]
			for i := 0; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != v {
					pos = i
					break
				}
			}
		case []uint64:
			m = len(x)
			y := sd.stash[j].([]uint64)
			v := y[len(y)-1]
			for i := 0; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != v {
					pos = i
					break
				}
			}
		case []int8:
			m = len(x)
			y := sd.stash[j].([]int8)
			v := y[len(y)-1]
			for i := 0; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != v {
					pos = i
					break
				}
			}
		case []int16:
			m = len(x)
			y := sd.stash[j].([]int16)
			v := y[len(y)-1]
			for i := 0; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != v {
					pos = i
					break
				}
			}
		case []int32:
			m = len(x)
			y := sd.stash[j].([]int32)
			v := y[len(y)-1]
			for i := 0; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != v {
					pos = i
					break
				}
			}
		case []int64:
			m = len(x)
			y := sd.stash[j].([]int64)
			v := y[len(y)-1]
			for i := 0; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != v {
					pos = i
					break
				}
			}
		case []float32:
			m = len(x)
			y := sd.stash[j].([]float32)
			v := y[len(y)-1]
			for i := 0; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != v {
					pos = i
					break
				}
			}
		case []float64:
			m = len(x)
			y := sd.stash[j].([]float64)
			v := y[len(y)-1]
			for i := 0; i < len(x); i++ {
				if pos != -1 && i >= pos {
					break
				}
				if x[i] != v {
					pos = i
					break
				}
			}
		default:
			msg := fmt.Sprintf("Segment: unknown type %T", x)
			panic(msg)
		}
	}

	if pos != -1 {
		return pos, true
	} else {
		return m, false
	}
}
