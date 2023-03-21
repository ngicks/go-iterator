package adapter

import "fmt"

type Windower[T any] struct {
	width uint
	sl    []T
}

func NewWindower[T any](sl []T, width uint) *Windower[T] {
	if len(sl) < int(width) {
		panic(
			fmt.Sprintf(
				"NewWindower: len(sl) must be greater than / equal to width,"+
					" len(sl) = %d, width = %d",
				len(sl), width,
			),
		)
	}
	return &Windower[T]{
		width: width,
		sl:    sl,
	}
}

func (w *Windower[T]) Next() (next []T, ok bool) {
	if len(w.sl) < int(w.width) {
		return nil, false
	}

	next = w.sl[:w.width]
	w.sl = w.sl[1:]

	return next, true
}
