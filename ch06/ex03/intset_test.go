// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "fmt"

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println(x.Len()) // "4"

	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
	// 4
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	x.Remove(1)
	fmt.Println(x.String()) // "{9 42}"

	x.Remove(144)
	fmt.Println(x.String()) // "{9 42}"
	fmt.Println(x)          // "{[4398046511616]}"

	x.Clear()
	fmt.Println(x.String()) // "{}"
	fmt.Println(x)          // "{[0]}"

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
	// {9 42 144}
	// {9 42}
	// {[4398046511616]}
	// {}
	// {[0]}
}

func Example_three() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y = *(x.Copy())
	y.Add(200)
	fmt.Println(y.String())
	fmt.Println(x.String())

	x.AddAll(1, 3, 96)
	fmt.Println(x.String())

	// Output:
	// {1 9 144 200}
	// {1 9 144}
	// {1 3 9 96 144}
}

func Example_four() {
	var x, y IntSet
	x.AddAll(1, 36, 9, 144)
	y.AddAll(9, 36, 42)
	x.IntersectWith(&y)
	fmt.Println(x.String())

	x.Clear()
	y.Clear()
	x.AddAll(1, 36, 9)
	y.AddAll(3, 9, 36, 42)
	x.IntersectWith(&y)
	fmt.Println(x.String())

	x.Clear()
	y.Clear()
	x.AddAll(1, 36, 9)
	y.AddAll(3, 9, 36, 42)
	x.SymmetricDifferenceWith(&y)
	fmt.Println(x.String())

	x.Clear()
	y.Clear()
	x.AddAll(1, 36, 9)
	y.AddAll(3, 9, 36, 42)
	x.DifferenceWith(&y)
	fmt.Println(x.String())
	// Output:
	// {9 36}
	// {9 36}
	// {1 3 42}
	// {1}
}
