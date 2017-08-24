// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "testing"

var intSets = map[string]IntSet{
	"has":   IntSet{[]uint64{9, 1}},
	"add":   IntSet{[]uint64{1}},
	"union": IntSet{[]uint64{9}},
}

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
	intSet := intSets["has"]
	for _, test := range tests {
		if got := intSet.Has(test.input); got != test.want {
			t.Errorf("IntSet %s.Has(%d) = %v", intSets["has"], test.input, got)
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		input int
		want  IntSet
	}{
		{6, IntSet{[]uint64{65}}},
		{128, IntSet{[]uint64{65, 0, 1}}},
	}
	intSet := intSets["add"]
	for _, test := range tests {
		intSet.Add(test.input)
		if intSet.String() != test.want.String() {
			t.Errorf("IntSet: %s", intSet.String())
		}
	}
}

func TestUnion(t *testing.T) {
	var x IntSet
	x.Add(5)
	x.Add(128)
	var tests = []struct {
		input IntSet
		want  IntSet
	}{
		{x, IntSet{[]uint64{41, 0, 1}}},
	}
	intSet := intSets["union"]
	for _, test := range tests {
		intSet.UnionWith(&test.input)
		if intSet.String() != test.want.String() {
			t.Errorf("IntSet: %s", intSet.String())
		}
	}

}
