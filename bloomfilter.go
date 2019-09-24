package bloomfilter

import (
	"hash"
	"hash/fnv"
	"math"
	"strconv"
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

func (b *BloomFilter) getHash(seed int, key string) (uint64, error) {
	t := []byte(strconv.Itoa(seed))
	var err error
	_, err = b.HashFunction.Write(t)
	_, err = b.HashFunction.Write([]byte(key))
	return b.HashFunction.Sum64(), err
}

// NewBloomFilter returns newly created BloomFilter struct
func NewBloomFilter(elements uint64, acceptableFalsePositiveProbability float64) BloomFilter {
	size := getSizeOfBitArray(elements, acceptableFalsePositiveProbability)
	hashFuncs := getOptimumNumOfHashFuncs(size, elements)

	return BloomFilter{size, make([]byte, size), hashFuncs, fnv.New64(), acceptableFalsePositiveProbability}
}

// Add adds new element in bloomfilter instance
func (b *BloomFilter) Add(element string) error {
	for i := 0; i < int(b.NumberOfHashFunctions); i++ {
		t, err := b.getHash(i, element)
		if err != nil {
			return err
		}
		position := t % uint64(b.Size)
		b.BitArray[position] = 1
	}
	return nil
}
