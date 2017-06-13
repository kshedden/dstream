package formula

import (
	"fmt"
	"math"
	"testing"

	"github.com/gonum/floats"
	"github.com/kshedden/dstream/dstream"
)

func tokEq(a, b []*token) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if *b[i] != *x {
			return false
		}
	}
	return true
}

func colSetEq(a, b *ColSet) bool {

	if len(a.Names) != len(b.Names) {
		return false
	}
	for i, x := range a.Names {
		if x != b.Names[i] {
			return false
		}
	}

	if len(a.Data) != len(b.Data) {
		return false
	}
	eq := func(x, y float64) bool { return math.Abs(x-y) < 1e-5 }
	for i, x := range a.Data {
		if !floats.EqualFunc(x, b.Data[i], eq) {
			return false
		}
	}

	return true
}

func TestLexParse(t *testing.T) {

	v, err := lex("(A + b)*c + d*f(e)")
	if err != nil {
		t.Fail()
		return
	}
	exp := []*token{&token{symbol: leftp}, &token{name: "A"},
		&token{symbol: plus}, &token{name: "b"}, &token{symbol: rightp},
		&token{symbol: times}, &token{name: "c"}, &token{symbol: plus},
		&token{name: "d"}, &token{symbol: times},
		&token{symbol: funct, name: "f(e)", funcn: "f", arg: "e"}}

	if !tokEq(v, exp) {
		t.Fail()
	}

	b, err := parse(v)
	if err != nil {
		t.Fail()
		return
	}
	exp = []*token{&token{name: "A"}, &token{name: "b"},
		&token{symbol: plus}, &token{name: "c"}, &token{symbol: times},
		&token{name: "d"},
		&token{symbol: funct, name: "f(e)", funcn: "f", arg: "e"},
		&token{symbol: times}, &token{symbol: plus},
	}

	if !tokEq(b, exp) {
		t.Fail()
	}
}

// Create some functions
func makeFuncs() map[string]Func {
	funcs := make(map[string]Func)
	funcs["square"] = func(na string, x []float64) *ColSet {
		y := make([]float64, len(x))
		for i, v := range x {
			y[i] = v * v
		}
		return &ColSet{Names: []string{na}, Data: [][]float64{y}}
	}
	funcs["pbase"] = func(na string, x []float64) *ColSet {
		y := make([]float64, len(x))
		z := make([]float64, len(x))
		for i, v := range x {
			y[i] = v * v
			z[i] = v * v * v
		}
		return &ColSet{Names: []string{na + "^2", na + "^3"}, Data: [][]float64{y, z}}
	}
	return funcs
}

func simpleData() ([][]interface{}, []string) {

	rawData := make([][]interface{}, 4)
	rawNames := []string{"x1", "x2", "x3", "x4"}

	rawData[0] = []interface{}{
		[]float64{0, 1, 2},
		[]float64{3, 4},
	}

	rawData[1] = []interface{}{
		[]string{"0", "0", "0"},
		[]string{"1", "1"},
	}

	rawData[2] = []interface{}{
		[]string{"a", "b", "a"},
		[]string{"b", "a"},
	}

	rawData[3] = []interface{}{
		[]float64{-1, 0, 1},
		[]float64{0, -1},
	}

	return rawData, rawNames
}

func TestSingle(t *testing.T) {

	rawData, rawNames := simpleData()
	funcs := makeFuncs()

	type rec struct {
		formula   string
		reflevels map[string]string
		expected  []*ColSet
		funcs     map[string]Func
	}

	for ip, pr := range []rec{
		{
			formula:   "x1 + x2 + x1*x2",
			reflevels: map[string]string{"x2": "0"},
			expected: []*ColSet{
				&ColSet{
					Names: []string{"x1", "x2[1]", "x1:x2[1]"},
					Data: [][]float64{
						[]float64{0, 1, 2},
						[]float64{0, 0, 0},
						[]float64{0, 0, 0},
					},
				},
				&ColSet{
					Names: []string{"x1", "x2[1]", "x1:x2[1]"},
					Data: [][]float64{
						[]float64{3, 4},
						[]float64{1, 1},
						[]float64{3, 4},
					},
				},
			},
		},
		{
			formula:   "x1 + x2 + x1*x2",
			reflevels: map[string]string{"x2": "1"},
			expected: []*ColSet{
				&ColSet{
					Names: []string{"x1", "x2[0]", "x1:x2[0]"},
					Data: [][]float64{
						[]float64{0, 1, 2},
						[]float64{1, 1, 1},
						[]float64{0, 1, 2},
					},
				},
				&ColSet{
					Names: []string{"x1", "x2[0]", "x1:x2[0]"},
					Data: [][]float64{
						[]float64{3, 4},
						[]float64{0, 0},
						[]float64{0, 0},
					},
				},
			},
		},
		{
			formula: "x1",
			expected: []*ColSet{
				&ColSet{
					Names: []string{"x1"},
					Data: [][]float64{
						[]float64{0, 1, 2},
					},
				},
				&ColSet{
					Names: []string{"x1"},
					Data: [][]float64{
						[]float64{3, 4},
					},
				},
			},
		},
		{
			formula:   "( ( x2*x3))",
			reflevels: map[string]string{"x2": "0", "x3": "a"},
			expected: []*ColSet{
				&ColSet{
					Names: []string{"x2[1]:x3[b]"},
					Data: [][]float64{
						[]float64{0, 0, 0},
					},
				},
				&ColSet{
					Names: []string{"x2[1]:x3[b]"},
					Data: [][]float64{
						[]float64{1, 0},
					},
				},
			},
		},
		{
			formula:   "(x1+x2)*(x3+x4)",
			reflevels: map[string]string{"x2": "0", "x3": "a"},
			expected: []*ColSet{
				&ColSet{
					Names: []string{"x1:x3[b]", "x1:x4", "x2[1]:x3[b]", "x2[1]:x4"},
					Data: [][]float64{
						[]float64{0, 1, 0},
						[]float64{0, 0, 2},
						[]float64{0, 0, 0},
						[]float64{0, 0, 0},
					},
				},
				&ColSet{
					Names: []string{"x1:x3[b]", "x1:x4", "x2[1]:x3[b]", "x2[1]:x4"},
					Data: [][]float64{
						[]float64{3, 0},
						[]float64{0, -4},
						[]float64{1, 0},
						[]float64{0, -1},
					},
				},
			},
		},
		{
			formula:   "x4 + (x1+x2)*x3",
			reflevels: map[string]string{"x2": "1", "x3": "a"},
			expected: []*ColSet{
				&ColSet{
					Names: []string{"x4", "x1:x3[b]", "x2[0]:x3[b]"},
					Data: [][]float64{
						[]float64{-1, 0, 1},
						[]float64{0, 1, 0},
						[]float64{0, 1, 0},
					},
				},
				&ColSet{
					Names: []string{"x4", "x1:x3[b]", "x2[0]:x3[b]"},
					Data: [][]float64{
						[]float64{0, -1},
						[]float64{3, 0},
						[]float64{0, 0},
					},
				},
			},
		},
		{
			formula: "1 + x1",
			expected: []*ColSet{
				&ColSet{
					Names: []string{"icept", "x1"},
					Data: [][]float64{
						[]float64{1, 1, 1},
						[]float64{0, 1, 2},
					},
				},
				&ColSet{
					Names: []string{"icept", "x1"},
					Data: [][]float64{
						[]float64{1, 1},
						[]float64{3, 4},
					},
				},
			},
		},
		{
			formula: "x1 + 1",
			expected: []*ColSet{
				&ColSet{
					Names: []string{"x1", "icept"},
					Data: [][]float64{
						[]float64{0, 1, 2},
						[]float64{1, 1, 1},
					},
				},
				&ColSet{
					Names: []string{"x1", "icept"},
					Data: [][]float64{
						[]float64{3, 4},
						[]float64{1, 1},
					},
				},
			},
		},
		{
			formula: "square(x1) + 1",
			expected: []*ColSet{
				&ColSet{
					Names: []string{"square(x1)", "icept"},
					Data: [][]float64{
						[]float64{0, 1, 4},
						[]float64{1, 1, 1},
					},
				},
				&ColSet{
					Names: []string{"square(x1)", "icept"},
					Data: [][]float64{
						[]float64{9, 16},
						[]float64{1, 1},
					},
				},
			},
		},
		{
			formula: "1 + pbase(x1)",
			expected: []*ColSet{
				&ColSet{
					Names: []string{"icept", "pbase(x1)^2", "pbase(x1)^3"},
					Data: [][]float64{
						[]float64{1, 1, 1},
						[]float64{0, 1, 4},
						[]float64{0, 1, 8},
					},
				},
				&ColSet{
					Names: []string{"icept", "pbase(x1)^2", "pbase(x1)^3"},
					Data: [][]float64{
						[]float64{1, 1},
						[]float64{9, 16},
						[]float64{27, 64},
					},
				},
			},
		},
		{
			formula: "1 + square(x1)",
			expected: []*ColSet{
				&ColSet{
					Names: []string{"icept", "square(x1)"},
					Data: [][]float64{
						[]float64{1, 1, 1},
						[]float64{0, 1, 4},
					},
				},
				&ColSet{
					Names: []string{"icept", "square(x1)"},
					Data: [][]float64{
						[]float64{1, 1},
						[]float64{9, 16},
					},
				},
			},
		},
	} {
		dp := dstream.NewFromArrays(rawData, rawNames)
		fp := New(pr.formula, dp).RefLevels(pr.reflevels).Funcs(funcs).Done()

		chunk := 0
		for fp.Next() {
			if !colSetEq(pr.expected[chunk], fp.Data) {
				fmt.Printf("Mismatch:\nip=%d\n", ip)
				fmt.Printf("chunk=%d\n", chunk)
				fmt.Printf("Expected: %v\n", pr.expected[chunk])
				fmt.Printf("Observed: %v\n", fp.Data)
				t.Fail()
			}
			chunk++
		}

		if fp.ErrorState != nil {
			fmt.Printf("ip=%d %v\n", ip, fp.ErrorState)
			t.Fail()
		}

		if chunk != len(pr.expected) {
			fmt.Printf("ip=%d wrong number of chunks\n", ip)
			t.Fail()
		}
	}
}

func TestMulti(t *testing.T) {

	rawData, rawNames := simpleData()
	funcs := makeFuncs()

	type rec struct {
		formulas  []string
		reflevels map[string]string
		expected  []*ColSet
		funcs     map[string]Func
	}

	for ip, pr := range []rec{
		{
			formulas:  []string{"x1"},
			reflevels: nil,
			expected: []*ColSet{
				&ColSet{
					Names: []string{"x1"},
					Data: [][]float64{
						[]float64{0, 1, 2},
					},
				},
				&ColSet{
					Names: []string{"x1"},
					Data: [][]float64{
						[]float64{3, 4},
					},
				},
			},
		},
		{
			formulas:  []string{"x1", "x2"},
			reflevels: map[string]string{"x2": "1"},
			expected: []*ColSet{
				&ColSet{
					Names: []string{"x1", "x2[0]"},
					Data: [][]float64{
						[]float64{0, 1, 2},
						[]float64{1, 1, 1},
					},
				},
				&ColSet{
					Names: []string{"x1", "x2[0]"},
					Data: [][]float64{
						[]float64{3, 4},
						[]float64{0, 0},
					},
				},
			},
		},
		{
			formulas:  []string{"x1", "square(x1) + x2"},
			reflevels: map[string]string{"x2": "1"},
			expected: []*ColSet{
				&ColSet{
					Names: []string{"x1", "square(x1)", "x2[0]"},
					Data: [][]float64{
						[]float64{0, 1, 2},
						[]float64{0, 1, 4},
						[]float64{1, 1, 1},
					},
				},
				&ColSet{
					Names: []string{"x1", "square(x1)", "x2[0]"},
					Data: [][]float64{
						[]float64{3, 4},
						[]float64{9, 16},
						[]float64{0, 0},
					},
				},
			},
		},
	} {
		dp := dstream.NewFromArrays(rawData, rawNames)
		fp := NewMulti(pr.formulas, dp).RefLevels(pr.reflevels).Funcs(funcs).Done()

		chunk := 0
		for fp.Next() {
			if !colSetEq(pr.expected[chunk], fp.Data) {
				fmt.Printf("Mismatch:\nip=%d\n", ip)
				fmt.Printf("chunk=%d\n", chunk)
				fmt.Printf("Expected: %v\n", pr.expected[chunk])
				fmt.Printf("Observed: %v\n", fp.Data)
				t.Fail()
			}
			chunk++
		}
	}
}

func TestReg(t *testing.T) {

	rawData := make([][]interface{}, 3)
	rawNames := []string{"y", "x1", "x2"}

	rawData[0] = []interface{}{
		[]float64{1, 2, 3},
		[]float64{4, 5},
	}
	rawData[1] = []interface{}{
		[]float64{1, 1, 1},
		[]float64{0, 0},
	}
	rawData[2] = []interface{}{
		[]float64{0, 1, 0},
		[]float64{1, 0},
	}

	dp := dstream.NewFromArrays(rawData, rawNames)
	fdp := New("x1 + x2", dp).Done()
	_ = dstream.NewReg(fdp, "y", []string{"x1", "x2", "", ""}, "", "")
}
