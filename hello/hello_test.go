package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("David", "")
		want := "Hello David"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello World when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Becky", "spanish")
		want := "Ola! Becky"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Andy", "french")
		want := "Bonjour Andy"

		assertCorrectMessage(t, got, want)
	})
}
