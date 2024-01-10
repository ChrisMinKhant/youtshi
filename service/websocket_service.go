package service

import (
	"net/http"
)

type WebsocketService struct {
}

func (websocketService *WebsocketService) pushNotification(w http.ResponseWriter, r *http.Request) {
	// create websocket manager
	websocketManager := NewManager()

	// start websocket
	websocketManager.startWebsocket(w, r)
}
