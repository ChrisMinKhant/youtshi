package handler

import (
	"log"
	"net/http"
	"v1/config/serviceprovider"
	"v1/service"
)

type WebsocketHandler struct {
	Id          string
	Name        string
	BusId       int
	PickUpPoint string
}

var websocketService service.WebsocketService

func init() {
	go func() {
		websocketService = serviceprovider.GetService("websocketService").(service.WebsocketService)
	}()
}

// The main entry point of handler
func (websocketHandler *WebsocketHandler) Handle(w http.ResponseWriter, r *http.Request) {

	log.Print("Reached Handle Function.")

	websocketHandler.websocketHandlerGroup()

	findTheFunction(r.Method, w, r)
}

// Grouping the register under the same api name "/users"
func (websocketHandler *WebsocketHandler) websocketHandlerGroup() {
	funcMap["GET"] = websocketHandler.notifyUsers
}

func (websocketHandler *WebsocketHandler) notifyUsers(w http.ResponseWriter, r *http.Request) {

}
