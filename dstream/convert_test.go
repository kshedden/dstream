package dstream

import (
	"bytes"
	"testing"
)

func TestConvert(t *testing.T) {

	x := [][]interface{}{
		{
			[]float64{0, 0, 0},
			[]float64{1, 1, 1},
			[]float64{2, 2, 3},
		},
		{
			[]string{"a", "b", "c"},
			[]string{"d", "e", "f"},
			[]string{"g", "h", "i"},
		},
		{
			[]float64{1, 2, 3},
			[]float64{4, 5, 6},
			[]float64{7, 8, 9},
		},
	}
	na := []string{"x1", "x2", "x3"}
	da := NewFromArrays(x, na)

	z := [][]interface{}{
		{
			[]int32{0, 0, 0},
			[]int32{1, 1, 1},
			[]int32{2, 2, 3},
		},
		{
			[]string{"a", "b", "c"},
			[]string{"d", "e", "f"},
			[]string{"g", "h", "i"},
		},
		{
			[]float32{1, 2, 3},
			[]float32{4, 5, 6},
			[]float32{7, 8, 9},
		},
	}
	db := NewFromArrays(z, na)

	// Perform two conversions
	da = Convert(da, "x1", Int32)
	da = Convert(da, "x3", Float32)

	if !EqualReport(da, db, true) {
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

	// Generate a new variable that is 100 times the id variable.
	times100 := func(v map[string]interface{}, z interface{}) {
		id := v["id"].([]float64)
		y := z.([]float64)
		for i := range id {
			y[i] = id[i] * 100
		}
	}

	b1 := bytes.NewReader([]byte(data1))
	tc := &CSVTypeConf{
		Float64: []string{"id", "v1", "v2", "v3"},
	}
	d1 := FromCSV(b1).TypeConf(tc).HasHeader().Done()

	d1 = Generate(d1, "id100", times100, Float64)
	d1 = Convert(d1, "id100", Uint64)
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
