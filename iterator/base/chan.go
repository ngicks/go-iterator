package base

type ChanIter[T any] struct {
	channel <-chan T
}

// NewChanIter makes ChanIter associated with a given channel.
// To signal end of the iterator, close the passed channel.
//
// *ChanIter[T] only implements SeIterator[T].
func NewChanIter[T any](channel <-chan T) *ChanIter[T] {
	return &ChanIter[T]{
		channel: channel,
	}
}

// Next blocks until internal channel receives a value and returns the value.
func (ci *ChanIter[T]) Next() (next T, ok bool) {
	next, ok = <-ci.channel
	return
}
