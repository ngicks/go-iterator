package adapter

import "github.com/ngicks/iterator/def"

// Mapper applies mapper function.
type Mapper[T, U any] struct {
	inner  def.SeIterator[T]
	mapper func(T) U
}

func NewMapper[T, U any](iter def.SeIterator[T], mapper func(T) U) Mapper[T, U] {
	return Mapper[T, U]{
		inner:  iter,
		mapper: mapper,
	}
}

func (m Mapper[T, U]) Next() (next U, ok bool) {
	v, ok := m.inner.Next()
	if ok {
		return m.mapper(v), ok
	}
	return
}

// SameTyMapper applies mapper function that returns value of same type to input.
type SameTyMapper[T any] struct {
	inner  def.SeIterator[T]
	mapper func(T) T
}

func NewSameTyMapper[T any](iter def.SeIterator[T], mapper func(T) T) SameTyMapper[T] {
	return SameTyMapper[T]{
		inner:  iter,
		mapper: mapper,
	}
}

func (m SameTyMapper[T]) Next() (next T, ok bool) {
	v, ok := m.inner.Next()
	if ok {
		return m.mapper(v), ok
	}
	return
}
