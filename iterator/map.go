package iterator

import (
	"github.com/ngicks/go-iterator/iterator/adapter"
	"github.com/ngicks/go-iterator/iterator/def"
)

func Map[T, U any](iter def.SeIterator[T], mapper func(T) U) Iterator[U] {
	return Iterator[U]{
		SeIterator: adapter.NewMapper(iter, mapper),
	}
}
