// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"time"
)

type CSVTypeConf struct {

	// Names of the variables
	String  []string
	Time    []string
	Uint8   []string
	Uint16  []string
	Uint32  []string
	Uint64  []string
	Int8    []string
	Int16   []string
	Int32   []string
	Int64   []string
	Float32 []string
	Float64 []string

	// Positions of the variables in the CSV file
	StringPos  []int
	TimePos    []int
	Uint8Pos   []int
	Uint16Pos  []int
	Uint32Pos  []int
	Uint64Pos  []int
	Int8Pos    []int
	Int16Pos   []int
	Int32Pos   []int
	Int64Pos   []int
	Float32Pos []int
	Float64Pos []int

	// Function used to parse time strings
	ParseTime func(string) time.Time
}

func (tc *CSVTypeConf) hasValidPositions() bool {
	if len(tc.String) != len(tc.StringPos) {
		return false
	}
	if len(tc.Time) != len(tc.TimePos) {
		return false
	}
	if len(tc.Uint8) != len(tc.Uint8Pos) {
		return false
	}
	if len(tc.Uint16) != len(tc.Uint16Pos) {
		return false
	}
	if len(tc.Uint32) != len(tc.Uint32Pos) {
		return false
	}
	if len(tc.Uint64) != len(tc.Uint64Pos) {
		return false
	}
	if len(tc.Int8) != len(tc.Int8Pos) {
		return false
	}
	if len(tc.Int16) != len(tc.Int16Pos) {
		return false
	}
	if len(tc.Int32) != len(tc.Int32Pos) {
		return false
	}
	if len(tc.Int64) != len(tc.Int64Pos) {
		return false
	}
	if len(tc.Float32) != len(tc.Float32Pos) {
		return false
	}
	if len(tc.Float64) != len(tc.Float64Pos) {
		return false
	}

	return true
}

func (cs *CSVReader) setNames() {

	tc := cs.typeConf

	cs.names = cs.names[0:0]
	cs.names = append(cs.names, tc.String...)
	cs.names = append(cs.names, tc.Time...)
	cs.names = append(cs.names, tc.Uint8...)
	cs.names = append(cs.names, tc.Uint16...)
	cs.names = append(cs.names, tc.Uint32...)
	cs.names = append(cs.names, tc.Uint64...)
	cs.names = append(cs.names, tc.Int8...)
	cs.names = append(cs.names, tc.Int16...)
	cs.names = append(cs.names, tc.Int32...)
	cs.names = append(cs.names, tc.Int64...)
	cs.names = append(cs.names, tc.Float32...)
	cs.names = append(cs.names, tc.Float64...)
}

// Next advances to the next chunk.
func (cs *CSVReader) Next() bool {

	if cs.done {
		return false
	}

	if cs.limitchunk > 0 && cs.limitchunk <= cs.chunknum {
		cs.done = true
		return false
	}

	cs.chunknum++

	tc := cs.typeConf
	truncate(cs.bdata)

	for j := 0; j < cs.chunkSize; j++ {

		// Try to read a row, return false if done.
		var rec []string
		var err error
		if cs.firstrow != nil {
			rec = cs.firstrow
			cs.firstrow = nil
		} else {
			rec, err = cs.csvrdr.Read()
			if err == io.EOF {
				cs.done = true
				return ilen(cs.bdata[0]) > 0
			} else if err != nil {
				if cs.skipErrors {
					os.Stderr.WriteString(fmt.Sprintf("%v\n", err))
					continue
				}
				panic(err)
			}
		}
		cs.nobs++

		i := 0
		for _, pos := range tc.StringPos {
			x := rec[pos]
			u := cs.bdata[i].([]string)
			cs.bdata[i] = append(u, string(x))
			i++
		}
		for _, pos := range tc.TimePos {
			x := tc.ParseTime(rec[pos])
			u := cs.bdata[i].([]time.Time)
			cs.bdata[i] = append(u, time.Time(x))
			i++
		}
		for _, pos := range tc.Uint8Pos {
			x, err := strconv.Atoi(rec[pos])
			if err != nil {
				panic(err)
			}
			u := cs.bdata[i].([]uint8)
			cs.bdata[i] = append(u, uint8(x))
			i++
		}
		for _, pos := range tc.Uint16Pos {
			x, err := strconv.Atoi(rec[pos])
			if err != nil {
				panic(err)
			}
			u := cs.bdata[i].([]uint16)
			cs.bdata[i] = append(u, uint16(x))
			i++
		}
		for _, pos := range tc.Uint32Pos {
			x, err := strconv.Atoi(rec[pos])
			if err != nil {
				panic(err)
			}
			u := cs.bdata[i].([]uint32)
			cs.bdata[i] = append(u, uint32(x))
			i++
		}
		for _, pos := range tc.Uint64Pos {
			x, err := strconv.Atoi(rec[pos])
			if err != nil {
				panic(err)
			}
			u := cs.bdata[i].([]uint64)
			cs.bdata[i] = append(u, uint64(x))
			i++
		}
		for _, pos := range tc.Int8Pos {
			x, err := strconv.Atoi(rec[pos])
			if err != nil {
				panic(err)
			}
			u := cs.bdata[i].([]int8)
			cs.bdata[i] = append(u, int8(x))
			i++
		}
		for _, pos := range tc.Int16Pos {
			x, err := strconv.Atoi(rec[pos])
			if err != nil {
				panic(err)
			}
			u := cs.bdata[i].([]int16)
			cs.bdata[i] = append(u, int16(x))
			i++
		}
		for _, pos := range tc.Int32Pos {
			x, err := strconv.Atoi(rec[pos])
			if err != nil {
				panic(err)
			}
			u := cs.bdata[i].([]int32)
			cs.bdata[i] = append(u, int32(x))
			i++
		}
		for _, pos := range tc.Int64Pos {
			x, err := strconv.Atoi(rec[pos])
			if err != nil {
				panic(err)
			}
			u := cs.bdata[i].([]int64)
			cs.bdata[i] = append(u, int64(x))
			i++
		}
		for _, pos := range tc.Float32Pos {
			x, err := strconv.ParseFloat(rec[pos], 64)
			if err != nil {
				x = math.NaN()
			}
			u := cs.bdata[i].([]float32)
			cs.bdata[i] = append(u, float32(x))
			i++
		}
		for _, pos := range tc.Float64Pos {
			x, err := strconv.ParseFloat(rec[pos], 64)
			if err != nil {
				x = math.NaN()
			}
			u := cs.bdata[i].([]float64)
			cs.bdata[i] = append(u, float64(x))
			i++
		}
	}

	return true
}

// SetPos determines the positions in the provided list of column names
// of all configured variables.  The configured variables are given
// in the type-specific slices named Float64, Int64, etc.
func (tc *CSVTypeConf) SetPos(h []string) {

	m := make(map[string]int)
	for k, v := range h {
		m[v] = k
	}
	tc.StringPos = tc.StringPos[0:0]
	for _, v := range tc.String {
		p, ok := m[v]
		if !ok {
			msg := fmt.Sprintf("String variable '%s' not found.\n", v)
			panic(msg)
		}
		tc.StringPos = append(tc.StringPos, p)
	}

	tc.TimePos = tc.TimePos[0:0]
	for _, v := range tc.Time {
		p, ok := m[v]
		if !ok {
			msg := fmt.Sprintf("Time variable '%s' not found.\n", v)
			panic(msg)
		}
		tc.TimePos = append(tc.TimePos, p)
	}

	tc.Uint8Pos = tc.Uint8Pos[0:0]
	for _, v := range tc.Uint8 {
		p, ok := m[v]
		if !ok {
			msg := fmt.Sprintf("Uint8 variable '%s' not found.\n", v)
			panic(msg)
		}
		tc.Uint8Pos = append(tc.Uint8Pos, p)
	}

	tc.Uint16Pos = tc.Uint16Pos[0:0]
	for _, v := range tc.Uint16 {
		p, ok := m[v]
		if !ok {
			msg := fmt.Sprintf("Uint16 variable '%s' not found.\n", v)
			panic(msg)
		}
		tc.Uint16Pos = append(tc.Uint16Pos, p)
	}

	tc.Uint32Pos = tc.Uint32Pos[0:0]
	for _, v := range tc.Uint32 {
		p, ok := m[v]
		if !ok {
			msg := fmt.Sprintf("Uint32 variable '%s' not found.\n", v)
			panic(msg)
		}
		tc.Uint32Pos = append(tc.Uint32Pos, p)
	}

	tc.Uint64Pos = tc.Uint64Pos[0:0]
	for _, v := range tc.Uint64 {
		p, ok := m[v]
		if !ok {
			msg := fmt.Sprintf("Uint64 variable '%s' not found.\n", v)
			panic(msg)
		}
		tc.Uint64Pos = append(tc.Uint64Pos, p)
	}

	tc.Int8Pos = tc.Int8Pos[0:0]
	for _, v := range tc.Int8 {
		p, ok := m[v]
		if !ok {
			msg := fmt.Sprintf("Int8 variable '%s' not found.\n", v)
			panic(msg)
		}
		tc.Int8Pos = append(tc.Int8Pos, p)
	}

	tc.Int16Pos = tc.Int16Pos[0:0]
	for _, v := range tc.Int16 {
		p, ok := m[v]
		if !ok {
			msg := fmt.Sprintf("Int16 variable '%s' not found.\n", v)
			panic(msg)
		}
		tc.Int16Pos = append(tc.Int16Pos, p)
	}

	tc.Int32Pos = tc.Int32Pos[0:0]
	for _, v := range tc.Int32 {
		p, ok := m[v]
		if !ok {
			msg := fmt.Sprintf("Int32 variable '%s' not found.\n", v)
			panic(msg)
		}
		tc.Int32Pos = append(tc.Int32Pos, p)
	}

	tc.Int64Pos = tc.Int64Pos[0:0]
	for _, v := range tc.Int64 {
		p, ok := m[v]
		if !ok {
			msg := fmt.Sprintf("Int64 variable '%s' not found.\n", v)
			panic(msg)
		}
		tc.Int64Pos = append(tc.Int64Pos, p)
	}

	tc.Float32Pos = tc.Float32Pos[0:0]
	for _, v := range tc.Float32 {
		p, ok := m[v]
		if !ok {
			msg := fmt.Sprintf("Float32 variable '%s' not found.\n", v)
			panic(msg)
		}
		tc.Float32Pos = append(tc.Float32Pos, p)
	}

	tc.Float64Pos = tc.Float64Pos[0:0]
	for _, v := range tc.Float64 {
		p, ok := m[v]
		if !ok {
			msg := fmt.Sprintf("Float64 variable '%s' not found.\n", v)
			panic(msg)
		}
		tc.Float64Pos = append(tc.Float64Pos, p)
	}

}

func (cs *CSVReader) setbdata() {

	tc := cs.typeConf

	cs.bdata = make([]interface{}, len(cs.names))

	for _, na := range tc.String {
		p := cs.namepos[na]
		cs.bdata[p] = make([]string, 0)
	}
	for _, na := range tc.Time {
		p := cs.namepos[na]
		cs.bdata[p] = make([]time.Time, 0)
	}
	for _, na := range tc.Uint8 {
		p := cs.namepos[na]
		cs.bdata[p] = make([]uint8, 0)
	}
	for _, na := range tc.Uint16 {
		p := cs.namepos[na]
		cs.bdata[p] = make([]uint16, 0)
	}
	for _, na := range tc.Uint32 {
		p := cs.namepos[na]
		cs.bdata[p] = make([]uint32, 0)
	}
	for _, na := range tc.Uint64 {
		p := cs.namepos[na]
		cs.bdata[p] = make([]uint64, 0)
	}
	for _, na := range tc.Int8 {
		p := cs.namepos[na]
		cs.bdata[p] = make([]int8, 0)
	}
	for _, na := range tc.Int16 {
		p := cs.namepos[na]
		cs.bdata[p] = make([]int16, 0)
	}
	for _, na := range tc.Int32 {
		p := cs.namepos[na]
		cs.bdata[p] = make([]int32, 0)
	}
	for _, na := range tc.Int64 {
		p := cs.namepos[na]
		cs.bdata[p] = make([]int64, 0)
	}
	for _, na := range tc.Float32 {
		p := cs.namepos[na]
		cs.bdata[p] = make([]float32, 0)
	}
	for _, na := range tc.Float64 {
		p := cs.namepos[na]
		cs.bdata[p] = make([]float64, 0)
	}
}
