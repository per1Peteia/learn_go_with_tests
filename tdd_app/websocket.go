package poker

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type playerServerSocket struct {
	*websocket.Conn
}

func newPlayerServerSocket(w http.ResponseWriter, r *http.Request) *playerServerSocket {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading connection to websocket: %v", err)
	}

	return &playerServerSocket{conn}
}

func (n *playerServerSocket) WaitForMsg() string {
	_, msg, err := n.ReadMessage()
	if err != nil {
		log.Printf("error reading from websocket: %v", err)
	}
	return string(msg)
}
