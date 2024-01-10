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

var connectionList = make(map[*websocket.Conn]struct{})

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (websocketManager *Manager) startWebsocket(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started websocket...")

	connection, err := websocketUpgrader.Upgrade(w, r, nil)

	connectionList[connection] = struct{}{}

	if err != nil {
		connection.Close()
		log.Fatalf("Fatal error at websocket connection >>> %v", err)
	}
}

func (websockerManager *Manager) sendNotification(message string) {

	for conn := range connectionList {
		err := conn.WriteJSON(message)

		if err != nil {
			log.Fatalf("Fatal error in send notification from websocket >>> %v", err)
		}
	}
}
