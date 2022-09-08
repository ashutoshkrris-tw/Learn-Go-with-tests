package main

import "testing"

func TestHelloUser(t *testing.T) {

	t.Run("should say hello to users", func(t *testing.T) {
		got := Hello("Ashutosh", "")
		wanted := "Hello, Ashutosh"

		assertCorrectMessage(t, got, wanted)
	})

	t.Run("should say 'Hello, World' when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		wanted := "Hello, World"

		assertCorrectMessage(t, got, wanted)
	})

	t.Run("should greet in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		wanted := "Hola, Elodie"

		assertCorrectMessage(t, got, wanted)
	})
	
	t.Run("should greet in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		wanted := "Bonjour, Elodie"

		assertCorrectMessage(t, got, wanted)
	})

	t.Run("should greet in Hindi", func(t *testing.T) {
		got := Hello("Elodie", "Hindi")
		wanted := "Hello, Elodie"

		assertCorrectMessage(t, got, wanted)
	})
}

func assertCorrectMessage(t testing.TB, got string, wanted string) {
	t.Helper()
	if got != wanted {
		t.Errorf("\nActual: %q\nExpected: %q", got, wanted)
	}
}
