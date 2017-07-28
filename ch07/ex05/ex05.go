package ex05

import "io"

type Ex05Reader struct {
	reader io.Reader
	n      int64
}

func (r *Ex05Reader) Read(p []byte) (n int, err error) {
	if r.n < int64(len(p)) {
		p = p[:r.n]
	}
	// FIXME: ここで err に io.EOF が返る
	n, err = r.reader.Read(p)
	if err == nil {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	var reader = Ex05Reader{r, n}
	return &reader
}
