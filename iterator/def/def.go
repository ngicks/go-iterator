package def

// Singly ended iterator.
type SeIterator[T any] interface {
	Next() (next T, ok bool)
}

// Doubly ended iterator.
type DeIterator[T any] interface {
	SeIterator[T]
	NextBack() (next T, ok bool)
}

type SizeHinter interface {
	SizeHint() int
}

type Reverser[T any] interface {
	Reverse() (rev SeIterator[T], ok bool)
}

type Unwrapper[T any] interface {
	Unwrap() SeIterator[T]
}
