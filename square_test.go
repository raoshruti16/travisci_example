package main

import "testing"

func TestSq(t *testing.T) {
	const a, b = 81,9
	if x := Sq(b); x!= a {
		t.Errorf("Sq(%v) = %v, want %v", b, x, a)
	}
}
