package palindrome

import "testing"

func TestIsParindrome(t *testing.T) {
	p1 := Palindrome("pop pop")
	if !IsPalindrome(p1) {
		t.Errorf("TestIsParindrome: '%s' is failed", p1)
	}

	p2 := Palindrome("しんぶんし")
	if !IsPalindrome(p2) {
		t.Errorf("TestIsParindrome: '%s' is failed", p2)
	}

	p3 := Palindrome("execute")
	if IsPalindrome(p3) {
		t.Errorf("TestIsParindrome: '%s' is failed", p3)
	}
}
