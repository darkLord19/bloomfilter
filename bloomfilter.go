package bloomfilter

import (
	"hash"
	"hash/fnv"
	"math"
)

// BloomFilter data struct
type BloomFilter struct {
	Size                               byte
	BitArray                           []byte
	NumberOfHashFunctions              byte
	HashFunction                       hash.Hash64
	AcceptableFalsePositiveProbability float64
}

func getSizeOfBitArray(elements uint64, prob float64) byte {
	return byte(math.Round(
		-1 * float64(elements) * math.Log(prob) / math.Pow((math.Log(2)), 2)))
}

func getOptimumNumOfHashFuncs(sizeOfArray byte, elements uint64) byte {
	return byte((float64(sizeOfArray) / float64(elements)) * math.Log(2))
}

// NewBloomFilter returns newly created BloomFilter struct
func NewBloomFilter(elements uint64, acceptableFalsePositiveProbability float64) BloomFilter {
	size := getSizeOfBitArray(elements, acceptableFalsePositiveProbability)
	hashFuncs := getOptimumNumOfHashFuncs(size, elements)

	return BloomFilter{size, make([]byte, size), hashFuncs, fnv.New64(), acceptableFalsePositiveProbability}
}
