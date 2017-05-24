// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import "fmt"

//!+
import "crypto/sha256"

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	result := CountDefferentBit(c1, c2)
	fmt.Print("number of different bit: ")
	fmt.Println(result)
}

func CountDefferentBit(x, y [32]byte) int {
	count := 0
	for i, _ := range x {
		count += PopCount(x[i] ^ y[i]) // XOR
	}
	return count
}

func PopCount(x byte) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

//!-
