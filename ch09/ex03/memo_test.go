package ex03

import (
	"fmt"
	"testing"

	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := New(Func{httpGetBody, nil})
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := New(Func{httpGetBody, nil})
	defer m.Close()
	memotest.Concurrent(t, m)
}

func TestCancelGet(t *testing.T) {
	done := make(chan struct{})
	m := New(Func{httpGetBody, done})
	defer m.Close()
	fmt.Println(m.CancelableGet("https://godoc.org"))
}
