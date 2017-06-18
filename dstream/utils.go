package dstream

//go:generate go run gen.go -template=utils.template

// VarPos returns a map from variable names to their corresponding
// column positions.
func VarPos(d Dstream) map[string]int {

	mp := make(map[string]int)

	for k, v := range d.Names() {
		mp[v] = k
	}

	return mp
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
