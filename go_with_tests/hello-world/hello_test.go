package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("world!")
	want := "Hello, world!"
	if got != want {
		//  For tests %q is very useful as it wraps your values in double quotes.
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloYou(t *testing.T) {
	got := Hello("John")
	want := "Hello, John"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("John")
		want := "Hello, John"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
