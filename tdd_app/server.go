package poker

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	GetLeague() League
	RecordWin(name string)
}

type PlayerServer struct {
	store    PlayerStore
	template *template.Template
	game     Game
	http.Handler
}

type Player struct {
	Name string `json:"Name"`
	Wins int    `json:"Wins"`
}

const htmlTmplPath = "game.html"

func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
	p := new(PlayerServer)

	tmpl, err := template.ParseFiles(htmlTmplPath)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s: %v", htmlTmplPath, err)
	}

	p.store = store
	p.template = tmpl
	p.game = game
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	router.Handle("/game", http.HandlerFunc(p.gameHandler))
	router.Handle("/ws", http.HandlerFunc(p.socketHandler))
	p.Handler = router
	return p, nil
}

const jsonContentType = "application/json"

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
}

func (p *PlayerServer) socketHandler(w http.ResponseWriter, r *http.Request) {
	conn := newPlayerServerSocket(w, r)

	numberOfPlayersMsg := conn.WaitForMsg()
	numberOfPlayers, _ := strconv.Atoi(string(numberOfPlayersMsg))
	p.game.Start(numberOfPlayers, io.Discard)

	winner := conn.WaitForMsg()
	p.game.Finish(string(winner))
}

func (p *PlayerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	err := p.template.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("error executing html template: %s", err.Error()), http.StatusInternalServerError)
	}
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
