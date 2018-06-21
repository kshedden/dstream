package dstream

import (
	"math"
)

// Stats contains summary statistics for a float64 Dstream variable.
type Stats struct {

	// The mean value
	Mean float64

	// The minimum value
	Min float64

	// The maximum value
	Max float64

	// The standard deviation of the values
	SD float64

	// The number of non inf/nan values
	N int

	// The number of Nan values
	NaN int

	// The number of Inf values
	Inf int
}

// Describe computes summary statistics for the float64 columns of a dstream.
func Describe(data Dstream) map[string]Stats {

	data.Reset()

	p := data.NumVar()
	stats := make([]Stats, p)
	first := true

	// Get the min, max and sum.
	for data.Next() {
		for j := 0; j < p; j++ {
			u := data.GetPos(j)
			x, ok := u.([]float64)
			if !ok {
				continue
			}

			for i, y := range x {

				if math.IsNaN(y) {
					stats[j].NaN++
					continue
				}

				if math.IsInf(y, 0) {
					stats[j].Inf++
					continue
				}

				stats[j].N++

				stats[j].Mean += y

				if (first && i == 0) || y < stats[j].Min {
					stats[j].Min = y
				}

				if (first && i == 0) || y > stats[j].Max {
					stats[j].Max = y
				}
			}
		}

		first = false
	}

	// Convert sum to mean.
	for j := range stats {
		stats[j].Mean /= float64(stats[j].N)
	}

	// Get the standard deviation.
	data.Reset()
	for data.Next() {
		for j := 0; j < p; j++ {
			u := data.GetPos(j)
			x, ok := u.([]float64)
			if !ok {
				continue
			}

			for _, y := range x {
				u := y - stats[j].Mean
				stats[j].SD += u * u
			}
		}
	}

	// Convert sum of squares to SD.
	for j := range stats {
		stats[j].SD = math.Sqrt(stats[j].SD / float64(stats[j].N))
	}

	// Put the statistics into a map indexed by variable names.
	stm := make(map[string]Stats)
	names := data.Names()
	for j := 0; j < p; j++ {
		stm[names[j]] = stats[j]
	}

	return stm
}
