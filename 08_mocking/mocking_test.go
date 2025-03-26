package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type configSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *configSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const (
	sleep = "sleep"
	write = "write"
)

type SpyCountdownOpertations struct {
	Calls []string
}

func (s *SpyCountdownOpertations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOpertations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
		}
	})

	t.Run("sleeps before each write", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOpertations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{write, sleep, write, sleep, write, sleep, write}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := configSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should sleep %v but slept %v", sleepTime, spyTime.durationSlept)
	}

}
