package adapter_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/go-iterator/iterator"
)

type SelectorTestCase[T any] struct {
	input    []T
	expected []T
	filter   func(T) bool
}

func executeCases[T any](
	t *testing.T,
	method func(iter iterator.Iterator[T], filter func(v T) bool) iterator.Iterator[T],
	testCases []SelectorTestCase[T],
) {
	for _, testCase := range testCases {
		iter := iterator.FromSlice(testCase.input)
		iter = method(iter, testCase.filter)
		actual := iter.Collect()
		if cmp.Diff(testCase.expected, actual) != "" {
			t.Fatalf("not equal: expected = %+v, actual = %+v", testCase.expected, actual)
		}
	}
}

func TestExcluder(t *testing.T) {
	testCases := []SelectorTestCase[int]{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 3, 5},
			filter:   func(i int) bool { return i%2 == 0 },
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2},
			filter:   func(i int) bool { return i >= 3 },
		},
	}

	executeCases(
		t,
		func(iter iterator.Iterator[int], filter func(v int) bool) iterator.Iterator[int] {
			return iter.Exclude(filter)
		},
		testCases,
	)
}

func TestSelector(t *testing.T) {
	testCases := []SelectorTestCase[int]{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{2, 4},
			filter:   func(i int) bool { return i%2 == 0 },
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{3, 4, 5},
			filter:   func(i int) bool { return i >= 3 },
		},
	}

	executeCases(
		t,
		func(iter iterator.Iterator[int], filter func(v int) bool) iterator.Iterator[int] {
			return iter.Select(filter)
		},
		testCases,
	)
}
