package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("should repeat characters 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("should repeat characters 8 times", func(t *testing.T) {
		repeated := Repeat("a", 8)
		expected := "aaaaaaaa"

		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t testing.TB, got string, wanted string) {
	t.Helper()
	if got != wanted {
		t.Errorf("\nActual: %q\nExpected: %q", got, wanted)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}
