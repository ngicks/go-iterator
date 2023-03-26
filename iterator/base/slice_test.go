package base_test

import (
	"testing"

	"github.com/ngicks/go-iterator/iterator/base"
)

func TestSliceIterator(t *testing.T) {
	t.Run("input is nil", func(t *testing.T) {
		iter := base.NewSliceIterDe[int](nil)

		if iter.SizeHint() != 0 {
			t.Fatal("mismatched")
		}

		for i := 0; i < 100; i++ {
			if _, ok := iter.Next(); ok {
				t.Fatalf("must not be ok == true")
			}
		}

		for i := 0; i < 100; i++ {
			if _, ok := iter.NextBack(); ok {
				t.Fatalf("must not be ok == true")
			}
		}
	})
	t.Run("len = 0", func(t *testing.T) {
		iter := base.NewSliceIterDe([]int{})

		if iter.SizeHint() != 0 {
			t.Fatal("mismatched")
		}

		for i := 0; i < 100; i++ {
			if _, ok := iter.Next(); ok {
				t.Fatalf("must not be ok == true")
			}
		}

		for i := 0; i < 100; i++ {
			if _, ok := iter.NextBack(); ok {
				t.Fatalf("must not be ok == true")
			}
		}
	})
	t.Run("normal", func(t *testing.T) {
		sl := []int{}
		for i := 0; i < 25; i = i + 5 {
			sl = append(sl, i)
		}

		iter := base.NewSliceIterDe(sl)

		if iter.SizeHint() != 5 {
			t.Fatalf("mismatched. size = %d", iter.SizeHint())
		}

		if next, ok := iter.Next(); !(next == 0 && ok) {
			t.Fatalf("mismatched. next = %d, ok = %t", next, ok)
		}
		if iter.SizeHint() != 4 {
			t.Fatalf("mismatched = %d", iter.SizeHint())
		}

		if next, ok := iter.NextBack(); !(next == 20 && ok) {
			t.Fatalf("mismatched. next = %d, ok = %t", next, ok)
		}
		if iter.SizeHint() != 3 {
			t.Fatalf("mismatched = %d", iter.SizeHint())
		}

		if next, ok := iter.NextBack(); !(next == 15 && ok) {
			t.Fatalf("mismatched. next = %d, ok = %t", next, ok)
		}
		if iter.SizeHint() != 2 {
			t.Fatalf("mismatched = %d", iter.SizeHint())
		}

		if next, ok := iter.Next(); !(next == 5 && ok) {
			t.Fatalf("mismatched. next = %d, ok = %t", next, ok)
		}
		if iter.SizeHint() != 1 {
			t.Fatalf("mismatched = %d", iter.SizeHint())
		}

		if next, ok := iter.Next(); !(next == 10 && ok) {
			t.Fatalf("mismatched. next = %d, ok = %t", next, ok)
		}
		if iter.SizeHint() != 0 {
			t.Fatalf("mismatched = %d", iter.SizeHint())
		}

		for i := 0; i < 100; i++ {
			if _, ok := iter.Next(); ok {
				t.Fatalf("must not be ok == true")
			}
		}
	})
}
