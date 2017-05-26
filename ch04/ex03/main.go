// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import "fmt"

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s *[6]int) {
	for i, elem := range *s { // 素直に添え字を見て比較するのが模範解答
		s[len(*s)-i-1] = elem
	}
}

//!-rev
