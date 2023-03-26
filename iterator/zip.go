package iterator

import (
	"github.com/ngicks/go-iterator/iterator/adapter"
	"github.com/ngicks/go-iterator/iterator/def"
)

func Zip[T any, U any](iterFormer def.SeIterator[T], iterLatter def.SeIterator[U]) Iterator[adapter.Zipped[T, U]] {
	return Iterator[adapter.Zipped[T, U]]{
		SeIterator: adapter.NewZipper(iterFormer, iterLatter),
	}
}
