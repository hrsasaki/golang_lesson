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

func TestPopCountByShift(t *testing.T) {
	actual := PopCountByShift(158)
	if actual != 5 {
		t.Errorf("TestPopCount(158) failed: actual = %d", actual)
	}
}

func TestPopCountByClear(t *testing.T) {
	actual := PopCountByClear(158)
	if actual != 5 {
		t.Errorf("TestPopCount(158) failed: actual = %d", actual)
	}
}

// -- Benchmarks --

func benchmarkPopCount(b *testing.B, length uint) {
	var input uint64 = 0x1234567890ABCDEF
	if length <= 15 {
		input = input >> (16 - length) * 4
	}
	for i := 0; i < b.N; i++ {
		PopCount(input)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	benchmarkPopCount(b, 4)
}

func BenchmarkPopCount10(b *testing.B) {
	benchmarkPopCount(b, 10)
}

func BenchmarkPopCount16(b *testing.B) {
	benchmarkPopCount(b, 16)
}

func benchmarkPopCountByShift(b *testing.B, length uint) {
	var input uint64 = 0x1234567890ABCDEF
	if length <= 15 {
		input = input >> (16 - length) * 4
	}
	for i := 0; i < b.N; i++ {
		PopCountByShift(input)
	}
}

func BenchmarkPopCountByShift4(b *testing.B) {
	benchmarkPopCountByShift(b, 4)
}

func BenchmarkPopCountByShift10(b *testing.B) {
	benchmarkPopCountByShift(b, 10)
}

func BenchmarkPopCountByShift16(b *testing.B) {
	benchmarkPopCountByShift(b, 16)
}

func benchmarkPopCountByClear(b *testing.B, length uint) {
	var input uint64 = 0x1234567890ABCDEF
	if length <= 15 {
		input = input >> (16 - length) * 4
	}
	for i := 0; i < b.N; i++ {
		PopCountByClear(input)
	}
}

func BenchmarkPopCountByClear4(b *testing.B) {
	benchmarkPopCountByClear(b, 4)
}

func BenchmarkPopCountByClear10(b *testing.B) {
	benchmarkPopCountByClear(b, 10)
}

func BenchmarkPopCountByClear16(b *testing.B) {
	benchmarkPopCountByClear(b, 16)
}
