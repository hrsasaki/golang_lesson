package ex15

import "testing"

func TestMax(t *testing.T) {
	max, err := max(3, 6, 2, 4)
	if max != 6 {
		t.Error()
	} else if err != nil {
		t.Error(err)
	}
}

func TestMaxWithNoArgs(t *testing.T) {
	_, err := max()
	if err == nil {
		t.Error()
	}
}

func TestMin(t *testing.T) {
	min, err := min(3, 6, 2, 4)
	if min != 2 {
		t.Error()
	} else if err != nil {
		t.Error(err)
	}
}

func TestMinWithNoArgs(t *testing.T) {
	_, err := min()
	if err == nil {
		t.Error()
	}
}
