package service

import (
	"net/http"
)

type WebsocketService struct {
}

// create websocket manager
var websocketManager = NewManager()

func (websocketService *WebsocketService) RegisterConnection(w http.ResponseWriter, r *http.Request) {
	// start websocket
	websocketManager.startWebsocket(w, r)
}

func (websocketService *WebsocketService) PushNotification(busNumber int, message string) {
	websocketManager.sendNotification(busNumber, message)
}
