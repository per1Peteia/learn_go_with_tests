package poker

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"
)

type SpyGame struct {
	startedWith    int
	startCalled    bool
	finishedWith   string
	finishedCalled bool
	blindAlert     []byte
}

func (s *SpyGame) Start(numberOfPlayers int, out io.Writer) {
	s.startCalled = true
	s.startedWith = numberOfPlayers
	out.Write(s.blindAlert)

}

func (s *SpyGame) Finish(winner string) {
	s.finishedWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		in := strings.NewReader("7\n")
		out := &bytes.Buffer{}
		game := &SpyGame{}

		cli := NewCLI(in, out, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, out, PlayerPrompt)
		assertGameStartedWith(t, game, 7)
	})

	t.Run("chris wins input", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		game := &SpyGame{}

		cli := NewCLI(in, &dummyStdOut, game)
		cli.PlayPoker()

		assertGameFinishedWith(t, game, "Chris")
	})

	t.Run("cleo wins input", func(t *testing.T) {
		in := strings.NewReader("1\nCleo wins\n")
		game := &SpyGame{}

		cli := NewCLI(in, &dummyStdOut, game)
		cli.PlayPoker()

		assertGameFinishedWith(t, game, "Cleo")
	})

	t.Run("player prompts with invalid input", func(t *testing.T) {
		in := strings.NewReader("Hi\n")
		out := &bytes.Buffer{}
		game := &SpyGame{}

		cli := NewCLI(in, out, game)
		cli.PlayPoker()

		assertGameStarted(t, game)
		assertMessageSentToUser(t, out, PlayerPrompt, BadPlayerInputErrMsg)
	})
}

func assertMessageSentToUser(t testing.TB, out *bytes.Buffer, msgs ...string) {
	t.Helper()
	got := out.String()
	want := strings.Join(msgs, "")
	if got != want {
		t.Errorf("got stdout %q, want %+v", got, want)
	}
}

func assertGameFinishedWith(t testing.TB, game *SpyGame, want string) {
	t.Helper()

	passed := retryUntil(time.Millisecond*500, func() bool {
		return game.finishedWith == want
	})

	if !passed {
		t.Errorf("expected winner %q, but got %q", want, game.finishedWith)
	}
}

func assertGameStarted(t testing.TB, game *SpyGame) {
	t.Helper()
	if game.startCalled {
		t.Errorf("game should not have been started")
	}
}

func assertGameStartedWith(t testing.TB, game *SpyGame, want int) {
	t.Helper()

	passed := retryUntil(time.Millisecond*500, func() bool {
		return game.startedWith == want
	})

	if !passed {
		t.Errorf("expected finish with %q, but got %q", want, game.startedWith)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}
