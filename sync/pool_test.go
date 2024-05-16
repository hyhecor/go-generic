package sync

import (
	"crypto/sha1"
	"hash"
	"hash/crc32"
	"sync"
	"testing"
)

func Test_Pool(t *testing.T) {
	var pool = NewPool(func() hash.Hash { return sha1.New() })

	v := pool.Get()

	pool.Put(v)
}

func Benchmark_Pool(b *testing.B) {

	b.Run("sync.Pool (base)", func(b *testing.B) {

		var pool = sync.Pool{New: func() any { return crc32.NewIEEE() }}

		for i := 0; i < b.N; i++ {

			h := pool.Get().(hash.Hash32)
			h.Sum([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
			pool.Put(h)
		}
	})

	// b.Run("Pool", func(b *testing.B) {

	// 	var pool = pool[hash.Hash32]{
	// 		getter: func() hash.Hash32 {
	// 			return crc32.NewIEEE()
	// 		},
	// 	}

	// 	for i := 0; i < b.N; i++ {

	// 		h := pool.Get()
	// 		h.Sum([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	// 		pool.Put(h)
	// 	}
	// })

	b.Run("Pool", func(b *testing.B) {

		var pool = NewPool(crc32.NewIEEE)

		for i := 0; i < b.N; i++ {

			h := pool.Get()
			h.Sum([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
			pool.Put(h)
		}
	})

}
