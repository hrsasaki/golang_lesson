package ex19

import "testing"

func Test(t *testing.T) {
	println(ex19())
	if ex19() != 1 {
		t.Error()
	}
}
