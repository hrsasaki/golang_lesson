package palindrome

import "sort"

type Palindrome []rune

func (p Palindrome) Len() int {
	return len(p)
}

func (p Palindrome) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Palindrome) Swap(i, j int) {}

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i <= s.Len()/2; i++ {
		if s.Less(i, s.Len()-1-i) || s.Less(s.Len()-1-i, i) {
			return false
		}
	}
	return true
}
