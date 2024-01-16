package service

import (
	"net/http"
	"v1/model"
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

func (websocketService *WebsocketService) PushNotification(busNumber int, message string) *model.Error {

	return websocketManager.sendNotification(busNumber, message)
}
