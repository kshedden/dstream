package dstream

import (
	"bytes"
	"fmt"
)

func ExampleDStream_FromCSV() {

	data := `Food,Type,Weight,Price
Banana,Fruit,13,9
Cucumber,Vegetable,15,5
Cheese,Dairy,12,35
Lamb,Meat,40,76
`

	b := bytes.NewBuffer([]byte(data))

	da := FromCSV(b).SetStringVars([]string{"Food", "Type"}).SetFloatVars([]string{"Weight"}).HasHeader()

	da.Next() // Always call Next before first call to Get or GetPos

	y := da.Get("Type").([]string)
	fmt.Printf("%v\n", y)

	x := da.Get("Weight").([]float64)
	fmt.Printf("%v\n", x)

	// Output:
	// [Fruit Vegetable Dairy Meat]
	// [13 15 12 40]
}

func ExampleDStream_Mutate() {

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
	da := FromCSV(b).SetFloatVars([]string{"V1", "V2", "V3", "V4"})

	dx := Mutate(da, "V2", timesTwo)

	dx.Next() // Always call Next before first call to Get or GetPos

	y := dx.Get("V2")
	fmt.Printf("%v\n", y)

	// Output:
	// [4 6 8 10]
}
