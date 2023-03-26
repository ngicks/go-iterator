package iterator_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/go-iterator/iterator"
	"golang.org/x/exp/constraints"
)

func testIterator[T comparable](t *testing.T, input []T) {
	t.Run("Next", func(t *testing.T) {
		iter := iterator.FromSlice(input)
		nexted := []T{}
		for {
			next, ok := iter.Next()
			if !ok {
				break
			}
			nexted = append(nexted, next)
		}
		if cmp.Diff(input, nexted) != "" {
			t.Fatalf("must equal. actual = %+v, expected = %+v", nexted, input)
		}
	})

	t.Run("Collect", func(t *testing.T) {
		iter := iterator.FromSlice(input)
		collected := iter.Collect()
		if cmp.Diff(input, collected) != "" {
			t.Fatalf("must equal. actual = %+v, expected = %+v", collected, input)
		}
	})

	t.Run("ForEach", func(t *testing.T) {
		iter := iterator.FromSlice(input)
		eached := []T{}
		iter.ForEach(func(i T) {
			eached = append(eached, i)
		})
		if cmp.Diff(input, eached) != "" {
			t.Fatalf("must equal. actual = %+v, expected = %+v", eached, input)
		}
	})

	t.Run("Map", func(t *testing.T) {
		iter := iterator.FromSlice(input)
		iter = iter.Map(func(t T) T {
			var zero T
			return zero
		})
		collected := iter.Collect()
		for _, v := range collected {
			var zero T
			if v != zero {
				t.Fatalf("must equal. actual = %+v, expected = %+v", v, zero)
			}
		}
	})
}

func TestIterator(t *testing.T) {
	input := []int{5, 6, 7, 12, 8, 9, 10}

	t.Run("Find", func(t *testing.T) {
		iter := iterator.FromSlice(input)
		v, _ := iter.Find(func(i int) bool { return i > 6 && i%2 == 0 })
		if v != 12 {
			t.Fatalf("invalid find: %d", v)
		}
	})
	t.Run("Find no match", func(t *testing.T) {
		iter := iterator.FromSlice(input)
		_, found := iter.Find(func(i int) bool { return i > 125 })
		if found {
			t.Fatalf("invalid find")
		}
	})
	t.Run("Reduce", func(t *testing.T) {
		iter := iterator.FromSlice(input)
		reduced := iter.Reduce(max[int])
		if reduced != 12 {
			t.Fatalf("invalid reduce: %d", reduced)
		}
	})
}

func max[T constraints.Ordered](i, j T) T {
	if i > j {
		return i
	} else {
		return j
	}
}
