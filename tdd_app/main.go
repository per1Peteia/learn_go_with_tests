package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &PlayerServer{}
	log.Fatal(http.ListenAndServe(":6969", srv))
}
