package poker

import (
	"bytes"
	"strings"
	"testing"
)

var dummySpyAlerter = &SpyBlindAlerter{}
var dummyStdOut = bytes.Buffer{}
var dummyStdIn = bytes.Buffer{}
var dummyPlayerStore = &StubPlayerStore{}

type SpyGame struct {
	startedWith  int
	finishedWith string
	startCalled  bool
}

func (s *SpyGame) Start(numberOfPlayers int) {
	s.startCalled = true
	s.startedWith = numberOfPlayers
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
	if game.finishedWith != want {
		t.Errorf("expected Finish call with %q but got %q", want, game.finishedWith)
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
	if game.startedWith != want {
		t.Errorf("expected game to start with %d players but got %d", want, game.startedWith)
	}
}
