package transform

import "testing"

func TestPrint(t *testing.T) {
	want := "Integer is  4"
	if got := Print(4); got != want {
		t.Errorf("Transform() = %q, want %q", got, want)
	}
}
