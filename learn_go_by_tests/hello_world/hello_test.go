package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Domen")
	want := "Hello, Domen"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
