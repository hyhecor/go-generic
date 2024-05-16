package hof

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pair(t *testing.T) {

	w := Pair(123, "123")

	assert.Equal(t, 123, w.First())
	assert.Equal(t, "123", w.Second())

}

func Benchmark_Pair(b *testing.B) {

	b.Run("Pair", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			_ = Pair(123, "123")
		}
	})

}
