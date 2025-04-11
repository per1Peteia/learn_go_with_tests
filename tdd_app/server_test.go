package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
	}
	srv := &PlayerServer{&store}

	t.Run("returns accepted on POST", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		res := httptest.NewRecorder()

		srv.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusAccepted)
	})
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	srv := &PlayerServer{&store}

	t.Run("returns some player's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()
		srv.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("return different score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		srv.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
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

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("status code: got %d, want %d", got, want)
	}
}
