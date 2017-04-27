// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount

import (
	"testing"
)

// -- Input-Output --

func TestPopCount(t *testing.T) {
	actual := PopCount(158)
	if actual != 5 {
		t.Errorf("TestPopCount(158) failed: actual = %d", actual)
	}
}

func TestPopCountByLoop(t *testing.T) {
	actual := PopCountByLoop(158)
	if actual != 5 {
		t.Errorf("TestPopCount(158) failed: actual = %d", actual)
	}
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoop(0x1234567890ABCDEF)
	}
}
