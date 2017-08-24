package ex05

import "testing"

func TestExec(t *testing.T) {
	Exec(10000)
}

func BenchmarkExec(b *testing.B) {
	Exec(10000)
}
