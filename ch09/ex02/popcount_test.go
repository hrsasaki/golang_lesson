package popcount

import "testing"

func TestPopcount(t *testing.T) {
	expected := 3
	actual := PopCount(11)
	if expected != actual {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
