package main

import "testing"

func TestHelloUser(t *testing.T) {
	got := Hello("Ashutosh")
	wanted := "Hello, Ashutosh"

	if got != wanted {
		t.Errorf("\nActual: %q\nExpected: %q", got, wanted)
	}
}
