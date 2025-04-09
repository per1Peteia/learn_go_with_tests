package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		aStackOfInts := new(Stack[int])

		AssertTrue(t, aStackOfInts.IsEmpty())

		aStackOfInts.Push(123)
		AssertFalse(t, aStackOfInts.IsEmpty())

		aStackOfInts.Push(456)
		val, _ := aStackOfInts.Pop()
		AssertEqual(t, val, 456)
		val, _ = aStackOfInts.Pop()
		AssertEqual(t, val, 123)

		AssertTrue(t, aStackOfInts.IsEmpty())
	})
}

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "bye")
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("did not want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
