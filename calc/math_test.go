package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Logf("Testing TestAdd")
	want := 16
	if _, got := Add(2, 2, 3, 4, 5); got != want {
		t.Errorf("Add() = %q, want %q", got, want)
	}
}

func TestAddError(t *testing.T) {
	t.Logf("Testing TestAddError")
	if err, _ := Add(2); err != nil {
		t.Logf("Less than 2 parameter passed")
	}
}
