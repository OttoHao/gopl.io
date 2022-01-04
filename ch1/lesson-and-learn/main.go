package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	files := os.Args[1:]
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
		}
		defer f.Close()
		readByIoutil(f)
	}
}

func readByBufio(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func readByIoutil(reader io.Reader) {
	input, _ := ioutil.ReadAll(reader)
	fmt.Println(string(input))
}
