package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(15, 15)

	if total != 30 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 30)
	}
}
