package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestGame(t *testing.T) {
	t.Run("returns 200 on /game", func(t *testing.T) {
		store := StubPlayerStore{}
		server, err := NewPlayerServer(&store, dummyGame)
		if err != nil {
			t.Fatalf("error creating server: %v", err)
		}

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("start game with 3 players and declare Ruth the winner", func(t *testing.T) {
		game := &SpyGame{}
		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		winner := "Ruth"
		sock := mustMakeAConnection(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

		defer server.Close()
		defer sock.Close()

		writeSocketMessage(t, sock, "3")
		writeSocketMessage(t, sock, winner)

		time.Sleep(time.Millisecond * 10)
		assertGameStartedWith(t, game, 3)
		assertGameFinishedWith(t, game, winner)

	})
}

func writeSocketMessage(t testing.TB, socket *websocket.Conn, winner string) {
	t.Helper()
	if err := socket.WriteMessage(websocket.TextMessage, []byte(winner)); err != nil {
		t.Fatalf("error writing to websocket connection: %v", err)
	}
}

func mustMakeAConnection(t testing.TB, url string) *websocket.Conn {
	t.Helper()
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open a websocket on %s: %v", url, err)
	}
	return ws
}

func mustMakePlayerServer(t testing.TB, p *StubPlayerStore, g *SpyGame) *PlayerServer {
	t.Helper()
	server, err := NewPlayerServer(p, g)
	if err != nil {
		t.Fatalf("error creating server: %v", err)
	}
	return server
}

func newGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
}

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}
	store := StubPlayerStore{nil, nil, wantedLeague}
	server, _ := NewPlayerServer(&store, dummyGame)

	t.Run("returns 200 on /league", func(t *testing.T) {
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		AssertContentType(t, response, jsonContentType)
		AssertStatus(t, response.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
	})
}

// this function helps parsing response body json into a league struct
func getLeagueFromResponse(t testing.TB, body io.Reader) []Player {
	t.Helper()
	league, err := NewLeague(body)
	if err != nil {
		t.Fatalf("unable to parse response body %q into league struct: %v", body, err)
	}
	return league
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{},
	}
	srv, _ := NewPlayerServer(&store, dummyGame)

	t.Run("returns accepted on POST", func(t *testing.T) {
		player := "Pepper"

		req := newPostWinRequest("Pepper")
		res := httptest.NewRecorder()

		srv.ServeHTTP(res, req)

		AssertStatus(t, res.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, wanted %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("false winner: got %q, want %q", store.winCalls[0], "Pepper")
		}
	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	srv, _ := NewPlayerServer(&store, dummyGame)

	t.Run("returns some player's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()
		srv.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, got, want)
	})

	t.Run("return different score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		srv.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, got, want)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newGetScoreRequest("Apollo")
		res := httptest.NewRecorder()

		srv.ServeHTTP(res, req)
		got := res.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}
