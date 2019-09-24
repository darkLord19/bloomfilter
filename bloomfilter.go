package bloomfilter

import (
	"math"
)

// BloomFilter data struct
type BloomFilter struct {
	Size                               byte
	HashFunctions                      byte
	AcceptableFalsePositiveProbability float64
	BitArray                           []byte
}

func getSizeOfBitArray(elements uint64, prob float64) byte {
	return byte(math.Round(
		-1 * float64(elements) * math.Log(prob) / math.Pow((math.Log(2)), 2)))
}
