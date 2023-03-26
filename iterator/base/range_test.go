package base_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/go-iterator/iterator"
)

type RangeTestCases struct {
	expected []int
	start    int
	end      int
}

func TestRange(t *testing.T) {
	testCases := []RangeTestCases{
		{
			expected: []int{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
				10, 11, 12, 13, 14, 15, 16, 17,
				18, 19, 20,
			},
			start: 0,
			end:   21,
		},
		{
			expected: []int{5, 6, 7, 8, 9},
			start:    5,
			end:      10,
		},
		{
			expected: []int{-11, -10, -9, -8, -7},
			start:    -11,
			end:      -6,
		},
	}

	for _, testCase := range testCases {
		t.Run(
			fmt.Sprintf("start=%d,end=%d", testCase.start, testCase.end),
			func(t *testing.T) {
				r := iterator.FromRange(testCase.start, testCase.end)
				collected := r.Collect()
				if cmp.Diff(testCase.expected, collected) != "" {
					t.Fatalf("must be equal. expected = %+v, actual = %+v", testCase.expected, collected)
				}
			},
		)
	}

	for _, testCase := range testCases {
		var expected sort.IntSlice = testCase.expected
		sort.Sort(sort.Reverse(expected))
		t.Run(
			fmt.Sprintf("reversed:start=%d,end=%d", testCase.start, testCase.end),
			func(t *testing.T) {
				r := iterator.FromRange(testCase.start, testCase.end).MustReverse()
				collected := r.Collect()
				if cmp.Diff([]int(expected), collected) != "" {
					t.Fatalf("must be equal. expected = %+v, actual = %+v", expected, collected)
				}
			},
		)
	}
}
