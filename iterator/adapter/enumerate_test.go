package adapter_test

import (
	"testing"

	"github.com/ngicks/go-iterator/iterator"
)

func TestEnumerate(t *testing.T) {
	input := []string{"foo", "bar", "baz", "qux", "quux"}
	iter := iterator.FromSlice(input)
	enum := iterator.Enumerate[string](iter)

	for idx, v := range input {
		next := enum.MustNext()
		if next.Count != idx {
			t.Fatalf("%d, %d\n", idx, next.Count)
		}
		if next.Next != v {
			t.Fatalf("%s, %s\n", v, next.Next)
		}
	}

	if _, ok := enum.Next(); ok {
		t.Fatal("must be exhausted")
	}
}
