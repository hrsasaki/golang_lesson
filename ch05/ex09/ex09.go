package ex09

import (
	"bufio"
	"strings"
)

func expand(s string, f func(string) string) string {
	input := bufio.NewScanner(strings.NewReader(s))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		if strings.HasPrefix(input.Text(), "$") {
			word := strings.Split(input.Text(), "$")[1]
			s = strings.Replace(s, input.Text(), f(word), 1)
		}
	}
	return s
}
