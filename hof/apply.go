package hof

// Apply
//
//	([a] -> b) -> [a] -> b
func Apply[A any, B any](fn func(a []A) B, aa ...A) B {
	return fn(aa)
}

// Apply with no return
//
//	([a] -> b) -> [a] -> none
func ApplyN[A any](fn func(a []A), aa ...A) {
	fn(aa)
}
