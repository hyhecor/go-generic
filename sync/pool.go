package sync

import (
	"sync"
)

type Pool[A any] struct {
	once sync.Once
	pool sync.Pool

	New func() A
}

func (p *Pool[A]) Get() A {
	p.once.Do(func() {
		p.pool.New = func() any {
			return p.New()
		}
	})

	return p.pool.Get().(A)
}

func (p *Pool[A]) Put(a A) {
	p.pool.Put(a)
}
