package adapter_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/genericcontainer/list"
	"github.com/ngicks/go-iterator/iterator"
)

type takeNTestCase struct {
	input        []int
	takeN        int
	expectedSize int
	expected     []int
}

func TestTakeNSizeHint(t *testing.T) {
	testCases := []takeNTestCase{
		{
			input:        []int{1, 2, 3},
			takeN:        3,
			expectedSize: 3,
		},
		{
			input:        []int{3, 5, 6, 76, 1, 34, 8},
			takeN:        5,
			expectedSize: 5,
		},
		{
			input:        []int{3, 5, 6},
			takeN:        120,
			expectedSize: 3,
		},
	}

	for _, testCase := range testCases {
		iter := iterator.FromSlice(testCase.input).TakeN(testCase.takeN)
		if size := iter.SizeHint(); size != testCase.expectedSize {
			t.Fatalf("mismatched: expected = %d, actual = %d", testCase.expectedSize, size)
		}
	}

	iter := iterator.FromList(list.New[int]()).TakeN(5)
	if size := iter.SizeHint(); size != -1 {
		t.Fatalf("mismatched: expected = %d, actual = %d", -1, size)
	}
}

func TestTakeN(t *testing.T) {
	testCases := []takeNTestCase{
		{
			input:    []int{1, 2, 3},
			takeN:    3,
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{1, 2, 3},
			takeN:    0,
			expected: []int{},
		},
		{
			input:    []int{3, 5, 6, 76, 1, 34, 8},
			takeN:    5,
			expected: []int{3, 5, 6, 76, 1},
		},
		{
			input:    []int{3, 5, 6},
			takeN:    120,
			expected: []int{3, 5, 6},
		},
	}

	for _, testCase := range testCases {
		iter := iterator.FromSlice(testCase.input).TakeN(testCase.takeN)
		if collected := iter.Collect(); cmp.Diff(testCase.expected, collected) != "" {
			t.Errorf("mismatched: expected = %+v, actual = %+v", testCase.expected, collected)
		}
	}
}

func TestTakeWhileSizeHint(t *testing.T) {
	for _, iter := range []iterator.Iterator[int]{
		iterator.FromList(list.New[int]()),
		iterator.FromSlice([]int{1, 2, 3}),
		iterator.FromRange(0, 25),
	} {
		size := iter.TakeWhile(func(i int) bool { return true }).SizeHint()
		if size != -1 {
			t.Fatalf("mismatched: expected = %d, actual = %d", -1, size)
		}
	}
}

type takeWhileTestCase struct {
	input    []int
	takeIf   func(int) bool
	expected []int
}

func TestTakeWhile(t *testing.T) {
	testCases := []takeWhileTestCase{
		{
			input:    []int{1, 2, 3},
			takeIf:   func(i int) bool { return true },
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{1, 2, 3},
			takeIf:   func(i int) bool { return false },
			expected: []int{},
		},
		{
			input:    []int{3, 5, 6, 76, 1, 34, 8},
			takeIf:   func(i int) bool { return i < 15 },
			expected: []int{3, 5, 6},
		},
		{
			input:    []int{3, 5, 6},
			takeIf:   func(i int) bool { return i%2 != 0 },
			expected: []int{3, 5},
		},
	}

	for _, testCase := range testCases {
		{
			iter := iterator.FromSlice(testCase.input).TakeWhile(func(i int) bool { return true })
			if collected := iter.Collect(); cmp.Diff(testCase.input, collected) != "" {
				t.Errorf("mismatched: expected = %+v, actual = %+v", testCase.input, collected)
			}
		}
		{
			iter := iterator.FromSlice(testCase.input).TakeWhile(testCase.takeIf)
			if collected := iter.Collect(); cmp.Diff(testCase.expected, collected) != "" {
				t.Errorf("mismatched: expected = %+v, actual = %+v", testCase.expected, collected)
			}
		}
	}
}
