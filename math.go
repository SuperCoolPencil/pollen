package bloom

import "math"

// OptimalM calculates the required bit array size (m).
// n: expected number of items
// p: acceptable false positive probability
func OptimalM(n uint64, p float64) uint64 {
	// Formula: m = -(n * ln(p)) / (ln(2)^2)
	numerator := -float64(n) * math.Log(p)
	denominator := math.Pow(math.Log(2), 2)

	return uint64(math.Ceil(numerator / denominator))
}

// OptimalK calculates the optimal number of hash functions (k).
// m: bit array size
// n: expected number of items
func OptimalK(m uint64, n uint64) uint32 {
	// Formula: k = (m/n)*ln(2)
	k := (float64(m) / float64(n)) * math.Log(2)
	return uint32(k)
}
