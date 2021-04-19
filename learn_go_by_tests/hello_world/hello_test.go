package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// t.Helper helps report the correct line when test fails
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Domen", "")
		want := "Hello, Domen"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' using empty parameter", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Domen", "ES")
		want := "Hola, Domen"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Domen", "FR")
		want := "Bonjour, Domen"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Slovenian", func(t *testing.T) {
		got := Hello("Domen", "SI")
		want := "Zdravo, Domen"
		assertCorrectMessage(t, got, want)
	})
}
