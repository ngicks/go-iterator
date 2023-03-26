package adapter

import "github.com/ngicks/go-iterator/iterator/def"

//methodgen:ignore=sizehint
type NTaker[T any] struct {
	inner def.SeIterator[T]
	n     int
}

func NewNTaker[T any](iter def.SeIterator[T], n int) *NTaker[T] {
	return &NTaker[T]{
		inner: iter,
		n:     n,
	}
}

func (t *NTaker[T]) SizeHint() int {
	if hint, ok := t.inner.(def.SizeHinter); ok {
		size := hint.SizeHint()
		if size > t.n {
			return t.n
		} else {
			// maybe -1
			return size
		}
	}
	return -1
}

func (t *NTaker[T]) Next() (next T, ok bool) {
	var v T
	v, ok = t.inner.Next()
	if !ok {
		return
	}
	if t.n > 0 {
		t.n--
		return v, ok
	}
	return v, false
}

//methodgen:ignore=sizehint
type WhileTaker[T any] struct {
	inner        def.SeIterator[T]
	isOutOfWhile bool
	takeIf       func(T) bool
}

func NewWhileTaker[T any](iter def.SeIterator[T], takeIf func(T) bool) *WhileTaker[T] {
	return &WhileTaker[T]{
		inner:  iter,
		takeIf: takeIf,
	}
}

func (t WhileTaker[T]) SizeHint() int {
	return -1
}

func (t *WhileTaker[T]) Next() (next T, ok bool) {
	var v T
	if t.isOutOfWhile {
		return
	}

	v, ok = t.inner.Next()
	if !ok {
		t.isOutOfWhile = true
		return
	}
	if t.takeIf(v) {
		return v, ok
	}
	t.isOutOfWhile = true
	return next, false
}
