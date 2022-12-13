package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(4, 6)
	want := 10
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSubstract(t *testing.T) {
	got := Subtract(10, 5)
	want := 5
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
