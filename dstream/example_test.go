package dstream

import (
	"bytes"
	"fmt"
)

func ExampleFromCSV() {

	data := `Food,Type,Weight,Price
Banana,Fruit,13,9
Cucumber,Vegetable,15,5
Cheese,Dairy,12,35
Lamb,Meat,40,76
`

	b := bytes.NewBuffer([]byte(data))

	da := FromCSV(b).SetStringVars([]string{"Food", "Type"}).SetFloatVars([]string{"Weight"}).HasHeader().Done()
	da.Next() // Always call Next before first call to Get or GetPos

	y := da.Get("Type").([]string)
	fmt.Printf("%v\n", y)

	x := da.Get("Weight").([]float64)
	fmt.Printf("%v\n", x)

	// Output:
	// [Fruit Vegetable Dairy Meat]
	// [13 15 12 40]
}

func ExampleMutate() {

	data := `1,2,3,4
2,3,4,5
3,4,5,6
4,5,6,7
`

	// A mutating function, scales all values by 2.
	timesTwo := func(x interface{}) {
		v := x.([]float64)
		for i := range v {
			v[i] *= 2
		}
	}

	b := bytes.NewBuffer([]byte(data))
	da := FromCSV(b).SetFloatVars([]string{"V1", "V2", "V3", "V4"}).Done()
	da = Mutate(da, "V2", timesTwo)

	da.Next() // Always call Next before first call to Get or GetPos

	y := da.Get("V2")
	fmt.Printf("%v\n", y)

	// Output:
	// [4 6 8 10]
}

func ExampleFilter() {

	data := `V1,V2,V3,V4
1,2,3,4
2,0,4,5
3,4,5,6
4,0,6,7
`

	// A filtering function, selects if not equal to 0.
	f := func(x interface{}, b []bool) bool {
		v := x.([]float64)
		var any bool
		for i := range v {
			b[i] = v[i] != 0
			any = any || !b[i]
		}
		return any
	}

	b := bytes.NewBuffer([]byte(data))
	da := FromCSV(b).SetFloatVars([]string{"V1", "V2", "V3", "V4"}).HasHeader().Done()
	da = Filter(da, map[string]FilterFunc{"V2": f})

	da.Next() // Always call Next before first call to Get or GetPos

	y := da.Get("V1")
	fmt.Printf("%v\n", y)

	// Output:
	// [1 3]
}

func ExampleSegment() {

	data := `V1,V2,V3,V4
1,2,3,4
1,0,4,5
2,4,5,6
3,0,6,7
`

	b := bytes.NewBuffer([]byte(data))
	da := FromCSV(b).SetFloatVars([]string{"V1", "V2", "V3", "V4"}).HasHeader().Done()
	da = Segment(da, []string{"V1"})

	for da.Next() {
		y := da.Get("V2")
		fmt.Printf("%v\n", y)
	}

	// Output:
	// [2 0]
	// [4]
	// [0]
}

func ExampleApply() {

	data := `V1,V2,V3,V4
1,2,3,4
1,0,4,5
2,4,5,6
3,0,6,7
`

	f := func(v map[string]interface{}, x interface{}) {
		v1 := v["V1"].([]float64)
		v2 := v["V2"].([]float64)
		y := x.([]float64)
		for i := range v1 {
			y[i] = v1[i] + v2[i]
		}
	}

	b := bytes.NewBuffer([]byte(data))
	da := FromCSV(b).SetFloatVars([]string{"V1", "V2", "V3", "V4"}).HasHeader().Done()
	da = Apply(da, "V1p2", f, "float64")

	for da.Next() {
		y := da.Get("V1p2")
		fmt.Printf("%v\n", y)
	}

	// Output:
	// [3 1 6 3]
}

func ExampleDiffChunk() {

	data := `V1,V2,V3,V4
1,2,3,4
1,0,4,5
2,4,5,6
3,0,6,8
3,1,5,9
`

	b := bytes.NewBuffer([]byte(data))
	da := FromCSV(b).SetFloatVars([]string{"V1", "V2", "V3", "V4"}).HasHeader().Done()
	da = DiffChunk(da, map[string]int{"V2": 1, "V4": 2})

	for da.Next() {
		y := da.Get("V2$d1")
		fmt.Printf("%v\n", y)
		y = da.Get("V4$d2")
		fmt.Printf("%v\n", y)
	}

	// Output:
	// [4 -4 1]
	// [0 1 -1]
}

func ExampleLagChunk() {

	data := `1,2,3,4
2,3,4,5
3,4,5,6
4,5,6,7
`

	b := bytes.NewBuffer([]byte(data))
	da := FromCSV(b).SetFloatVars([]string{"V1", "V2", "V3", "V4"}).Done()
	da = LagChunk(da, map[string]int{"V2": 2})

	da.Next() // Always call Next before first call to Get or GetPos

	y := da.Get("V2[0]")
	fmt.Printf("%v\n", y)
	y = da.Get("V2[-1]")
	fmt.Printf("%v\n", y)

	// Output:
	// [4 5]
	// [3 4]
}

func ExampleJoin() {

	data1 := `1,2,3,4
1,3,4,5
3,4,5,6
3,5,6,7
`

	data2 := `1,2,3
1,3,4
1,4,5
3,5,6
`

	data3 := `2,2,3,5,6
2,3,4,7,5
3,4,5,3,4
4,5,6,2,3
`

	names := [][]string{{"V1", "V2", "V3", "V4"},
		{"V1", "V2", "V3"},
		{"V1", "V2", "V3", "V4", "V5"}}
	var da []Dstream

	for j, data := range []string{data1, data2, data3} {
		b := bytes.NewBuffer([]byte(data))
		d := FromCSV(b).SetFloatVars(names[j]).Done()
		d = Convert(d, "V1", "uint64")
		d = Segment(d, []string{"V1"})
		da = append(da, d)
	}

	join := NewJoin(da, []string{"V1", "V1", "V1"})

	for join.Next() {
		fmt.Printf("%v\n", da[0].GetPos(0))
		if join.Status[1] {
			fmt.Printf("%v\n", da[1].GetPos(0))
		}
		if join.Status[2] {
			fmt.Printf("%v\n\n", da[2].GetPos(0))
		}
	}

	// Output:
	// [1 1]
	// [1 1 1]
	// [3 3]
	// [3]
	// [3]
}

func ExampleRegroup() {

	data := `V1,V2,V3
1,2,3
3,3,4
2,4,5
2,5,6
5,2,3
0,3,4
1,4,5
5,5,6
`

	b := bytes.NewBuffer([]byte(data))
	d := FromCSV(b).SetFloatVars([]string{"V1", "V2", "V3"}).HasHeader().Done()
	d = Convert(d, "V1", "uint64")
	d = Regroup(d, "V1", true)

	for d.Next() {
		fmt.Printf("%v\n", d.GetPos(0))
		fmt.Printf("%v\n", d.GetPos(1))
		fmt.Printf("%v\n\n", d.GetPos(2))
	}

	// Output:
	// [0]
	// [3]
	// [4]
	//
	// [1 1]
	// [2 4]
	// [3 5]
	//
	// [2 2]
	// [4 5]
	// [5 6]
	//
	// [3]
	// [3]
	// [4]
	//
	// [5 5]
	// [2 5]
	// [3 6]
}
