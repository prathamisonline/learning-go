package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Pratham")

	want := "Hello, Pratham"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
