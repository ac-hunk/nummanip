package calc

import "testing"

func TestAdd(t *testing.T) {
	want := 16
	if got := Add(2, 2, 3, 4, 5); got != want {
		t.Errorf("Add() = %q, want %q", got, want)
	}
}
