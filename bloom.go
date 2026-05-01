package bloom

type BloomFilter struct {
	bitset []uint64 // array storing packed bits
	m      uint64   // number of bits
	k      uint64   // number of hash functions
}

func New(m uint64, k uint64) *BloomFilter {
	size := (m + 63) / 64 // 64 bit blocks needed to hold m bits

	return &BloomFilter{
		bitset: make([]uint64, size),
		m:      m,
		k:      k,
	}
}

func NewWithEstimates(n uint64, p float64) *BloomFilter

func (f *BloomFilter) Add(data []byte) {

}
func (f *BloomFilter) Check(data []byte) {

}

func (f *BloomFilter) Reset() {
	for i := range f.bitset {
		f.bitset[i] = 0
	}
}
