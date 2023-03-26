package iterator

import (
	"bytes"
	"io"

	"github.com/ngicks/go-iterator/iterator/def"
)

func Reader[T any](iter def.SeIterator[T], marshaler func(T) ([]byte, error)) io.Reader {
	return &IterReader[T]{
		iter:      iter,
		marshaler: marshaler,
	}
}

type IterReader[T any] struct {
	iter      def.SeIterator[T]
	marshaler func(T) ([]byte, error)
	savedErr  error
	read      bytes.Buffer
}

func (i *IterReader[T]) Read(p []byte) (n int, err error) {
	if i.savedErr != nil && (i.savedErr != io.EOF || i.read.Len() == 0) {
		return 0, i.savedErr
	}

	if i.savedErr == nil && i.read.Len() < len(p) {
		next, ok := i.iter.Next()
		if !ok {
			i.savedErr = io.EOF
			if i.read.Len() == 0 {
				return 0, i.savedErr
			}
		} else {
			marshalled, err := i.marshaler(next)
			if err != nil {
				i.savedErr = err
				return 0, i.savedErr
			}
			i.read.Write(marshalled)
		}
	}

	copied := copy(p, i.read.Next(len(p)))
	return copied, nil
}

func (i *IterReader[T]) WriteTo(w io.Writer) (n int64, err error) {
	if i.savedErr != nil {
		return 0, i.savedErr
	}

	const (
		maxReadSize = 64 * 1024 * 1024
	)
	var nn int64

	i.read.Grow(maxReadSize)

	for {
		for i.read.Len() < maxReadSize/2 {
			next, ok := i.iter.Next()
			if !ok {
				i.savedErr = io.EOF
				break
			}
			marshalled, err := i.marshaler(next)
			if err != nil {
				i.savedErr = err
				break
			}
			i.read.Write(marshalled)
		}

		if i.savedErr != nil && i.savedErr != io.EOF {
			return 0, i.savedErr
		}

		n, err := i.read.WriteTo(w)
		nn += n
		if err != nil {
			i.savedErr = err
			return nn, i.savedErr
		}

		if i.savedErr == io.EOF {
			break
		}
	}

	return nn, nil
}
