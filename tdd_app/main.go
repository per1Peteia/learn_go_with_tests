package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store := &FileSystemPlayerStore{db}
	srv := NewPlayerServer(store)

	fmt.Println("serving on port :6969")
	if err := http.ListenAndServe(":6969", srv); err != nil {
		log.Fatalf("could not listen on port :6969 %v", err)
	}
}
