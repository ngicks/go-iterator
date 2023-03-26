package iterator_test

import "testing"

func FuzzIterator(f *testing.F) {
	f.Add([]byte("abcdefghijklmnopqrstuvwxyz1234567890-=[];',./!@#$%^&*()_+{}:|<>?"))
	f.Fuzz(func(t *testing.T, input []byte) {
		testIterator(t, input)
	})
}
