package adapter

import (
	"fmt"

	"github.com/ngicks/go-iterator/iterator/def"
)

func Reverse[T any](iter def.SeIterator[T]) (rev def.SeIterator[T], ok bool) {
	switch x := iter.(type) {
	case def.Reverser[T]:
		return x.Reverse()
	case def.DeIterator[T]:
		return ReversedDeIter[T]{DeIterator: x}, true
	case def.Unwrapper[T]:
		return Reverse(x.Unwrap())
	}
	return nil, false
}

func MustReverse[T any](iter def.SeIterator[T]) (rev def.SeIterator[T]) {
	rev, ok := Reverse(iter)
	if !ok {
		panic(fmt.Sprintf("MustReverse: failed: %+v", iter))
	}
	return
}

//methodgen:ignore=reverse
type ReversedDeIter[T any] struct {
	def.DeIterator[T]
}

func (rev ReversedDeIter[T]) Next() (next T, ok bool) {
	return rev.DeIterator.NextBack()
}
func (rev ReversedDeIter[T]) NextBack() (next T, ok bool) {
	return rev.DeIterator.Next()
}

// Reverse implements Reverser[T].
// This simply unwrap iterator.
func (iter ReversedDeIter[T]) Reverse() (rev def.SeIterator[T], ok bool) {
	return iter.DeIterator, true
}
