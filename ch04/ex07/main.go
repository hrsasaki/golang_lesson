// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import "fmt"

func main() {
	//!+slice
	s := []byte("I have a pen.")
	fmt.Println(string(s))
	reverse(s)
	fmt.Println(string(s))
	//!-slice
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev
