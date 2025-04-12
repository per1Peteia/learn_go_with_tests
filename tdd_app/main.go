package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	srv := NewPlayerServer(NewInMemoryPlayerStore())
	fmt.Println("serving on port :6969")
	log.Fatal(http.ListenAndServe(":6969", srv))
}
