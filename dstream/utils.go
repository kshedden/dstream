package dstream

import "fmt"

//go:generate go run gen.go utils.template

func VarPos(d Dstream) map[string]int {

	mp := make(map[string]int)

	for k, v := range d.Names() {
		mp[v] = k
	}

	return mp
}

func VarTypes(d Dstream) {
	for k, na := range d.Names() {
		v := d.GetPos(k)
		switch v.(type) {
		case []float64:
			fmt.Printf("%s float64\n", na)
		case []string:
			fmt.Printf("%s string\n", na)
		case nil:
			fmt.Printf("%s nil\n", na)
		default:
			fmt.Printf("%s other\n", na)
		}
	}
}

func wherefalse(ma []bool, pos []int) []int {
	pos = pos[0:0]
	for i, v := range ma {
		if !v {
			pos = append(pos, i)
		}
	}
	return pos
}

func resizeBool(x []bool, n int) []bool {
	if cap(x) < n {
		x = make([]bool, n)
	}
	return x[0:n]
}

func zeroFloat(z []float64) {
	z = z[0:cap(z)]
	for i, _ := range z {
		z[i] = 0
	}
}
