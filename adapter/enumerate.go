package adapter

import "github.com/ngicks/iterator/def"

type Enumerated[T any] struct {
	Count int
	Next  T
}

//methodgen:ignore=reverse
type Enumerator[T any] struct {
	count int
	inner def.SeIterator[T]
}

func NewEnumerator[T any](iter def.SeIterator[T]) *Enumerator[T] {
	return &Enumerator[T]{
		inner: iter,
	}
}

func (e *Enumerator[T]) Next() (next Enumerated[T], ok bool) {
	nextInner, ok := e.inner.Next()
	if !ok {
		return Enumerated[T]{}, false
	}
	c := e.count
	e.count++
	return Enumerated[T]{Count: c, Next: nextInner}, true
}
