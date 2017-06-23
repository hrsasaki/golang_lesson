// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
	"io"
	"os"
)

//!+bytecounter
type CountWriter struct {
	Writer io.Writer
	Count  int64
}

func (w *CountWriter) Write(p []byte) (int, error) {
	n, err := w.Writer.Write(p)
	w.Count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cw = CountWriter{w, 0}
	return &cw, &cw.Count
}

//!-bytecounter

func main() {
	//!+main
	w, n := CountingWriter(os.Stdout)
	fmt.Fprintln(w, "Hello, World")
	fmt.Println(*n)
	//!main
}
