package base

import (
	"github.com/ngicks/genericcontainer/list"
)

// Doubly ended iterator made from List.
type ListIterDe[T any] struct {
	listLen      int
	done         bool
	advanceFront int
	advanceBack  int
	eleFront     *list.Element[T]
	eleBack      *list.Element[T]
}

// NewListIterDe makes *ListIterDe[T] from list.List[T].
// Range is fixed at the time NewListIterDe returns.
// Mutating passed list outside this iterator may cause undefined behavior.
func NewListIterDe[T any](l *list.List[T]) *ListIterDe[T] {
	return &ListIterDe[T]{
		listLen:  l.Len(),
		eleFront: l.Front(),
		eleBack:  l.Back(),
	}
}

func (li *ListIterDe[T]) Next() (next T, ok bool) {
	if li.done {
		return
	}
	if li.eleFront == li.eleBack {
		li.done = true
	}
	if li.eleFront == nil {
		return
	}
	next, ok = li.eleFront.Get()
	if !ok {
		return
	}
	li.eleFront = li.eleFront.Next()
	li.advanceFront++
	return
}

func (li *ListIterDe[T]) NextBack() (next T, ok bool) {
	if li.done {
		return
	}
	if li.eleFront == li.eleBack {
		li.done = true
	}
	if li.eleBack == nil {
		return
	}
	next, _ = li.eleBack.Get()
	ok = true
	li.eleBack = li.eleBack.Prev()
	li.advanceBack++
	return
}

// SizeHint hints size of remaining elements.
// Size would be incorrect if and only if new element is inserted
// into between head and tail of the iterator.
func (li *ListIterDe[T]) SizeHint() int {
	return li.listLen - li.advanceFront - li.advanceBack
}

// ListIterSe is monotonic list iterator. It only advances to tail.
// ListIterSe is not fused, its Next might return ok=true after it returns ok=false.
// This happens when passed list grows its tail afterwards.
type ListIterSe[T any] struct {
	root     *list.List[T]
	ele      *list.Element[T]
	advanced bool
}

func NewListIterSe[T any](l *list.List[T]) *ListIterSe[T] {
	return &ListIterSe[T]{
		root:     l,
		ele:      l.Front(),
		advanced: true,
	}
}

func (li *ListIterSe[T]) Next() (next T, ok bool) {
	if li.ele == nil {
		if li.root.Front() == nil {
			return
		}
		li.ele = li.root.Front()
	}

	if !li.advanced {
		nextEle := li.ele.Next()
		if nextEle == nil {
			return next, false
		} else {
			li.ele = nextEle
		}
	}

	ele, ok := li.ele.Get()
	nextEle := li.ele.Next()
	if nextEle == nil {
		li.advanced = false
	} else {
		li.ele = nextEle
		li.advanced = true
	}
	return ele, ok
}
