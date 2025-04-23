package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}
	store := StubPlayerStore{nil, nil, wantedLeague}
	server := NewPlayerServer(&store)

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
	srv := NewPlayerServer(&store)

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
	srv := NewPlayerServer(&store)

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
