package base

import (
	"io"
	"os"
)

type ReadResult struct {
	N    int
	Data []byte
	Err  error
}

type ReaderIter struct {
	reader   io.Reader
	bufSize  int
	reuseBuf []byte
	savedErr error
}

func NewReaderIter(reader io.Reader, bufSize int, reuseBuf bool) *ReaderIter {
	i := &ReaderIter{
		reader:  reader,
		bufSize: bufSize,
	}
	if reuseBuf {
		i.reuseBuf = make([]byte, bufSize)
	}
	return i
}

func (i *ReaderIter) Next() (next ReadResult, ok bool) {
	if i.savedErr != nil {
		return ReadResult{Err: i.savedErr}, false
	}

	var buf []byte
	if i.reuseBuf != nil {
		buf = i.reuseBuf
	} else {
		buf = make([]byte, i.bufSize)
	}

	var result ReadResult
	result.N, result.Err = i.reader.Read(buf)
	result.Data = buf[:result.N]
	if result.Err != nil {
		i.savedErr = result.Err
	}
	if result.Err == io.EOF {
		return result, false
	}
	return result, true
}

func (i *ReaderIter) SizeHint() int {
	if f, ok := i.reader.(*os.File); ok {
		stat, err := f.Stat()
		if err != nil {
			return -1
		}
		return int(stat.Size())
	}
	return -1
}
