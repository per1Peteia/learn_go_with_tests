package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	want := "Hello, World!"
	got := Hello()

	if want != got {
		t.Errorf("got %s want %s", got, want)
	}
}
