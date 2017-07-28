// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"testing"
)

//!+String
func TestString(t *testing.T) {
	l1 := literal(1.5)
	l2 := literal(2.4)
	l3 := literal(5)
	fmt.Printf("l1 => %s\n", l1.String())
	v1 := Var("foo")
	v2 := Var("bar")
	fmt.Printf("v1 => %s\n", v1.String())
	u1 := unary{'-', l2}
	fmt.Printf("u1 => %s\n", u1.String())
	b1 := binary{'*', l3, v2}
	fmt.Printf("b1 => %s\n", b1.String())
	c1 := call{"sqrt", []Expr{b1, v1, l3}}
	fmt.Printf("c1 => %s\n", c1.String())
}

//!-String

/*
//!+output
l1 => 1.5
v1 => foo
u1 => -[2.4]
b1 => *[5, bar]
c1 => sqrt[*[5, bar], foo, 5]
//!-output
*/
