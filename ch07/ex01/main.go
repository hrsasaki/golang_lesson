// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

//!+bytecounter

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanWords)
	*c = 0
	for input.Scan() {
		*c++
	}
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanLines)
	*c = 0
	for input.Scan() {
		*c++
	}
	return len(p), nil
}

//!-bytecounter

func main() {
	var wc WordCounter
	wc.Write([]byte("hello world\naaa"))
	fmt.Println(wc)

	var lc LineCounter
	lc.Write([]byte("hello world\naaa"))
	fmt.Println(lc)
}
