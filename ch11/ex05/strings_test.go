package testspl

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		str  string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"apple,orange,banana,,grape", ",", 5},
		{" 1 2 ", " ", 4},
	}
	for _, test := range tests {
		words := strings.Split(test.str, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d words\n", test.str, test.sep, got, test.want)
		}
	}
}
