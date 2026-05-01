package hash

// 64 bit implementation of FNV-1a
// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function

const (
	fnv_offset_basis uint64 = 0xcbf29ce484222325
	fnv_prime        uint64 = 0x100000001b3
)

// Hash takes a byte slice and returns a 64-bit FNV-1a hash
func Hash(input []byte) uint64 {

	// algorithm fnv-1a is:
	//     hash := FNV_offset_basis
	//
	//     for each byte_of_data to be hashed do
	//         hash := hash XOR byte_of_data
	//         hash := hash × FNV_prime
	//
	//     return hash

	hash := fnv_offset_basis

	for _, b := range input {
		hash ^= uint64(b)
		hash *= fnv_prime
	}

	return hash
}
