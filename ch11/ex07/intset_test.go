// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"math/rand"
	"testing"
	"time"
)

func TestHas(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{0, true},
		{1, false},
		{3, true},
		{64, true},
		{66, false},
	}
	intSet := IntSet{[]uint64{9, 1}}
	for _, test := range tests {
		if got := intSet.Has(test.input); got != test.want {
			t.Errorf("IntSet %s.Has(%d) = %v", intSet, test.input, got)
		}
	}
}

func TestAddAndUnion(t *testing.T) {
	var tests = []struct {
		input int
		want  IntSet
	}{
		{6, IntSet{[]uint64{65}}},
		{128, IntSet{[]uint64{65, 0, 1}}},
	}
	intSet := IntSet{[]uint64{1}}
	for _, test := range tests {
		intSet.Add(test.input)
		if intSet.String() != test.want.String() {
			t.Errorf("IntSet: %s", intSet.String())
		}
	}
}

func TestUnion(t *testing.T) {
	// var x IntSet
	// x.Add(5)
	// x.Add(128)
	var tests = []struct {
		input IntSet
		want  IntSet
	}{
		{IntSet{[]uint64{4}}, IntSet{[]uint64{13}}},
		{IntSet{[]uint64{32, 0, 1, 5, 24}}, IntSet{[]uint64{41, 0, 1, 5, 24}}},
		{IntSet{[]uint64{32, 0, 1}}, IntSet{[]uint64{41, 0, 1}}},
	}
	for _, test := range tests {
		intSet := IntSet{[]uint64{9}}
		intSet.UnionWith(&test.input)
		if intSet.String() != test.want.String() {
			t.Errorf("IntSet: %s want %s", intSet.String(), test.want.String())
		}
	}
}

// Benchmark
func benchmarkAdd(b *testing.B, max int) {
	rand.Seed(time.Now().UnixNano())
	intSet := IntSet{[]uint64{1}}
	for i := 0; i < b.N; i++ {
		// 毎回同じデータを渡すようにしたいので、forループの外であらかじめデータを作っておく
		// 一つの IntSet に Add し続けているのはテストとしてまずい
		intSet.Add(rand.Intn(max))
	}
}

func BenckmarkAddRand100(b *testing.B) {
	benchmarkAdd(b, 100)
}

func BenckmarkAddRand1000(b *testing.B) {
	benchmarkAdd(b, 1000)
}

func benchmarkUnion(b *testing.B, max, length int) {
	rand.Seed(time.Now().UnixNano())
	var x, y IntSet
	for i := 0; i < length; i++ {
		x.Add(rand.Intn(max))
		y.Add(rand.Intn(max))
	}
	for i := 0; i < b.N; i++ {
		x.UnionWith(&y)
	}
}

func BenchmarkUnionM100L10(b *testing.B) {
	benchmarkUnion(b, 100, 100)
}

func BenchmarkUnionM1000L10(b *testing.B) {
	benchmarkUnion(b, 1000, 100)
}

func BenchmarkUnionM100L100(b *testing.B) {
	benchmarkUnion(b, 100, 1000)
}

func BenchmarkUnionM1000L100(b *testing.B) {
	benchmarkUnion(b, 1000, 1000)
}
