// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

type ComplexFloat struct {
	real big.Float
	imag big.Float
}

type ComplexRat struct {
	real big.Rat
	imag big.Rat
}

func calcComplex64() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot64(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func calcComplex128() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot128(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func calcFloat() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot128(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if abs64(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrot128(z complex128) color.Color {
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

func abs64(z complex64) float32 {
	abs := math.Hypot(float64(real(z)), float64(imag(z)))
	return float32(abs)
}

func absFloat(z ComplexFloat) big.Float {
	real, _ := z.real.Float64()
	imag, _ := z.imag.Float64()
	abs := math.Hypot(real, imag)
	return *(new(big.Float)).SetFloat64(abs)
}

func absRat(z ComplexRat) big.Rat {

}

func squareFloat(z ComplexFloat) ComplexFloat {
	// (a+bi)^2 = a^2 - b^2 + 2abi
}

func squareRat(z ComplexRat) big.Rat {
	// (a+bi)^2 = a^2 - b^2 + 2abi
	real := big.Rat.Mul(z.real, z.real) - big.Rat.Mul(z.imag, z.imag)
	imag := big.Rat.Mul(big.NewRat(2, 1), big.Rat.Mul(z.real, z.imag))
	return ComplexFloat{real, imag}
}
