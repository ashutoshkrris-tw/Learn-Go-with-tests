package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"Hello": "World"}

	t.Run("letter in dictionary", func(t *testing.T) {
		got, _ := dictionary.Search("Hello")
		wanted := "World"

		assertEqual(t, got, wanted)
	})

	t.Run("letter not in dictionary", func(t *testing.T) {
		_, err := dictionary.Search("123")

		assertError(t, err, ErrNotFound)
	})
}

func assertEqual(t testing.TB, got, wanted string) {
	t.Helper()
	if got != wanted {
		t.Errorf("\nActual: %q\nExpected: %q", got, wanted)
	}
}

func assertError(t testing.TB, got, wanted error) {
	t.Helper()
	if got != wanted {
		t.Errorf("\nActual: %q\nExpected: %q", got, wanted)
	}
}
