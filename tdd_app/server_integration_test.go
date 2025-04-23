package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	db, cleanDB := createTempFile(t, "[]")
	defer cleanDB()
	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newGetScoreRequest(player))
		AssertStatus(t, res.Code, http.StatusOK)
		AssertResponseBody(t, res.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newLeagueRequest())
		AssertStatus(t, res.Code, http.StatusOK)

		got := getLeagueFromResponse(t, res.Body)
		want := []Player{{"Pepper", 3}}

		AssertLeague(t, got, want)
	})
}
