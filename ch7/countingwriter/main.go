package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	writer io.Writer
	count  int64
}

func (cw *CountWriter) Write(data []byte) (int, error) {
	cw.count += int64(len(data))
	return cw.writer.Write(data)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CountWriter{writer: w}
	return &cw, &cw.count
}

func main() {
	w, c := CountingWriter(os.Stdout)
	w.Write([]byte("hello, world"))
	fmt.Printf("\ncount: %d\n", *c)

	w.Write([]byte("hello, world"))
	fmt.Printf("\ncount: %d\n", *c)
}
