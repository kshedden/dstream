package dstream

import (
	"bytes"
	"testing"
)

func TestConvert(t *testing.T) {

	x1 := []interface{}{
		[]float64{0, 0, 0},
		[]float64{1, 1, 1},
		[]float64{2, 2, 3},
	}
	x2 := []interface{}{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
	}
	x3 := []interface{}{
		[]float64{1, 2, 3},
		[]float64{4, 5, 6},
		[]float64{7, 8, 9},
	}
	dat := [][]interface{}{x1, x2, x3}
	na := []string{"x1", "x2", "x3"}
	da := NewFromArrays(dat, na)

	z1 := []interface{}{
		[]int32{0, 0, 0},
		[]int32{1, 1, 1},
		[]int32{2, 2, 3},
	}
	z2 := []interface{}{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
	}
	z3 := []interface{}{
		[]float32{1, 2, 3},
		[]float32{4, 5, 6},
		[]float32{7, 8, 9},
	}
	dat = [][]interface{}{z1, z2, z3}
	db := NewFromArrays(dat, na)

	dx := Convert(da, "x1", "int32")
	dx = Convert(dx, "x3", "float32")

	if !EqualReport(dx, db, true) {
		t.Fail()
	}
}

func TestConvert2(t *testing.T) {

	data1 := `id,v1,v2,v3
1,2,3,4
1,3,4,5
2,4,5,6
3,5,6,7
3,99,99,99
3,100,101,102
4,200,201,202
`
	times100 := func(v map[string]interface{}, z interface{}) {
		id := v["id"].([]float64)
		y := z.([]float64)
		for i := range id {
			y[i] = id[i] * 100
		}
	}
	_ = times100

	b1 := bytes.NewReader([]byte(data1))
	d1 := FromCSV(b1).SetFloat64Vars("id", "v1", "v2", "v3").HasHeader().Done()
	d1 = Generate(d1, "id100", times100, "float64")
	d1 = Convert(d1, "id100", "uint64")
	d1 = DropCols(d1, "id")

	if !d1.Next() {
		t.Fail()
	}
	z := d1.Get("id100").([]uint64)
	y := []uint64{100, 100, 200, 300, 300, 300, 400}
	for i, u := range z {
		if u != y[i] {
			t.Fail()
		}
	}
}
