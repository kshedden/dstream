package dstream

//go:generate go run gen.go comparisons.template

// Equal returns true if the two Data values have equal contents.
func Equal(x, y Dstream) bool {
	return EqualReport(x, y, false)
}
