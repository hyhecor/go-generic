package slice

type slice[A any] []A

func Wrap[S ~[]A, A any](aa S) slice[A] {
	return []A(aa)
}

func (sl slice[A]) Slice() []A {
	return sl
}

func (sl slice[A]) Len() int {
	return len(sl)
}

func (sl slice[A]) Clone() []A {

	aa := make([]A, len(sl))

	copy(aa, sl)

	return aa
}

func (sl slice[A]) Car() (a A, ok bool) {

	if len(sl) == 0 {
		return a, false
	}

	return sl[0], true
}

func (sl slice[A]) Cdr() (aa slice[A], ok bool) {

	if len(sl) < 1 {
		return []A{}, false
	}

	return sl[1:], true
}
