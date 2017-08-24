package ex04

import "testing"

func TestPipeline(t *testing.T) {
	Exec(10000)
}

// func TestPipeline2(t *testing.T) {
// 	Exec(2)
// }

func BenchmarkTestPipeline(b *testing.B) {
	Exec(10000)
}

// func BenchmarkTestPipeline2(b *testing.B) {
// 	Exec(2)
// }
