package hof

func Hash[A any, B any, C any](fn func(A) (B, C)) func(A) func(*B) func(*C) {
	return func(a A) func(*B) func(*C) {
		return func(b *B) func(*C) {
			return func(c *C) {
				*b, *c = fn(a)
			}
		}
	}
}
