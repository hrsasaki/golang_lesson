// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

//!-nonempty

func main() {
	//!+main
	data := []string{"one", "one", "two", "three", "three", "four"}
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", removedup(data)) // `["one" "two" "three" "four"]`
	//!-main
}

func removedup(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for i := 0; i < len(strings); i++ {
		if i == 0 {
			out = append(out, strings[i])
		} else if strings[i] != strings[i-1] {
			out = append(out, strings[i])
		}
	}
	return out
}

//!-
