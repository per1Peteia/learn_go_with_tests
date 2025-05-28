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

func (w *playerServerSocket) WaitForMsg() string {
	_, msg, err := w.ReadMessage()
	if err != nil {
		log.Printf("error reading from websocket: %v", err)
	}
	return string(msg)
}

func (w *playerServerSocket) Write(p []byte) (n int, err error) {
	err = w.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
