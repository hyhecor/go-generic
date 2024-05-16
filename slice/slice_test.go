package slice

import (
	"testing"

	"github.com/hyhecor/go-generic/hof"
	"github.com/stretchr/testify/assert"
)

func Test_Cdr(t *testing.T) {

	fold := hof.Fold(func(a, b int) int { return a + b }, 0)

	aa := Wrap([]int{1, 2, 3})

	aa, ok := aa.Cdr()

	assert.Equal(t, fold([]int{2, 3}), fold(aa))
	assert.Equal(t, true, ok)

	aa, ok = aa.Cdr()

	assert.Equal(t, fold([]int{3}), fold(aa))
	assert.Equal(t, true, ok)

	aa, ok = aa.Cdr()

	assert.Equal(t, fold([]int{}), fold(aa))
	assert.Equal(t, true, ok)

	aa, ok = aa.Cdr()

	assert.Equal(t, fold([]int{}), fold(aa))
	assert.Equal(t, false, ok)
}

func Test_Cdr_loop(t *testing.T) {

	for aa, ok := Wrap([]int{1, 2, 3}).Cdr(); ok; aa, ok = aa.Cdr() {

		switch len(aa) {
		case 2:
			assert.Equal(t, true, ok)
		case 1:
			assert.Equal(t, true, ok)
		case 0:
			assert.Equal(t, true, ok)
		}
	}
}
