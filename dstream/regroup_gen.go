// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"fmt"
	"time"
)

const (
	bc = 100
)

func addchunk(da Dstream, rslt *DataFrame) {
	for j := 0; j < da.NumVar(); j++ {
		v := da.GetPos(j)
		switch v := v.(type) {

		case []string:
			rslt.data[j] = append(rslt.data[j], make([]string, 0, bc))

		case []time.Time:
			rslt.data[j] = append(rslt.data[j], make([]time.Time, 0, bc))

		case []uint8:
			rslt.data[j] = append(rslt.data[j], make([]uint8, 0, bc))

		case []uint16:
			rslt.data[j] = append(rslt.data[j], make([]uint16, 0, bc))

		case []uint32:
			rslt.data[j] = append(rslt.data[j], make([]uint32, 0, bc))

		case []uint64:
			rslt.data[j] = append(rslt.data[j], make([]uint64, 0, bc))

		case []int8:
			rslt.data[j] = append(rslt.data[j], make([]int8, 0, bc))

		case []int16:
			rslt.data[j] = append(rslt.data[j], make([]int16, 0, bc))

		case []int32:
			rslt.data[j] = append(rslt.data[j], make([]int32, 0, bc))

		case []int64:
			rslt.data[j] = append(rslt.data[j], make([]int64, 0, bc))

		case []float32:
			rslt.data[j] = append(rslt.data[j], make([]float32, 0, bc))

		case []float64:
			rslt.data[j] = append(rslt.data[j], make([]float64, 0, bc))

		default:
			msg := fmt.Sprintf("Regroup: unknown type %T\n", v)
			panic(msg)
		}
	}
}

func doRegroup(da Dstream, varpos int) *DataFrame {

	bucket := make(map[uint64]int)

	rslt := &DataFrame{
		xform: xform{
			names: da.Names(),
		},
		data: make([][]interface{}, da.NumVar()),
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
				switch v := rslt.data[k][b].(type) {
				case []string:
					u := da.GetPos(k).([]string)
					v = append(v, u[i])
					rslt.data[k][b] = v
				case []time.Time:
					u := da.GetPos(k).([]time.Time)
					v = append(v, u[i])
					rslt.data[k][b] = v
				case []uint8:
					u := da.GetPos(k).([]uint8)
					v = append(v, u[i])
					rslt.data[k][b] = v
				case []uint16:
					u := da.GetPos(k).([]uint16)
					v = append(v, u[i])
					rslt.data[k][b] = v
				case []uint32:
					u := da.GetPos(k).([]uint32)
					v = append(v, u[i])
					rslt.data[k][b] = v
				case []uint64:
					u := da.GetPos(k).([]uint64)
					v = append(v, u[i])
					rslt.data[k][b] = v
				case []int8:
					u := da.GetPos(k).([]int8)
					v = append(v, u[i])
					rslt.data[k][b] = v
				case []int16:
					u := da.GetPos(k).([]int16)
					v = append(v, u[i])
					rslt.data[k][b] = v
				case []int32:
					u := da.GetPos(k).([]int32)
					v = append(v, u[i])
					rslt.data[k][b] = v
				case []int64:
					u := da.GetPos(k).([]int64)
					v = append(v, u[i])
					rslt.data[k][b] = v
				case []float32:
					u := da.GetPos(k).([]float32)
					v = append(v, u[i])
					rslt.data[k][b] = v
				case []float64:
					u := da.GetPos(k).([]float64)
					v = append(v, u[i])
					rslt.data[k][b] = v
				default:
					msg := fmt.Sprintf("Regroup: unkown type %T\n", v)
					panic(msg)
				}
			}
		}
	}

	return rslt
}
