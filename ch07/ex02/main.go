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
	Count  *int64
}

func (w *CountWriter) Write(p []byte) (int, error) {
	num := int64(len(p))
	w.Count = &num
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := new(CountWriter)
	cw.Write([]byte("hello"))
	return cw.Writer, cw.Count
}

//!-bytecounter

func main() {
	//!+main
	_, n := CountingWriter(os.Stdout)
	fmt.Println(&n)
	//!main
}
