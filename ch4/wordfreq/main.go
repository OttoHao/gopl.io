package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	freqs := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		freqs[strings.ToLower(scanner.Text())]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Printf("word\tcount\n")
	for k, v := range freqs {
		fmt.Printf("%q\t%d\n", k, v)
	}

}
