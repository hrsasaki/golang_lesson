package main

import (
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 2 {
		println("required just 2 arguments.")
		return
	}
	if isAnagrams(os.Args[1], os.Args[2]) {
		println(os.Args[1] + ", " + os.Args[2] + " : anagrams")
		return
	}
	println(os.Args[1] + ", " + os.Args[2] + " : not anagrams")
}

func isAnagrams(s1, s2 string) bool {
	for _, c := range s1 {
		count := strings.Count(s1, string(c))
		if strings.Count(s2, string(c)) != count {
			return false
		}
	}
	return true
}
