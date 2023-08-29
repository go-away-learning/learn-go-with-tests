package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("says hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("says hello to people, in spanish", func(t *testing.T) {
		got := Hello("Cristiano", "spanish")
		want := "Hola, Cristiano!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("says hello to people, in french", func(t *testing.T) {
		got := Hello("Chrétien", "french")
		want := "Bonjour, Chrétien!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
