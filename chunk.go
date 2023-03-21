package iterator

import "github.com/ngicks/iterator/adapter"

func Chunks[T any](sl []T, size uint) Iterator[[]T] {
	return Iterator[[]T]{
		SeIterator: adapter.NewChunker(sl, size),
	}
}
