// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 25.
//!+

// ch01/ex12 is a lissajous server.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
)

var xmin, ymin, xmax, ymax = -2, -2, +2, +2
var width, height = 1024, 1024
var xorigin, yorigin = 0, 0
var scale = 1

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Form {
		switch k {
		case "x":
			tmp, err := strconv.Atoi(v[0])
			if err != nil {
				handleError(err)
			}
			xorigin = tmp
		case "y":
			tmp, err := strconv.Atoi(v[0])
			if err != nil {
				handleError(err)
			}
			yorigin = tmp
		case "scale":
			tmp, err := strconv.Atoi(v[0])
			if err != nil {
				handleError(err)
			}
			scale = tmp
		}
	}
	generatePng(w)
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
}

// from gopl.io/ch3/mandelbrot.
func generatePng(out io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*float64(ymax-ymin) + float64(ymin)
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*float64(xmax-xmin) + float64(xmin)
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-
