package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Corey")
	want := "Hello Corey"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
