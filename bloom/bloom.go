package bloom

import (
	"fmt"
	"hash"

	"github.com/google/uuid"
	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	filter []bool // TODO: Optimise this
	size   int32
}

var mHasher hash.Hash32

func init() {
	// mHasher = murmur3.New32WithSeed(uint32(time.Now().Unix()))
	mHasher = murmur3.New32WithSeed(uint32(11))
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

	// fmt.Println("wrote ", key, " at ", idx)
}

func (bf *BloomFilter) Print() {
	fmt.Println(bf.filter)
}

func (bf *BloomFilter) Exists(key string) (string, int32, bool) {
	idx := murmurHash(key, bf.size)
	return key, idx, bf.filter[idx]
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

	for j := 100; j < 30000; j += 1000 {

		bloom := NewBloomFilter(int32(j))

		// Insert the key in bloom filter
		for key, _ := range dataset_exists {
			bloom.Add(key)
		}

		falsePositives := 0
		// truePositives := 0
		for _, key := range dataset {
			_, _, exists := bloom.Exists(key)

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

		fmt.Println("size: ", j, "fp_rate: ", float64(falsePositives)/float64(len(dataset)),
			"fp_rate_perc: ", 100*float64(falsePositives)/float64(len(dataset)))
		// fmt.Println(100 * (float64(falsePositives) / float64(len(dataset))))
	}

	// keys := []string{"a", "b", "c", "d"}
	// keys := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"}

	// for _, key := range keys {
	// 	bloom.Add(key)
	// }

	// for _, key := range keys {
	// 	fmt.Println(bloom.Exists(key))
	// }

	// fmt.Println(bloom.Exists("e"))
	// fmt.Println(bloom.Exists("f"))
	// fmt.Println(bloom.Exists("g"))
	// fmt.Println(bloom.Exists("h"))
}
