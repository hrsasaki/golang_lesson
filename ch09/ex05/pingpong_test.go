package ex05

import "testing"

func TestExec(t *testing.T) {
	Exec(10000)
}

func BenchmarkExec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Exec(10000)
	}
}
