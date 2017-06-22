package ex16

import "testing"

func TestJoin(t *testing.T) {
	if Join(",", "foo", "bar", "baz") != "foo,bar,baz" {
		t.Error()
	}
	if Join("", "あああ") != "あああ" {
		t.Error()
	}
	if Join("") != "" {
		t.Error()
	}
}
