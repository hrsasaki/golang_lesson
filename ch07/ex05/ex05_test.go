package ex05

import (
	"io"
	"os"
	"testing"
)

func TestReaderEOF(t *testing.T) {
	r := LimitReader(os.Stdin, 5)
	_, err := r.Read([]byte("Hello, World"))
	if err != io.EOF {
		t.Error("not EOF")
	}
}
