// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex02: %v\n", err)
		os.Exit(1)
	}
	elementMap := make(map[string]int)
	for element, count := range countElements(elementMap, doc) {
		fmt.Println(element, count)
	}
}

//!-main

func countElements(elementMap map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elementMap[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elementMap = countElements(elementMap, c)
	}
	return elementMap
}
