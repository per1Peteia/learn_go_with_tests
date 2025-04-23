package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("chris wins input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &StubPlayerStore{}

		cli := NewCLI(store, in)
		cli.PlayPoker()

		AssertPlayerWin(t, store, "Chris")
	})

	t.Run("cleo wins input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := &StubPlayerStore{}

		cli := NewCLI(store, in)
		cli.PlayPoker()

		AssertPlayerWin(t, store, "Cleo")
	})
}
