package goself

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2, 3) != 6 {
		t.Fatalf("error")
	}
}
