package bloom

import (
	"fmt"
	"hash"
	"math/rand"

	"github.com/google/uuid"
	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	filter []uint8
	size   int32
}

var hashFns []hash.Hash32

func init() {
	// mHasher = murmur3.New32WithSeed(uint32(time.Now().Unix()))
	hashFns = make([]hash.Hash32, 0)
	for i := 0; i < 20; i++ {
		hashFns = append(hashFns, murmur3.New32WithSeed(rand.Uint32()))
	}
}

func murmurHash(key string, size int32, hashFnIdx int) int32 {
	hashFns[hashFnIdx].Write([]byte(key))
	var result = hashFns[hashFnIdx].Sum32() % uint32(size)
	hashFns[hashFnIdx].Reset()

	return int32(result)
}

// Initialise the bloom filter
func NewBloomFilter(size int32) *BloomFilter {
	return &BloomFilter{
		filter: make([]uint8, size),
		size:   size,
	}
}

func (bf *BloomFilter) Add(key string, numHashFuns int) {
	for i := 0; i < numHashFuns; i++ {

		idx := murmurHash(key, bf.size, i)
		arrayIdx := idx / 8
		bitIdx := idx % 8
		bf.filter[arrayIdx] = bf.filter[arrayIdx] | (1 << bitIdx)
	}

	// fmt.Println("wrote ", key, " at ", idx)
}

func (bf *BloomFilter) Print() {
	fmt.Println(bf.filter)
}

func (bf *BloomFilter) Exists(key string, numHashFuns int) (string, int32, bool) {
	for i := 0; i < numHashFuns; i++ {

		idx := murmurHash(key, bf.size, i)
		arrayIdx := idx / 8
		bitIdx := idx % 8
		exist := bf.filter[arrayIdx]&(1<<bitIdx) > 0

		if !exist {
			return key, idx, false
		}
	}

	return key, 0, true
}

// create the bloom filter
func CreateBloomFilter() {

	dataset := make([]string, 0)
	dataset_exists := make(map[string]bool)
	dataset_notexists := make(map[string]bool)

	for i := 0; i < 500; i++ {
		u := uuid.New()
		dataset = append(dataset, u.String())
		dataset_exists[u.String()] = true
	}

	for i := 0; i < 500; i++ {
		u := uuid.New()
		dataset = append(dataset, u.String())
		dataset_notexists[u.String()] = false
	}

	for j := 1; j <= len(hashFns); j++ {

		bloom := NewBloomFilter(int32(10000))

		// Insert the key in bloom filter
		for key, _ := range dataset_exists {
			bloom.Add(key, j)
		}

		falsePositives := 0
		// truePositives := 0
		for _, key := range dataset {
			_, _, exists := bloom.Exists(key, j)

			// if key exists
			if exists {

				// and in reality it exists
				// if _, ok := dataset_exists[key]; ok {
				// 	truePositives++
				// }

				// and in reality it doesn't exists
				if _, ok := dataset_notexists[key]; ok {
					falsePositives++
				}

			}
		}

		fmt.Println("number of hash functions: ", j, "fp_rate: ", float64(falsePositives)/float64(len(dataset)),
			"fp_rate_perc: ", 100*float64(falsePositives)/float64(len(dataset)))
		// fmt.Println(100 * (float64(falsePositives) / float64(len(dataset))))
	}
}
