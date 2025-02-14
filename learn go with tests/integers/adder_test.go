package main

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expeced := 4

	if sum != expeced {
		t.Errorf("expected '%d' but got '%d'", expeced, sum)
	}
}
