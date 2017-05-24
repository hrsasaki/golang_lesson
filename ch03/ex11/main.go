// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	abs := s
	if strings.HasPrefix(s, "-") {
		buf.WriteString("-")
		abs = s[1:]
	}
	sp := strings.Split(abs, ".")
	buf.WriteString(commaInt(sp[0]))
	if len(sp) > 1 {
		buf.WriteString(".")
		buf.WriteString(removeZero(sp[1]))
	}
	return buf.String()
}

func commaInt(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaInt(s[:n-3]) + "," + s[n-3:]
}

func removeZero(s string) string {
	if strings.HasSuffix(s, "0") {
		return removeZero(s[:len(s)-1])
	}
	return s
}

//!-
