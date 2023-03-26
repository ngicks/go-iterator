package iterator

import (
	"github.com/ngicks/go-iterator/iterator/adapter"
	"github.com/ngicks/go-iterator/iterator/def"
)

func Enumerate[T any](iter def.SeIterator[T]) Iterator[adapter.Enumerated[T]] {
	return Iterator[adapter.Enumerated[T]]{
		SeIterator: adapter.NewEnumerator(iter),
	}
}
