package iterator

import "github.com/ngicks/iterator/adapter"

func Windows[T any](sl []T, width uint) Iterator[[]T] {
	return Iterator[[]T]{
		SeIterator: adapter.NewWindower(sl, width),
	}
}
