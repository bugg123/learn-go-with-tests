package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to corey", func(t *testing.T) {
		got := Hello("Corey")
		want := "Hello Corey"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello World when empty string provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
}
