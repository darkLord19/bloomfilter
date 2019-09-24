/*
Package bloomfilter implements bloomfilter data structure.

A Bloom filter is a data structure designed to tell you, rapidly
and memory-efficiently, whether an element is present in a set.
The price paid for this efficiency is that a Bloom filter is a probabilistic data structure:
it tells us that the element either definitely is not in the set or may be in the set.

Example:

	package main

	import "github.com/darkLord19/bloomfilter"

	func main() {
		bf := bloomfilter.NewBloomFilter(10000, 0.10)
		bf.Add("A")
		bf.Add("B")
		status, err := bf.DoesNotExist("C")
	}

*/
package bloomfilter

import (
	"hash"
	"hash/fnv"
	"math"
	"strconv"
)

// BloomFilter data struct
type BloomFilter struct {
	Size                               uint64
	BitArray                           []bool
	NumberOfHashFunctions              uint8
	HashFunction                       hash.Hash64
	AcceptableFalsePositiveProbability float64
}

func getSizeOfBitArray(elements uint64, prob float64) uint64 {
	return uint64(math.Round(
		-1 * float64(elements) * math.Log(prob) / math.Pow((math.Log(2)), 2)))
}

func getOptimumNumOfHashFuncs(sizeOfArray uint64, elements uint64) uint8 {
	return uint8(float64(sizeOfArray) / float64(elements) * math.Log(2))
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

	return BloomFilter{size, make([]bool, size), hashFuncs, fnv.New64(), acceptableFalsePositiveProbability}
}

// Add adds new element in bloomfilter instance
func (b *BloomFilter) Add(element string) error {
	for i := 0; i < int(b.NumberOfHashFunctions); i++ {
		t, err := b.getHash(i, element)
		if err != nil {
			return err
		}
		position := t % uint64(b.Size)
		b.BitArray[position] = true
	}
	return nil
}

// DoesNotExist checks if element does not exist for sure in our dataset
func (b *BloomFilter) DoesNotExist(element string) (bool, error) {
	for i := 0; i < int(b.NumberOfHashFunctions); i++ {
		t, err := b.getHash(i, element)
		if err != nil {
			return false, err
		}
		position := t % uint64(b.Size)
		if !b.BitArray[position] {
			return true, nil
		}
	}
	return false, nil
}
