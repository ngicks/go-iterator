package base

import "github.com/ngicks/iterator/adapter"

// MapIterDe is doubly ended iterator,
// which is made from map.
type MapIterDe[T comparable, U any] struct {
	innerMap map[T]U
	keys     []T
	idxFront int
	idxBack  int
}

// NewMapIterDe makes MapIterDe[T, U] from map[T]U.
//
// m should not be mutated after the return of this function. Otherwise behavior is undefined.
//
// keySortOption is used to sort its iteration order.
// if nil, order is random
// (This is default behavior of range expression applied to a map,
// as per the Go programming language specification.)
func NewMapIterDe[T comparable, U any](m map[T]U, keySortOption func(keys []T) []T) *MapIterDe[T, U] {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	if keySortOption != nil {
		keys = keySortOption(keys)
	}

	return &MapIterDe[T, U]{
		innerMap: m,
		keys:     keys,
		idxFront: 0,
		idxBack:  len(keys) - 1,
	}
}

func (mi *MapIterDe[T, U]) Next() (next adapter.Zipped[T, U], ok bool) {
	if mi.idxFront > mi.idxBack {
		return
	}
	var val U
	key := mi.keys[mi.idxFront]
	val, ok = mi.innerMap[key]
	mi.idxFront++
	return adapter.Zipped[T, U]{
		Former: key,
		Latter: val,
	}, ok
}

func (mi *MapIterDe[T, U]) NextBack() (next adapter.Zipped[T, U], ok bool) {
	if mi.idxFront > mi.idxBack {
		return
	}
	var val U
	key := mi.keys[mi.idxBack]
	val, ok = mi.innerMap[key]
	mi.idxBack--
	return adapter.Zipped[T, U]{
		Former: key,
		Latter: val,
	}, ok
}

// SizeHint returns size of remaining elements.
// This is not an actual count of non-used elements in underlying map,
// since it ignores keys added after return of NewMapIterDe.
//
// Internally, MapIterDe stores keys of map as []T, in NewMapIterDe,
// and consumes the slice from both end.
func (mi *MapIterDe[T, U]) SizeHint() int {
	return mi.idxBack - mi.idxFront + 1
}
