// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
)

const (
	bc = 100
)

func addchunk(da Dstream, rslt *dataArrays) {
	for j := 0; j < da.NumVar(); j++ {
		v := da.GetPos(j)
		switch v := v.(type) {

		case []string:
			rslt.arrays[j] = append(rslt.arrays[j], make([]string, 0, bc))

		case []float64:
			rslt.arrays[j] = append(rslt.arrays[j], make([]float64, 0, bc))

		case []float32:
			rslt.arrays[j] = append(rslt.arrays[j], make([]float32, 0, bc))

		case []uint64:
			rslt.arrays[j] = append(rslt.arrays[j], make([]uint64, 0, bc))

		case []uint32:
			rslt.arrays[j] = append(rslt.arrays[j], make([]uint32, 0, bc))

		case []uint16:
			rslt.arrays[j] = append(rslt.arrays[j], make([]uint16, 0, bc))

		case []uint8:
			rslt.arrays[j] = append(rslt.arrays[j], make([]uint8, 0, bc))

		case []int64:
			rslt.arrays[j] = append(rslt.arrays[j], make([]int64, 0, bc))

		case []int32:
			rslt.arrays[j] = append(rslt.arrays[j], make([]int32, 0, bc))

		case []int16:
			rslt.arrays[j] = append(rslt.arrays[j], make([]int16, 0, bc))

		case []int8:
			rslt.arrays[j] = append(rslt.arrays[j], make([]int8, 0, bc))

		case []int:
			rslt.arrays[j] = append(rslt.arrays[j], make([]int, 0, bc))

		default:
			msg := fmt.Sprintf("Regroup: unknown type %T\n", v)
			panic(msg)
		}
	}
}

func doRegroup(da Dstream, varpos int) *dataArrays {

	bucket := make(map[uint64]int)

	rslt := &dataArrays{
		xform: xform{
			names: da.Names(),
		},
		arrays: make([][]interface{}, da.NumVar()),
	}

	for da.Next() {

		idv := da.GetPos(varpos).([]uint64)

		for i, id := range idv {

			b, ok := bucket[id]
			if !ok {
				b = len(bucket)
				bucket[id] = b
				addchunk(da, rslt)
			}

			for k := 0; k < da.NumVar(); k++ {
				switch v := rslt.arrays[k][b].(type) {
				case []string:
					u := da.GetPos(k).([]string)
					v = append(v, u[i])
					rslt.arrays[k][b] = v
				case []float64:
					u := da.GetPos(k).([]float64)
					v = append(v, u[i])
					rslt.arrays[k][b] = v
				case []float32:
					u := da.GetPos(k).([]float32)
					v = append(v, u[i])
					rslt.arrays[k][b] = v
				case []uint64:
					u := da.GetPos(k).([]uint64)
					v = append(v, u[i])
					rslt.arrays[k][b] = v
				case []uint32:
					u := da.GetPos(k).([]uint32)
					v = append(v, u[i])
					rslt.arrays[k][b] = v
				case []uint16:
					u := da.GetPos(k).([]uint16)
					v = append(v, u[i])
					rslt.arrays[k][b] = v
				case []uint8:
					u := da.GetPos(k).([]uint8)
					v = append(v, u[i])
					rslt.arrays[k][b] = v
				case []int64:
					u := da.GetPos(k).([]int64)
					v = append(v, u[i])
					rslt.arrays[k][b] = v
				case []int32:
					u := da.GetPos(k).([]int32)
					v = append(v, u[i])
					rslt.arrays[k][b] = v
				case []int16:
					u := da.GetPos(k).([]int16)
					v = append(v, u[i])
					rslt.arrays[k][b] = v
				case []int8:
					u := da.GetPos(k).([]int8)
					v = append(v, u[i])
					rslt.arrays[k][b] = v
				case []int:
					u := da.GetPos(k).([]int)
					v = append(v, u[i])
					rslt.arrays[k][b] = v
				default:
					msg := fmt.Sprintf("Regroup: unkown type %T\n", v)
					panic(msg)
				}
			}
		}
	}

	return rslt
}
