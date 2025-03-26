package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})

	t.Run("runs safe concurrently", func(t *testing.T) {
		counter := NewCounter()
		toCount := 1000

		var wg sync.WaitGroup
		wg.Add(toCount)

		for i := 0; i < toCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(t, counter, toCount)
	})
}

func NewCounter() *Counter {
	return &Counter{}
}

func assertCount(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
