package iterator

import "github.com/ngicks/go-iterator/iterator/def"

func Fold[T, U any](iter def.SeIterator[T], reducer func(accumulator U, next T) U, initial U) U {
	var accum U = initial
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		accum = reducer(accum, next)
	}
	return accum
}
