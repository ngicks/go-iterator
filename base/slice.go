package base

// SliceIterDe is doubly ended iterator,
// which is made of slice.
type SliceIterDe[T any] struct {
	innerSlice []T
	idxFront   int
	idxBack    int
}

// NewSliceIterDe makes SliceIterDe[T] from []T.
func NewSliceIterDe[T any](sl []T) *SliceIterDe[T] {
	return &SliceIterDe[T]{
		innerSlice: sl,
		idxFront:   0,
		idxBack:    len(sl) - 1,
	}
}

func (si *SliceIterDe[T]) Next() (next T, ok bool) {
	if si.idxFront > si.idxBack {
		return
	}
	next, ok = si.innerSlice[si.idxFront], true
	si.idxFront++
	return
}
func (si *SliceIterDe[T]) NextBack() (next T, ok bool) {
	if si.idxFront > si.idxBack {
		return
	}
	next, ok = si.innerSlice[si.idxBack], true
	si.idxBack--
	return
}

// SizeHint returns size of remaining elements.
func (si *SliceIterDe[T]) SizeHint() int {
	return si.idxBack - si.idxFront + 1
}
