// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	// $GOPATH 配下でなければ、相対パスによる参照も可能
	"../ex01/tempconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Input a decimal.")
		var t float64
		_, err := fmt.Scan(&t)
		handleError(err)
		printTemp(t)
	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			handleError(err)
			printTemp(t)
		}
	}
}

func printTemp(temp float64) {
	f := tempconv.Fahrenheit(temp)
	c := tempconv.Celsius(temp)
	k := tempconv.Kelvin(temp)
	fmt.Printf("%s = %s, %s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c), k, tempconv.KToF(k))
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
}

//!-
