package base_test

import (
	"testing"

	"github.com/ngicks/genericcontainer/list"
	"github.com/ngicks/iterator/base"
)

func TestListIteratorDe(t *testing.T) {
	t.Run("len = 0", func(t *testing.T) {
		list := list.New[int]()
		iter := base.NewListIterDe(list)

		if iter.SizeHint() != 0 {
			t.Fatalf("mismatched = %d", iter.SizeHint())
		}

		for i := 0; i < 100; i++ {
			if _, ok := iter.Next(); ok {
				t.Fatalf("must not be ok == true")
			}
		}

		list.PushBack(27)
		list.PushBack(56)
		list.PushBack(32)
		list.PushBack(42)

		for i := 0; i < 100; i++ {
			if _, ok := iter.Next(); ok {
				t.Fatalf("must not be ok == true")
			}
		}
	})

	t.Run("normal case", func(t *testing.T) {
		list := list.New[int]()
		for i := 0; i < 25; i = i + 5 {
			list.PushBack(i)
		}

		iter := base.NewListIterDe(list)

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

		list.PushBack(27)
		list.PushBack(56)
		list.PushBack(32)
		list.PushBack(42)

		for i := 0; i < 100; i++ {
			if _, ok := iter.Next(); ok {
				t.Fatalf("must not be ok == true")
			}
		}
	})
}

func TestListIteratorSe(t *testing.T) {
	t.Run("len = 0", func(t *testing.T) {
		list := list.New[int]()
		iter := base.NewListIterSe(list)

		for i := 0; i < 100; i++ {
			if _, ok := iter.Next(); ok {
				t.Fatalf("must not be ok == true")
			}
		}

		list.PushBack(27)
		list.PushBack(56)
		list.PushBack(32)
		list.PushBack(42)

		for _, v := range []int{27, 56, 32, 42} {
			next, ok := iter.Next()
			if !ok {
				t.Fatalf("must not be ok == false")
			}
			if next != v {
				t.Fatalf("mismatched. expected = %d, actual = %d", v, next)
			}
		}

		for i := 0; i < 100; i++ {
			if _, ok := iter.Next(); ok {
				t.Fatalf("must not be ok == true")
			}
		}
	})

	t.Run("normal case", func(t *testing.T) {
		list := list.New[int]()
		for i := 0; i < 25; i = i + 5 {
			list.PushBack(i)
		}

		iter := base.NewListIterSe(list)

		for i := 0; i < 25; i = i + 5 {
			next, ok := iter.Next()
			if !ok {
				t.Fatal("must not be ok == false")
			}
			if next != i {
				t.Fatalf("mismatched. expected = %d, actual = %d", i, next)
			}
		}

		for i := 0; i < 100; i++ {
			if _, ok := iter.Next(); ok {
				t.Fatalf("must not be ok == true")
			}
		}

		for i := 0; i < 25; i = i + 5 {
			list.PushBack(i)
		}

		for i := 0; i < 25; i = i + 5 {
			next, ok := iter.Next()
			if !ok {
				t.Fatal("must not be ok == false")
			}
			if next != i {
				t.Fatalf("mismatched. expected = %d, actual = %d", i, next)
			}
		}
	})
}
