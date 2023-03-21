package adapter_test

import (
	"testing"

	"github.com/ngicks/iterator"
	"github.com/ngicks/iterator/adapter"
	"github.com/ngicks/iterator/base"
)

func TestSizeHinter(t *testing.T) {
	innerSl := []int{1, 2, 3, 4, 5}
	sliceIter := base.NewSliceIterDe(innerSl)
	iter := iterator.Iterator[int]{SeIterator: sliceIter}

	assertEq := func(len int) {
		if iter.SizeHint() != len {
			t.Fatalf("Len incorrect: expected = %d, actual  %d", len, iter.SizeHint())
		}
	}

	assertEq(5)

	iter = iterator.Iterator[int]{
		SeIterator: adapter.NewExcluder[int](
			sliceIter,
			func(i int) bool { return i%3 == 0 },
		),
	}

	assertEq(5)
	iter = iterator.Iterator[int]{
		SeIterator: adapter.NewSelector[int](
			sliceIter,
			func(i int) bool { return i%3 == 0 },
		),
	}
	assertEq(5)
	iter = iterator.Iterator[int]{
		SeIterator: iterator.Iterator[int]{SeIterator: sliceIter},
	}
	assertEq(5)
	iter = iterator.Iterator[int]{
		SeIterator: adapter.NewMapper[int](
			sliceIter,
			func(i int) int { return i },
		),
	}
	assertEq(5)
	iter = iterator.Iterator[int]{SeIterator: adapter.ReversedDeIter[int]{DeIterator: sliceIter}}
	assertEq(5)
}
