package bloom

import (
	"fmt"
	"hash"
	"time"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	filter []bool // TODO: Optimise this
	size   int32
}

var mHasher hash.Hash32

func init() {
	mHasher = murmur3.New32WithSeed(uint32(time.Now().Unix()))
}

func murmurHash(key string, size int32) int32 {
	mHasher.Write([]byte(key))
	var result = mHasher.Sum32() % uint32(size)
	mHasher.Reset()

	return int32(result)
}

// Initialise the bloom filter
func NewBloomFilter(size int32) *BloomFilter {
	return &BloomFilter{
		filter: make([]bool, size),
		size:   size,
	}
}

func (bf *BloomFilter) Add(key string) {
	idx := murmurHash(key, bf.size)
	bf.filter[idx] = true

	fmt.Println("wrote ", key, " at ", idx)
}

func (bf *BloomFilter) Exists(key string) bool {
	idx := murmurHash(key, bf.size)
	return bf.filter[idx]
}

// create the bloom filter
func CreateBloomFilter() {
	bloom := NewBloomFilter(16)
	keys := []string{"a", "b", "c", "d"}
	// keys := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"}

	for _, key := range keys {
		bloom.Add(key)
	}

	for _, key := range keys {
		fmt.Println("key, ", bloom.Exists(key))
	}
}
