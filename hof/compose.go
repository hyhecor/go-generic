package hof

// Compose
//
//	(a -> b) -> (b -> c) -> a -> c
func Compose[A any, B any, C any](ab func(A) B, bc func(B) C) func(A) C {
	return func(a A) (c C) {
		return bc(ab(a))
	}
}

// a -> a
func Pure[A any](a A) A {
	return a
}
