// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

const (
	sha256Str = "sha256"
	sha384Str = "sha384"
	sha512Str = "sha512"
)

var alg = flag.String("a", "sha256", "hash algorithm")
var input = flag.String("v", "hello world", "input")

func main() {
	flag.Parse()

	if *alg == sha256Str {
		fmt.Printf("%s for %q: %x\n", sha256Str, *input, sha256.Sum256([]byte(*input)))
	} else if *alg == sha384Str {
		fmt.Printf("%s for %q: %x\n", sha384Str, *input, sha512.Sum384([]byte(*input)))
	} else if *alg == sha512Str {
		fmt.Printf("%s for %q: %x\n", sha512Str, *input, sha512.Sum512([]byte(*input)))
	} else {
		fmt.Printf("invalid alg argument: %v\n", *alg)
		os.Exit(1)
	}
}

//!+

func main1() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//!-
