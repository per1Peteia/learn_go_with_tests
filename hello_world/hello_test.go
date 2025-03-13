package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		want := "Hello, Justus!"
		got := Hello("Justus", "")

		if want != got {
			assertCorrectMessage(t, want, got)
		}
	})

	t.Run("in Spanish", func(t *testing.T) {
		want := "Hola, Justus!"
		got := Hello("Justus", "Spanish")

		if want != got {
			assertCorrectMessage(t, want, got)
		}
	})

	t.Run("in French", func(t *testing.T) {
		want := "Bonjour, Justus!"
		got := Hello("Justus", "French")

		if want != got {
			assertCorrectMessage(t, want, got)
		}
	})

	t.Run("defaulting to Hello, World", func(t *testing.T) {
		want := "Hello, World!"
		got := Hello("", "")

		if want != got {
			assertCorrectMessage(t, want, got)
		}
	})
}

func assertCorrectMessage(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("want %s got %s", want, got)
	}
}
