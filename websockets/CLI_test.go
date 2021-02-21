package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/bugg123/learn-go-with-tests/websocket"
)

var dummySpyAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("it prompts the user for number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := dummySpyAlerter

		cli := poker.NewCLI(dummyPlayerStore, in, stdout, blindAlerter)
		cli.PlayPoker()

		got := stdout.String()
		want := poker.PlayerPrompt

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {

				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduleAlert(t, got, want)

			})

		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("5\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, dummyStdOut, blindAlerter)
		cli.PlayPoker()

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}
		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduleAlert(t, got, want)

			})
		}

	})
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummyStdOut, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummyStdOut, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}

func assertScheduleAlert(t *testing.T, got, want scheduledAlert) {
	amountGot := got.amount

	if amountGot != want.amount {
		t.Errorf("got amount %d, want %d", amountGot, want.amount)
	}

	gotScheduledTime := got.scheduledAt
	if gotScheduledTime != want.scheduledAt {
		t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, want.scheduledAt)
	}
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

type scheduledAlert struct {
	scheduledAt time.Duration
	amount      int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.scheduledAt)
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}
