// See page 9.

// ch01/ex03 measures difference of time between 2 algorithms, substituting and using strings.Join.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//!+

func main() {
	// substituting (low performance)
	fmt.Println("--- substituting ---")
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	lowPerformanceTime := time.Since(start)
	fmt.Println(lowPerformanceTime)

	// using strings.Join (high performance)
	fmt.Println("--- using strings.Join ---")
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	highPerformanceTime := time.Since(start)
	fmt.Println(highPerformanceTime)

	// difference between 2 algorithms above
	fmt.Println("### DIFFERENCE ###")
	fmt.Println(lowPerformanceTime - highPerformanceTime)
}

//!-
