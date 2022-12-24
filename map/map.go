package basicmap

import "hash/fnv"

type IMap interface {
	Get(string) (int, bool)
	Set(string, int)
}

func topHash(i uint64) uint8 {
	return uint8(i >> 56)
}

func hash(s string) uint64 {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64()
}

func MakeMap(hint int) IMap {
	bucket := make([]*bmap, 1<<hint, 1<<hint)
	return &_map{
		B:       uint8(hint),
		hash0:   0,
		buckets: bucket,
	}
}

type _map struct {
	B       uint8
	hash0   uint32 // hash seed
	buckets []*bmap
}

type pair struct {
	tophash uint8
	key     string
	val     int
}

type bmap struct {
	tophash [8]*pair
	next    *bmap
}

func (m *_map) Get(key string) (int, bool) {
	hashVal := hash(key)
	mod := int(hashVal) % len(m.buckets)
	b := m.buckets[mod]

	topHash := topHash(hashVal)
	for b != nil {
		for i := 0; i < 8; i++ {
			if b.tophash[i] == nil || b.tophash[i].tophash != topHash {
				continue
			}

			if b.tophash[i].key == key {
				return b.tophash[i].val, true
			}
		}

		b = b.next
	}

	return 0, false
}

func (m *_map) Set(string, int) {

}
