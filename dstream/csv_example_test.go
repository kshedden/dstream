package dstream

import (
	"bytes"
	"fmt"
)

func ExampleCSV() {

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
