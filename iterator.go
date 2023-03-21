package iterator

import (
	"fmt"

	"github.com/ngicks/genericcontainer/list"
	"github.com/ngicks/iterator/adapter"
	"github.com/ngicks/iterator/base"
	"github.com/ngicks/iterator/def"
)

type Iterator[T any] struct {
	def.SeIterator[T]
}

func FromSlice[T any](sl []T) Iterator[T] {
	return Iterator[T]{
		SeIterator: base.NewSliceIterDe(sl),
	}
}

func FromMap[T comparable, U any](m map[T]U, keySortOption func(keys []T) []T) Iterator[adapter.Zipped[T, U]] {
	return Iterator[adapter.Zipped[T, U]]{
		SeIterator: base.NewMapIterDe(m, keySortOption),
	}
}

func FromFixedList[T any](l *list.List[T]) Iterator[T] {
	return Iterator[T]{
		SeIterator: base.NewListIterDe(l),
	}
}

func FromList[T any](l *list.List[T]) Iterator[T] {
	return Iterator[T]{
		SeIterator: base.NewListIterSe(l),
	}
}

func FromChannel[T any](channel <-chan T) Iterator[T] {
	return Iterator[T]{
		SeIterator: base.NewChanIter(channel),
	}
}

func FromRange(start, end int) Iterator[int] {
	return Iterator[int]{
		SeIterator: base.NewRange(start, end),
	}
}

func (iter Iterator[T]) Select(selector func(T) bool) Iterator[T] {
	return Iterator[T]{adapter.NewSelector[T](iter, selector)}
}
func (iter Iterator[T]) Exclude(excluder func(T) bool) Iterator[T] {
	return Iterator[T]{adapter.NewExcluder[T](iter, excluder)}
}
func (iter Iterator[T]) SkipN(n int) Iterator[T] {
	return Iterator[T]{adapter.NewNSkipper[T](iter, n)}
}
func (iter Iterator[T]) SkipWhile(skipIf func(T) bool) Iterator[T] {
	return Iterator[T]{adapter.NewWhileSkipper[T](iter, skipIf)}
}
func (iter Iterator[T]) TakeN(n int) Iterator[T] {
	return Iterator[T]{adapter.NewNTaker[T](iter, n)}
}
func (iter Iterator[T]) TakeWhile(takeIf func(T) bool) Iterator[T] {
	return Iterator[T]{adapter.NewWhileTaker[T](iter, takeIf)}
}
func (iter Iterator[T]) Chain(z def.SeIterator[T]) Iterator[T] {
	return Iterator[T]{adapter.NewChainer(iter.SeIterator, z)}
}
func (iter Iterator[T]) Map(mapper func(T) T) Iterator[T] {
	return Iterator[T]{adapter.NewSameTyMapper(iter.SeIterator, mapper)}
}
func (iter Iterator[T]) SizeHint() int {
	if sizeHinter, ok := iter.SeIterator.(def.SizeHinter); ok {
		return sizeHinter.SizeHint()
	}
	return -1
}

func (iter Iterator[T]) Unwrap() def.SeIterator[T] {
	return iter.SeIterator
}
func (iter Iterator[T]) MustNext() T {
	v, ok := iter.SeIterator.Next()
	if !ok {
		panic("NextMust: failed")
	}
	return v
}
func (iter Iterator[T]) Reverse() (rev Iterator[T], ok bool) {
	reversed, ok := adapter.Reverse(iter.SeIterator)
	if !ok {
		return
	}
	return Iterator[T]{reversed}, true
}
func (iter Iterator[T]) MustReverse() (rev Iterator[T]) {
	rev, ok := iter.Reverse()
	if !ok {
		panic(fmt.Sprintf("MustReverse: failed: %+v", iter))
	}
	return
}
func (iter Iterator[T]) iterateUntil(predicate func(T) (continueIteration bool)) {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if !predicate(next) {
			break
		}
	}
}
func (iter Iterator[T]) ForEach(each func(T)) {
	iter.iterateUntil(func(t T) bool {
		each(t)
		return true
	})
}
func (iter Iterator[T]) Collect() []T {
	var cap int
	if hint := iter.SizeHint(); hint > 0 {
		cap = hint
	}
	collected := make([]T, 0, cap)
	iter.iterateUntil(func(t T) bool {
		collected = append(collected, t)
		return true
	})
	return collected
}
func (iter Iterator[T]) Find(predicate func(T) bool) (v T, found bool) {
	var lastElement T
	iter.iterateUntil(func(t T) bool {
		b := predicate(t)
		if b {
			lastElement = t
			found = true
		}
		return !b
	})
	return lastElement, found
}

func (iter Iterator[T]) Reduce(reducer func(accumulator T, next T) T) T {
	var accum T
	iter.iterateUntil(func(t T) bool {
		accum = reducer(accum, t)
		return true
	})
	return accum
}
