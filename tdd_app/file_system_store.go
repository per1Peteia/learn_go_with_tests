package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	db     *json.Encoder
	league League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	file.Seek(0, io.SeekStart)
	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("couldn't load player store from %s: %v", file.Name(), err)
	}
	return &FileSystemPlayerStore{
		db:     json.NewEncoder(&tape{file}),
		league: league,
	}, nil
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.db.Encode(&f.league)
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}
