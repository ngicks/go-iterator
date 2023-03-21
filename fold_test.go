package iterator_test

import (
	"fmt"
	"testing"

	"github.com/ngicks/iterator"
)

func TestFold(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	iter := iterator.FromSlice(input)

	mm := iterator.Fold[int](iter, func(mm map[int]string, next int) map[int]string {
		mm[next] = fmt.Sprintf("foo%d", next)
		return mm
	}, make(map[int]string))

	for _, v := range input {
		if fmt.Sprintf("foo%d", v) != mm[v] {
			t.Fatalf("%+v\n", mm)
		}
	}
	if len(input) != len(mm) {
		t.Fatalf("%+v,\n%+v\n", input, mm)
	}
}
