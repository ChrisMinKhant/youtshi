package service

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (websocketManager *Manager) startWebsocket(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started websocket...")

	connection, err := websocketUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatalf("Fatal error at websocket connection >>> %v", err)
	}

	connection.Close()
}
