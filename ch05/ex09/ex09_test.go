package ex09

import (
	"strconv"
	"testing"
)

func TestExpand(t *testing.T) {
	s := expand("foo $bar  baz  $quux $", strconv.Quote)
	if s != "foo \"bar\"  baz  \"quux\" \"\"" {
		t.Errorf("%s", s)
	}
}

func TestExpandJPN(t *testing.T) {
	s := expand("　あああ　$いいい　 ううう", strconv.Quote)
	if s != "　あああ　\"いいい\"　 ううう" {
		t.Errorf("%s", s)
	}
}
