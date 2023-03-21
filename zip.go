package iterator

import (
	"github.com/ngicks/iterator/adapter"
	"github.com/ngicks/iterator/def"
)

func Zip[T any, U any](iterFormer def.SeIterator[T], iterLatter def.SeIterator[U]) Iterator[adapter.Zipped[T, U]] {
	return Iterator[adapter.Zipped[T, U]]{
		SeIterator: adapter.NewZipper(iterFormer, iterLatter),
	}
}
