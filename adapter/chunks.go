package adapter

import (
	"fmt"
	"math"
)

type Chunker[T any] struct {
	size uint
	sl   []T
}

func NewChunker[T any](sl []T, size uint) *Chunker[T] {
	if size > math.MaxInt {
		panic(
			fmt.Sprintf(
				"NewChunker: size must not be larger than max int. input = %d, want <= %d",
				size, math.MaxInt,
			),
		)
	}
	return &Chunker[T]{
		size: size,
		sl:   sl,
	}
}

// Next returns next element of this iterator.
// len(next) != size if remaining element is less than size
func (c *Chunker[T]) Next() (next []T, ok bool) {
	if len(c.sl) == 0 {
		return
	}
	if len(c.sl) < int(c.size) {
		next = c.sl[:]
		c.sl = c.sl[:0]
		return next, true
	}
	next, c.sl = c.sl[:c.size], c.sl[c.size:]
	return next, true
}
