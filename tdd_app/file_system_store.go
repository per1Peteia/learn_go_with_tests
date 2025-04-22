package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	db     io.ReadWriteSeeker
	league League
}

func NewFileSystemPlayerStore(db io.ReadWriteSeeker) *FileSystemPlayerStore {
	db.Seek(0, io.SeekStart)
	league, _ := NewLeague(db)
	return &FileSystemPlayerStore{
		db:     db,
		league: league,
	}
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.db.Seek(0, io.SeekStart)
	json.NewEncoder(f.db).Encode(&f.league)
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
