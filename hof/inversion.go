package hof

// Inversion
//
//	(a -> (b, c)) -> (*b, *c) -> (a)
func Inversion[A any, B any, C any](fn func(A) (B, C)) func(b *B, c *C) func(a A) {
	return func(b *B, c *C) func(a A) {
		return func(a A) {
			*b, *c = fn(a)
		}
	}
}
