package poker

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestCLI(t *testing.T) {
	t.Run("chris wins input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &StubPlayerStore{}
		var dummySpyAlerter = &SpyBlindAlerter{}

		cli := NewCLI(store, in, dummySpyAlerter)
		cli.PlayPoker()

		AssertPlayerWin(t, store, "Chris")
	})

	t.Run("cleo wins input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := &StubPlayerStore{}
		var dummySpyAlerter = &SpyBlindAlerter{}

		cli := NewCLI(store, in, dummySpyAlerter)
		cli.PlayPoker()

		AssertPlayerWin(t, store, "Cleo")
	})

	t.Run("schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := NewCLI(store, in, blindAlerter)
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
					t.Errorf("alert %d was not scheduled: %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
}

func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	t.Helper()
	if got.amount != want.amount {
		t.Errorf("got amount %d, want %d", got.amount, want.amount)
	}
	if got.at != want.at {
		t.Errorf("got time %v, want %d", got.at, want.at)
	}
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}
