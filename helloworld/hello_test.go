package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people",func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Eloide", "Spanish")
		want := "Hola, Eloide!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("François", "French")
		want := "Bonjour, François!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}