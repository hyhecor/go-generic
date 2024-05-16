package hof

// Filter
//
//	(a -> bool) -> [a] -> [a]
func Filter[A any, B bool](filter func(a A) (b B)) func(aa []A) (aa_ []A) {
	return func(aa []A) (aa_ []A) {

		aa_ = make([]A, 0, len(aa))

		for i := range aa {

			if filter(aa[i]) {
				aa_ = append(aa_, aa[i])
			}
		}
		return
	}
}
