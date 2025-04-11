package main

import (
	"log"
	"net/http"
)

type InMemPlayerStore struct{}

func (i *InMemPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemPlayerStore) RecordWin(name string) {}

func main() {
	srv := &PlayerServer{&InMemPlayerStore{}}
	log.Fatal(http.ListenAndServe(":6969", srv))
}
