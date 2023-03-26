package adapter

import "github.com/ngicks/go-iterator/iterator/def"

//methodgen:ignore=sizehint
type Excluder[T any] struct {
	inner    def.SeIterator[T]
	excluder func(T) bool
}

func NewExcluder[T any](iter def.SeIterator[T], excluder func(T) bool) Excluder[T] {
	return Excluder[T]{
		inner:    iter,
		excluder: excluder,
	}
}

func (e Excluder[T]) Next() (next T, ok bool) {
	var v T
	for {
		v, ok = e.inner.Next()
		if !ok {
			return
		}
		if e.excluder(v) {
			continue
		}
		return v, ok
	}
}

//methodgen:ignore=sizehint
type Selector[T any] struct {
	inner    def.SeIterator[T]
	selector func(T) bool
}

func NewSelector[T any](iter def.SeIterator[T], selector func(T) bool) Selector[T] {
	return Selector[T]{
		inner:    iter,
		selector: selector,
	}
}

func (s Selector[T]) Next() (next T, ok bool) {
	var v T
	for {
		v, ok = s.inner.Next()
		if !ok {
			return
		}
		if s.selector(v) {
			return v, ok
		}
	}
}
