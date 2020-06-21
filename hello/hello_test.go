package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, want, got string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", want, got)
		}
	}

	t.Run("saying hello to a names person", func(t *testing.T) {
		want := "Hello, Freddy"
		got := Hello("Freddy", "")

		assertCorrectMessage(t, want, got)
	})

	t.Run("Saying hello without a names person", func(t *testing.T) {
		want := "Hello, world"
		got := Hello("", "")

		assertCorrectMessage(t, want, got)
	})

	t.Run("In Spanish", func(t *testing.T) {
		want := "Hola, Elodie"
		got := Hello("Elodie", "Spanish")

		assertCorrectMessage(t, want, got)
	})

	t.Run("In French", func(t *testing.T) {
		want := "Bonjour, Elodie"
		got := Hello("Elodie", "French")

		assertCorrectMessage(t, want, got)
	})

	t.Run("In Italian", func(t *testing.T) {
		want := "Buongiorno, Elodie"
		got := Hello("Elodie", "Italian")

		assertCorrectMessage(t, want, got)
	})
}
