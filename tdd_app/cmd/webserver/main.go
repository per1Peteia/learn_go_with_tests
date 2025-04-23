package main

import (
	"fmt"
	poker "github.com/per1Peteia/learn_go_with_tests/tdd_app"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFunc, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	srv := poker.NewPlayerServer(store)

	fmt.Println("serving on port :6969")
	if err := http.ListenAndServe(":6969", srv); err != nil {
		log.Fatalf("could not listen on port :6969 %v", err)
	}
}
