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
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

var cycles int = 5      // number of complete x oscillator revolutions
var res float64 = 0.001 // angular resolution
var size int = 100      // image canvas covers [-size..+size]
var nframes int = 64    // number of animation frames
var delay int = 8       // delay between frames in 10ms units

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
	for k, v := range r.Form {
		switch k {
		case "cycles":
			tmp, err := strconv.Atoi(v[0])
			if err != nil {
				handleError(err)
			}
			cycles = tmp
		case "res":
			tmp, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				handleError(err)
			}
			res = tmp
		case "size":
			tmp, err := strconv.Atoi(v[0])
			if err != nil {
				handleError(err)
			}
			size = tmp
		case "nframes":
			tmp, err := strconv.Atoi(v[0])
			if err != nil {
				handleError(err)
			}
			nframes = tmp
		case "delay":
			tmp, err := strconv.Atoi(v[0])
			if err != nil {
				handleError(err)
			}
			delay = tmp
		}
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
}

// from gopl.io/ch1/lissajous.
func lissajous(out io.Writer) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-
