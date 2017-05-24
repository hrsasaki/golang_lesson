// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os" //!+
)

func main() {
	if len(os.Args) < 3 {
		c := sha256.Sum256([]byte(os.Args[1]))
		fmt.Printf("SHA-256\n %x\n", c)
		return
	}
	if os.Args[2] == "384" {
		c := sha512.Sum384([]byte(os.Args[1]))
		fmt.Printf("SHA-384\n %x\n", c)
	} else if os.Args[2] == "512" {
		c := sha512.Sum512([]byte(os.Args[1]))
		fmt.Printf("SHA-512\n %x\n", c)
	} else {
		c := sha256.Sum256([]byte(os.Args[1]))
		fmt.Printf("SHA-256\n %x\n", c)
	}
}

//!-
