package bloom

import "github.com/supercoolpencil/pollen/hash"

type BloomFilter struct {
	bitset []uint64 // array storing packed bits
	m      uint64   // number of bits
	k      uint32   // number of hash functions
}

// New Create a bloom filter with specified m and k
func New(m uint64, k uint32) *BloomFilter {
	size := (m + 63) / 64 // 64 bit blocks needed to hold m bits

	return &BloomFilter{
		bitset: make([]uint64, size),
		m:      m,
		k:      k,
	}
}

// NewWithEstimates Create a bloom filter with specified n and p
// n: number of entries
// p: probability of false positives
func NewWithEstimates(n uint64, p float64) *BloomFilter {
	m := OptimalM(n, p)
	k := OptimalK(m, n)
	return New(m, k)
}

func (f *BloomFilter) Add(data []byte) {
	baseHash := hash.Hash(data)

	// Kirsch-Mitzenmacher (https://www.eecs.harvard.edu/~michaelm/postscripts/rsa2008.pdf)
	h1 := uint32(baseHash >> 32)
	h2 := uint32(baseHash)

	for i := uint64(0); i < uint64(f.k); i++ {

		// Use double hashing to add non-linearity (https://en.wikipedia.org/wiki/Double_hashing#Enhanced_double_hashing)
		cubic := (i*i*i - i) / 6

		combinedHash := uint64(h1) + (i * uint64(h2)) + cubic

		targetBit := combinedHash % f.m

		blockIndex := targetBit / 64
		blockOffset := targetBit % 64

		f.bitset[blockIndex] |= 1 << blockOffset
	}
}
func (f *BloomFilter) Check(data []byte) bool {

}

func (f *BloomFilter) Reset() {
	for i := range f.bitset {
		f.bitset[i] = 0
	}
}
