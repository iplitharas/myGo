package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known world", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown world", func(t *testing.T) {
		_, err := dictionary.Search("something_else")
		assertError(t, err, ErrNotFound)
	})
	t.Run("add new key", func(t *testing.T) {
		dictionary.Add("new key", "value")
		got, err := dictionary.Search("new key")
		want := "value"
		assertStrings(t, got, want)
		if err != nil {
			t.Fatal("should find added word:", err)
		}

	})
	t.Run("existing key", func(t *testing.T) {
		err := dictionary.Add("test", "new value should not overwrite the existing")
		if err != nil {
			assertError(t, err, ErrExists)
		}
		got, err := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
		if err != nil {
			t.Fatal("should find added word:", err)
		}

	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
