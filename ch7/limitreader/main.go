package main

import (
	"io"
	"math"
	"os"
	"strings"
)

type LimitReader struct {
	reader io.Reader
	limit  int64
}

func (lr *LimitReader) Read(p []byte) (n int, err error) {
	if lr.limit <= 0 {
		return 0, io.EOF
	}
	len := int64(math.Min(float64(len(p)), float64(lr.limit)))
	p = p[:len]
	n, err = lr.reader.Read(p)
	lr.limit -= len
	return
}

func NewLimitReader(r io.Reader, n int64) io.Reader {
	return &LimitReader{reader: r, limit: n}
}

func main() {
	reader := strings.NewReader("hello world")
	lr := NewLimitReader(reader, 5)
	io.Copy(os.Stdout, lr)
}
