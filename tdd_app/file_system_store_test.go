package poker

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	db, cleanDb := createTempFile(t, `
		[{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
	defer cleanDb()

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		t.Fatalf("did expect no error but got: %v", err)
	}
	assertNoError(t, err)

	t.Run("get player score", func(t *testing.T) {
		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEqual(t, got, want)
	})

	t.Run("record a player score", func(t *testing.T) {
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEqual(t, got, want)
	})

	t.Run("record a player score for new player", func(t *testing.T) {
		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEqual(t, got, want)
	})

	t.Run("return league sorted", func(t *testing.T) {
		got := store.GetLeague()
		want := League{{"Chris", 34}, {"Cleo", 10}, {"Pepper", 1}}
		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

func TestEmptyFileSystemStore(t *testing.T) {
	db, cleanDB := createTempFile(t, "")
	defer cleanDB()

	_, err := NewFileSystemPlayerStore(db)
	assertNoError(t, err)
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("error creating temporary file: %v", err)
	}

	tmpFile.Write([]byte(initialData))

	rmvFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, rmvFile
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("expected no error but got one: %v", got)
	}
}

func assertScoreEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
