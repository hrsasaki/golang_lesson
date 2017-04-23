// See page 14.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Word struct {
	word  string
	count int
	files []string
}

func main() {
	var words []Word
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, words)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, words)
			f.Close()
		}
	}
	for _, w := range words {
		if w.count > 1 {
			fmt.Printf("%d\t%s found in ", w.count, w.word)
			for n, name := range w.files {
				if n != 0 {
					fmt.Printf(",")
				}
				fmt.Printf("%s ", name)
			}
			fmt.Println("")
		}
	}
}

func countLines(f *os.File, words []Word) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		var word Word
		if !containsWord(words, input.Text(), &word) {
			words = append(words, word)
		}
		word.count++
		word.files = append(word.files, f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}

func containsWord(words []Word, line string, word *Word) bool {
	for _, w := range words {
		if w.word == line {
			word = &w
			return true
		}
	}
	word = &Word{
		word:  line,
		count: 0,
		files: make([]string, 0),
	}
	return false
}

//!-
