package hof

type pair[A any, B any] struct {
	first  A
	second B
}

func (p pair[A, B]) First() A {
	return p.first
}

func (p pair[A, B]) Second() B {
	return p.second
}

func (p pair[A, B]) Set() (A, B) {
	return p.first, p.second
}

// Pair
//
//	(b, c) -> {b, c}
func Pair[A any, B any](a A, b B) pair[A, B] {

	return pair[A, B]{
		first:  a,
		second: b,
	}
}
