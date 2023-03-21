package adapter_test

import (
	"strconv"
	"testing"

	"github.com/ngicks/iterator"
	"github.com/ngicks/iterator/adapter"
	"github.com/ngicks/iterator/base"
)

func TestZip(t *testing.T) {
	t.Run("different type", func(t *testing.T) {
		zip := iterator.Zip[int, string](
			iterator.FromRange(0, 6),
			adapter.NewMapper[int](
				base.NewRange(0, 3),
				func(i int) string { return strconv.Itoa(i) },
			),
		)

		for i := 0; i < 3; i++ {
			v, ok := zip.Next()
			if !ok {
				t.Fatal("must be ok = true")
			}
			if v.Former != i {
				t.Fatalf("mismatched. expected = %d, actual = %d", i, v.Former)
			}
			if v.Latter != strconv.Itoa(i) {
				t.Fatalf("mismatched. expected = %s, actual = %s", strconv.Itoa(i), v.Latter)
			}
		}

		iterator.FromRange(0, 100).ForEach(func(int) {
			_, ok := zip.Next()
			if ok {
				t.Fatal("must be ok = false")
			}
		})
	})

	t.Run("iter is exhausted when either is", func(t *testing.T) {
		for _, testCase := range []struct {
			left  int
			right int
		}{{2, 3}, {3, 6}, {8, 13}, {26, 33}, {2789, 23}} {
			zip := iterator.Zip[int, int](
				iterator.FromRange(0, testCase.left),
				iterator.FromRange(0, testCase.right),
			)

			less := func() int {
				if testCase.left < testCase.right {
					return testCase.left
				} else {
					return testCase.right
				}
			}()

			if size := zip.SizeHint(); size != less {
				t.Fatalf("mismatched. expected = %d, actual = %d", less, size)
			}

			for i := 0; i < less; i++ {
				_, ok := zip.Next()
				if !ok {
					t.Fatal("must be ok = true")
				}
				if size := zip.SizeHint(); size != (less - (i + 1)) {
					t.Fatalf("mismatched. expected = %d, actual = %d", less-(i+1), size)
				}
			}

			iterator.FromRange(0, 100).ForEach(func(int) {
				_, ok := zip.Next()
				if ok {
					t.Fatal("must be ok = false")
				}
			})
		}
	})

	t.Run("Reverse", func(t *testing.T) {
		for _, testCase := range []struct {
			left  int
			right int
		}{{2, 3}, {3, 6}, {8, 13}, {26, 33}, {2789, 23}} {
			zip := iterator.Zip[int, int](
				iterator.FromRange(0, testCase.left),
				iterator.FromRange(0, testCase.right),
			)

			if _, ok := zip.Reverse(); ok {
				t.Fatal("must not be ok = true")
			}
		}

		for _, testCase := range []struct {
			left  int
			right int
		}{{2, 2}, {3, 3}, {15, 15}, {2789, 2789}} {
			zip := iterator.Zip[int, int](
				iterator.FromRange(0, testCase.left),
				iterator.FromRange(0, testCase.right).Map(func(i int) int { return i * 2 }),
			)

			rev, ok := iterator.Iterator[adapter.Zipped[int, int]]{SeIterator: zip}.Reverse()
			if !ok {
				t.Fatal("must not be ok = false")
			}

			size := rev.SizeHint()
			if size != testCase.right {
				t.Fatalf("mismatched. expected = %d, actual = %d", testCase.right, size)
			}

			for i := 0; i < testCase.right; i++ {
				next, ok := rev.Next()
				if !ok {
					t.Fatal("must not be ok = false")
				}
				if next.Latter != next.Former*2 {
					t.Fatal("mismatched")
				}
			}

			iterator.FromRange(0, 100).ForEach(func(int) {
				_, ok := zip.Next()
				if ok {
					t.Fatal("must be ok = false")
				}
			})
		}
	})

}
