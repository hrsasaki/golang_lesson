// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py += 2 {
		// TODO: 関数化する
		y1 := float64(py)/height*(ymax-ymin) + ymin
		y2 := float64(py+1)/height*(ymax-ymin) + ymin

		for px := 0; px < width; px += 2 {
			// TODO: 関数化する
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := float64(px+1)/width*(xmax-xmin) + xmin

			// (px, py)と(px+1, py+1)が張る矩形内のピクセルの色の平均をとる
			z := [2][2]complex128{{complex(x1, y1), complex(x1, y2)}, {complex(x2, y1), complex(x2, y2)}}
			r11, g11, b11, a11 := mandelbrot(z[0][0]).RGBA()
			r12, g12, b12, a12 := mandelbrot(z[0][1]).RGBA()
			r21, g21, b21, a21 := mandelbrot(z[1][0]).RGBA()
			r22, g22, b22, a22 := mandelbrot(z[1][1]).RGBA()
			avgColor := color.RGBA{
				uint8((r11 + r12 + r21 + r22) / 4),
				uint8((g11 + g12 + g21 + g22) / 4),
				uint8((b11 + b12 + b21 + b22) / 4),
				uint8((a11 + a12 + a21 + a22) / 4)}

			// Image point (px, py) represents complex value z.
			img.Set(px, py, avgColor)
			img.Set(px+1, py, avgColor)
			img.Set(px, py+1, avgColor)
			img.Set(px+1, py+1, avgColor)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
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

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
