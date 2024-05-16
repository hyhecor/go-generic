package hof

// Fold
//
//	(b -> a -> b) -> b -> [a] -> b
func Fold[A any, B any](reducer func(b B, a A) (b_ B), b B) func(aa []A) (b_ B) {
	return func(aa []A) (b_ B) {

		b_ = b

		for i := range aa {
			b_ = reducer(b_, aa[i])
		}
		return
	}
}
