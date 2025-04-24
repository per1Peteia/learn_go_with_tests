package poker

import (
	"bytes"
	"strings"
	"testing"
)

var dummySpyAlerter = &SpyBlindAlerter{}
var dummyStdOut = bytes.Buffer{}
var dummyPlayerStore = &StubPlayerStore{}

func TestCLI(t *testing.T) {
	t.Run("chris wins input", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		store := &StubPlayerStore{}
		game := NewGame(dummySpyAlerter, store)

		cli := NewCLI(in, &dummyStdOut, game)
		cli.PlayPoker()

		AssertPlayerWin(t, store, "Chris")
	})

	t.Run("cleo wins input", func(t *testing.T) {
		in := strings.NewReader("1\nCleo wins\n")
		store := &StubPlayerStore{}
		game := NewGame(dummySpyAlerter, store)

		cli := NewCLI(in, &dummyStdOut, game)
		cli.PlayPoker()

		AssertPlayerWin(t, store, "Cleo")
	})
}
