# bloomfilter

A Bloom filter is a data structure designed to tell you, rapidly and memory-efficiently, whether an element is present in a set. The price paid for this efficiency is that a Bloom filter is a probabilistic data structure: it tells us that the element either definitely is not in the set or may be in the set.

In this implementation, the hashing functions used is fnv hash, a non-cryptographic hashing function.

# Installation
```
go get -u github.com/darklord19/bloomfilter
```

# Usage
```go
// NewBloomFilter accepts two arguments. First is number of elements you want to track
// and second is acceptable false positive probability
bf := bloomfilter.NewBloomFilter(10000, 0.10) 
bf.Add("A")
bf.Add("B")
res, err := bf.DoesNotExist("C")
```
