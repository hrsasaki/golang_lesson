package ex04

import (
	"io"
	"strings"
)

type Ex04Reader struct {
	str    string
	reader io.Reader
}

func (r *Ex04Reader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	return
}

func NewReader(str string) io.Reader {
	var reader = Ex04Reader{str, strings.NewReader(str)}
	return &reader
}
