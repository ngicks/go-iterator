package iterator

import (
	"github.com/ngicks/iterator/adapter"
	"github.com/ngicks/iterator/def"
)

func Enumerate[T any](iter def.SeIterator[T]) Iterator[adapter.Enumerated[T]] {
	return Iterator[adapter.Enumerated[T]]{
		SeIterator: adapter.NewEnumerator(iter),
	}
}
