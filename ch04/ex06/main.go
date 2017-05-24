// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
	"unicode"
)

func main() {
	//!+slice
	s := "I have  a pen."
	fmt.Println(s)
	fmt.Println(removeSpace(s))
	//!-slice
}

// reverse reverses a slice of ints in place.
func removeSpace(s string) string {
	runes := []rune(s)
	out := runes[:0] // zero-length slice of original
	for i := 0; i < len(runes); i++ {
		if i == 0 {
			out = append(out, runes[i])
		} else if !unicode.IsSpace(runes[i]) || !unicode.IsSpace(runes[i-1]) {
			out = append(out, runes[i])
		}
	}
	return string(out)
}

//!-
