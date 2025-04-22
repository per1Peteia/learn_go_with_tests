package main

import (
	"fmt"
	poker "github.com/per1Peteia/learn_go_with_tests/tdd_app"
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

	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("could not create filesystem store: %v", err)
	}
	srv := poker.NewPlayerServer(store)

	fmt.Println("serving on port :6969")
	if err := http.ListenAndServe(":6969", srv); err != nil {
		log.Fatalf("could not listen on port :6969 %v", err)
	}
}
