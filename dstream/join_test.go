package dstream

import (
	"math"
	"os"
	"testing"
)

func TestJoin1(t *testing.T) {

	ar := make([][]uint64, 3)
	for j := 0; j < 3; j++ {
		x := make([]uint64, 20+3*j)
		for i, _ := range x {
			x[i] = uint64(math.Floor(float64(i) / float64(j+3)))
		}
		ar[j] = x
	}

	sizes := [][]int{
		{3, 4, 5},
		{3, 4, 5},
		{3, 4, 5},
		{3, 4, 5},
		{3, 4, 5},
		{3, 3, 1},
		{2, 0, 0},
	}

	da0 := NewFromArrays([][]interface{}{[]interface{}{ar[0]}}, []string{"id"})
	da1 := NewFromArrays([][]interface{}{[]interface{}{ar[1]}}, []string{"id"})
	da2 := NewFromArrays([][]interface{}{[]interface{}{ar[2]}}, []string{"id"})
	da := []Dstream{da0, da1, da2}

	for j := 0; j < 3; j++ {
		da[j] = Segment(da[j], []string{"id"})
	}

	for da[0].Next() {

	}

	join := NewJoin(da, []string{"id", "id", "id"})

	jj := 0
	for join.Next() {
		n0 := len(join.Data[0].GetPos(0).([]uint64))
		n1 := len(join.Data[1].GetPos(0).([]uint64))
		n2 := len(join.Data[2].GetPos(0).([]uint64))
		s := sizes[jj]
		if n0 != s[0] || n1 != s[1] || n2 != s[2] {
			t.Fail()
		}
		jj++
	}

	for k := 0; k < 3; k++ {
		f, msg := checkPosName(join.Data[k])
		if !f {
			os.Stderr.WriteString(msg)
		}
	}
}
