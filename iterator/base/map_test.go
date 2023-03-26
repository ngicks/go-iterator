package base_test

import (
	"sort"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/go-iterator/iterator/adapter"
	"github.com/ngicks/go-iterator/iterator/base"
	"github.com/stretchr/testify/require"
)

func TestMapIterator(t *testing.T) {
	t.Run("input is nil", func(t *testing.T) {
		require := require.New(t)

		iter := base.NewMapIterDe[string, int](nil, nil)

		require.Equal(0, iter.SizeHint())

		for i := 0; i < 100; i++ {
			_, ok := iter.Next()
			require.False(ok)
		}

		for i := 0; i < 100; i++ {
			_, ok := iter.NextBack()
			require.False(ok)
		}
	})
	t.Run("len = 0", func(t *testing.T) {
		require := require.New(t)

		iter := base.NewMapIterDe(map[string]int{}, nil)

		require.Equal(0, iter.SizeHint())

		for i := 0; i < 100; i++ {
			_, ok := iter.Next()
			require.False(ok)
		}

		for i := 0; i < 100; i++ {
			_, ok := iter.NextBack()
			require.False(ok)
		}
	})
	t.Run("random order", func(t *testing.T) {
		require := require.New(t)

		mmap := map[string]int{}
		values := make([]int, 0)
		for i := 0; i < 25; i = i + 5 {
			values = append(values, i)
			mmap[strconv.FormatInt(int64(i), 10)] = i
		}

		iter := base.NewMapIterDe(mmap, nil)

		require.Equal(5, iter.SizeHint())
		yielded := make([]int, 0)
		for i := 0; i < 5; i++ {
			val, ok := iter.Next()
			require.True(ok)
			require.Equal(val.Former, strconv.FormatInt(int64(val.Latter), 10))
			yielded = append(yielded, val.Latter)
		}

		sortInt := func(x []int) {
			sort.Slice(x, func(i, j int) bool {
				return x[i] < x[j]
			})
		}
		sortInt(yielded)

		require.Conditionf(
			func() bool { return cmp.Equal(yielded, values) },
			"diff = %s",
			cmp.Diff(yielded, values),
		)
		// not necessary to be same order.
	})

	t.Run("int asc code order", func(t *testing.T) {
		require := require.New(t)

		mmap := map[string]int{}
		for i := 0; i < 25; i = i + 5 {
			mmap[strconv.FormatInt(int64(i), 10)] = i
		}

		iter := base.NewMapIterDe(
			mmap,
			func(keys []string) []string {
				sort.Slice(keys, func(i, j int) bool {
					a, _ := strconv.ParseInt(keys[i], 10, 64)
					b, _ := strconv.ParseInt(keys[j], 10, 64)
					return a < b
				})
				return keys
			},
		)

		require.Equal(5, iter.SizeHint())

		ans := []int{0, 20, 5, 10, 15}
		yielded := make([]int, 0)

		var val adapter.Zipped[string, int]
		var ok bool
		checkVal := func() {
			require.True(ok)
			require.Equal(val.Former, strconv.FormatInt(int64(val.Latter), 10))
		}

		val, ok = iter.Next()
		checkVal()
		yielded = append(yielded, val.Latter)

		val, ok = iter.NextBack()
		checkVal()
		yielded = append(yielded, val.Latter)

		val, ok = iter.Next()
		checkVal()
		yielded = append(yielded, val.Latter)

		val, ok = iter.Next()
		checkVal()
		yielded = append(yielded, val.Latter)

		val, ok = iter.Next()
		checkVal()
		yielded = append(yielded, val.Latter)

		require.Conditionf(
			func() bool { return cmp.Equal(yielded, ans) },
			"diff = %s",
			cmp.Diff(yielded, ans))
	})
}
