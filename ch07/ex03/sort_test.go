// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestTree(t *testing.T) {
	tr := add(nil, 5)
	add(tr, 4)
	add(tr, 2)
	add(tr, 9)
	add(tr, 6)
	add(tr, 11)

	expected := "5\n\tleft: 4\n\t\tleft: 2\n\tright: 9\n\t\tleft: 6\n\t\tright: 11\n"
	actual := tr.String()
	fmt.Printf("%s\n", actual)
	if actual != expected {
		t.Errorf("")
	}
}
