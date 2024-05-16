package slice

type Slice[A any] []A

func New[A any](aa []A) Slice[A] {
	return aa
}

func (sl Slice[A]) Slice() []A {
	return sl
}

func (sl Slice[A]) Len() int {
	return len(sl)
}

func (sl Slice[A]) Clone() []A {

	aa := make([]A, len(sl))

	copy(aa, sl)

	return aa
}

func (sl Slice[A]) Car() (a A, ok bool) {

	if len(sl) == 0 {
		return a, false
	}

	return sl[0], true
}

func (sl Slice[A]) Cdr() (aa Slice[A], ok bool) {

	if len(sl) < 1 {
		return []A{}, false
	}

	return sl[1:], true
}
