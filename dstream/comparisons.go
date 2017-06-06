package dstream

//go:generate go run gen.go comparisons.template

// Equal returns true if the two Dstream values have equal contents.
func Equal(x, y Dstream) bool {
	return EqualReport(x, y, false)
}
