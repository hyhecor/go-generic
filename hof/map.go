package hof

// Map
//
//	(a -> b) -> [a] -> [b]
func Map[A any, B any](mapper func(a A) (b B)) func(aa []A) (bb []B) {
	return func(aa []A) (bb []B) {

		bb = make([]B, len(aa))

		for i := range aa {
			bb[i] = mapper(aa[i])
		}

		return
	}
}
