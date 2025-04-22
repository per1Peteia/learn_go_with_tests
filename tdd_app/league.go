package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

// Iterates over League and returns pointer to Player with matching name; returns nil if player name is not found.
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

// Reads from rdr and parses byte-string into a Player slice; returns errors happening while parsing.
func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("error parsing league: %v", err)
	}
	return league, err
}
