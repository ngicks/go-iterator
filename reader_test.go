package iterator_test

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"io"
	"testing"

	"github.com/ngicks/iterator"
)

type testJson struct {
	Count      int
	RandomByte []byte
}

type testNDJsonIter struct {
	limit int
	count int
}

func (iter *testNDJsonIter) Next() (testJson, bool) {
	if iter.count > iter.limit {
		return testJson{}, false
	}
	buf := new(bytes.Buffer)
	buf.Grow(32)
	io.CopyN(buf, rand.Reader, 32)
	ret := testJson{
		Count:      iter.count,
		RandomByte: buf.Bytes(),
	}
	iter.count++
	return ret, true
}

func FuzzReader(f *testing.F) {
	f.Add(20, 13, 23, 83) // prime numbers
	f.Fuzz(func(t *testing.T, limitSize, bufSize0, bufSize1, bufSize2 int) {
		if limitSize <= 0 || bufSize0 <= 0 || bufSize1 <= 0 || bufSize2 <= 0 {
			t.Skip()
		}

		{
			iter := &testNDJsonIter{
				limit: limitSize,
			}

			r := iterator.Reader[testJson](
				iter,
				func(tj testJson) ([]byte, error) { return json.Marshal(tj) },
			)
			var (
				err  error
				n    int
				read bytes.Buffer
			)
			bufSize := [...]int{bufSize0, bufSize1, bufSize2}
			for i := 0; err == nil; i++ {
				buf := make([]byte, bufSize[i%3])
				n, err = r.Read(buf)
				read.Write(buf[:n])
			}

			dec := json.NewDecoder(&read)
			for dec.More() {
				var j testJson
				err := dec.Decode(&j)
				if err != nil {
					t.Fatalf(
						"reader must return stream of valid JSONs, but decoder returned error = %+v",
						err,
					)
				}
			}
		}
		{
			iter := &testNDJsonIter{
				limit: limitSize,
			}

			r := iterator.Reader[testJson](
				iter,
				func(tj testJson) ([]byte, error) { return json.Marshal(tj) },
			)
			var buf bytes.Buffer
			_, err := r.(io.WriterTo).WriteTo(&buf)
			if err != nil {
				t.Fatalf("reader.WriteTo must not return error but is %+v", err)
			}

			dec := json.NewDecoder(&buf)
			for dec.More() {
				var j testJson
				err := dec.Decode(&j)
				if err != nil {
					t.Fatalf(
						"reader must return stream of valid JSONs, but decoder returned error = %+v",
						err,
					)
				}
			}
		}
	})
}
