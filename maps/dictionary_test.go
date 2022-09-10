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

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"Hello": "World"}

	t.Run("should add new words", func(t *testing.T) {
		key := "NewWord"
		value := "Added"
		err := dictionary.Add(key, value)

		assertDefinition(t, dictionary, key, value)
		assertError(t, err, nil)
	})

	t.Run("should not add existing words", func(t *testing.T) {
		key := "NewWord"
		value := "Added"
		err := dictionary.Add(key, value)

		assertDefinition(t, dictionary, key, value)
		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {

	key := "Hello"
	value := "World"

	t.Run("should update existing word", func(t *testing.T) {

		dictionary := Dictionary{key: value}
		newValue := "World!"

		err := dictionary.Update(key, newValue)

		assertDefinition(t, dictionary, key, newValue)
		assertError(t, err, nil)
	})

	t.Run("should not update new word", func(t *testing.T) {

		dictionary := Dictionary{}

		err := dictionary.Update(key, value)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != value {
		t.Errorf("\nActual: %q\nExpected: %q", got, value)
	}
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
