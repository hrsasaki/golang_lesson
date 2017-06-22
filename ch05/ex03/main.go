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
	for _, textNodes := range searchTextNodes(nil, doc) {
		fmt.Println(textNodes)
	}
}

//!-main

func searchTextNodes(textNodes []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		textNodes = append(textNodes, n.Data)
	}
	if n.Data != "script" && n.Data != "style" {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			textNodes = searchTextNodes(textNodes, c)
		}
	}
	return textNodes
}
