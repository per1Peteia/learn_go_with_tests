package main

import (
	"fmt"
	poker "github.com/per1Peteia/learn_go_with_tests/tdd_app"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFunc, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	alerter := poker.BlindAlerterFunc(poker.StdOutAlerter)

	game := poker.NewGame(alerter, store)

	fmt.Println("let's play poker")
	fmt.Println("type '{name} wins' to record a win")

	poker.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()

}
