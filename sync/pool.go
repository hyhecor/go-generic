package sync

import (
	"sync"
)

type pool[A any] struct {
	*sync.Pool
}

func NewPool[A any](fn func() A) *pool[A] {
	return &pool[A]{
		Pool: &sync.Pool{
			New: func() any { return fn() },
		},
	}
}

func (p *pool[A]) Get() A {
	return p.Pool.Get().(A)
}

func (p *pool[A]) Put(a A) {
	p.Pool.Put(a)
}
