package service

import (
	"net/http"
)

type WebsocketService struct {
}

func NewWebsocketService() *WebsocketService {
	return &WebsocketService{}
}

// create websocket manager
var websocketManager = NewManager()

var busService = NewBusService()

func (websocketService *WebsocketService) RegisterConnection(w http.ResponseWriter, r *http.Request) {
	// start websocket
	websocketManager.startWebsocket(w, r)
}
