// rotate a slice.
package main

import "fmt"

func main() {
	//!+array
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(a, 3)) // "[3, 4, 5, 0, 1, 2]"
	//!-array
}

func rotate(s []int, n int) []int {
	for i := 0; i < n; i++ {
		s = append(s, s[i])
	}
	return s[n:]
}
