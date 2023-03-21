package main

import (
	"fmt"
	"strings"

	"github.com/ngicks/iterator"
	"github.com/ngicks/iterator/adapter"
)

func main() {
	fmt.Println(
		iterator.Fold[adapter.Enumerated[string]](
			iterator.Enumerate[string](
				iterator.
					FromSlice([]string{"foo", "bar", "baz"}).
					Map(func(s string) string { return s + s }).
					Exclude(func(s string) bool { return strings.Contains(s, "az") }).
					MustReverse(),
			),
			func(accumulator map[string]int, next adapter.Enumerated[string]) map[string]int {
				accumulator[next.Next] = next.Count
				return accumulator
			},
			map[string]int{"initialKey": -1},
		),
	)
}
