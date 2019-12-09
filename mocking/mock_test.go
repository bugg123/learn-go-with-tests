package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}
func (s *SpySleeper) Sleep(){
	s.Calls++
}
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q but expected %q", got, want)
	}
	if spySleeper.Calls != 4 {
		t.Errorf("not the right calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}