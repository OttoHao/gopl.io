package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (w *WordCounter) Write(data []byte) (int, error) {
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*w++
	}
	return len(data), nil
}

func main() {
	var w WordCounter
	w.Write([]byte("hello"))
	fmt.Println(w)

	w = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&w, "hello, %s", name)
	fmt.Println(w)
}
