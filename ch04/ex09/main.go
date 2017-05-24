// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // counts of word
	invalid := 0                   // count of invalid UTF-8 characters

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex09: %v\n", err)
		os.Exit(1)
	}
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}

	fmt.Printf("word\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
