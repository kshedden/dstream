package dstream

// VarPos returns a map from variable names to their corresponding
// column positions.
func VarPos(d Dstream) map[string]int {

	mp := make(map[string]int)

	for k, v := range d.Names() {
		mp[v] = k
	}

	return mp
}

// VarMap returns a map from variable names to data slices, in the current chunk.
func VarMap(d Dstream) map[string]interface{} {

	mp := make(map[string]interface{})

	for _, na := range d.Names() {
		mp[na] = d.Get(na)
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
	for i := range z {
		z[i] = 0
	}
}
