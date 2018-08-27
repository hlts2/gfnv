package gfnv

/**
* see: https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function
**/

// group of fnv offset basis
const (
	// FnvOffsetBasis32 represents fnv offset basis for 32bit
	FnvOffsetBasis32 uint32 = 2166136261

	// FnvOffsetBasis64 represents fnv offset basis for 64bit
	FnvOffsetBasis64 uint64 = 14695981039346656037
)

// group of fnv prime
const (
	// FnvPrime32 represents fnv prime for 32bit
	FnvPrime32 uint32 = 16777619

	// FnvPrime64 represents fnv prime for 64bit
	FnvPrime64 uint64 = 1099511628211
)

// Fnv32 returns fnv32 hash value
func Fnv32(message []byte) uint32 {
	hash := FnvOffsetBasis32

	for i := 0; i < len(message); i++ {
		hash = (FnvPrime32 * hash) ^ uint32(message[i])
	}
	return hash
}

// Fnv64 returns fnv64 hash value
func Fnv64(message []byte) uint64 {
	hash := FnvOffsetBasis64
	for i := 0; i < len(message); i++ {
		hash = (FnvPrime64 * hash) ^ uint64(message[i])
	}
	return hash
}
