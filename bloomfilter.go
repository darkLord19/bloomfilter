package bloomfilter

// BloomFilter data struct
type BloomFilter struct {
	Size                               byte
	HashFunctions                      byte
	AcceptableFalsePositiveProbability float64
	BitArray                           []byte
}
