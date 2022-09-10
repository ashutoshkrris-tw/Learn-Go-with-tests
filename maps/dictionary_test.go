package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"A": 65}

	got := dictionary.Search("A")
	wanted := 65

	assertEqual(t, got, wanted)

}

func assertEqual(t testing.TB, got, wanted int) {
	t.Helper()
	if got != wanted {
		t.Errorf("\nActual: %d\nExpected: %d", got, wanted)
	}
}
