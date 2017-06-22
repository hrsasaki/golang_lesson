package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex01: %v\n", err)
		os.Exit(1)
	}
	for _, element := range ElementsByTagName(nil, doc, "div", "a") {
		fmt.Printf("tagname: %s\t", element.Data)
		for _, attr := range element.Attr {
			fmt.Printf("%s: %s\t", attr.Key, attr.Val)
		}
		fmt.Println()
	}
}

func ElementsByTagName(elements []*html.Node, doc *html.Node, name ...string) []*html.Node {
	if doc.Type == html.ElementNode {
		for _, n := range name {
			if doc.Data == n {
				elements = append(elements, doc)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		elements = ElementsByTagName(elements, c, name...)
	}
	return elements
}
