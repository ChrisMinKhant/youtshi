package handler

import (
	"net/http"
	"v1/config/serviceprovider"
	"v1/service"
)

type WebsocketHandler struct {
}

func NewWebsocketHandler() *WebsocketHandler {
	return &WebsocketHandler{}
}

// The main entry point of handler
func (websocketHandler *WebsocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	websocketHandler.websocketHandlerGroup()
	findTheFunction(r.Method, w, r)
}

// Grouping the handlers' functions under the same path.
// The functions are differentiated by its http methods.
func (websocketHandler *WebsocketHandler) websocketHandlerGroup() {

	// All the handler's functions must be mapped with
	// their respective http methods here. One http method can
	// have only one handler function.
	funcMap["GET"] = websocketHandler.notifyUsers
}

// New websocket connection is registered here, for notifying users
// with seemless connection.
func (websocketHandler *WebsocketHandler) notifyUsers(w http.ResponseWriter, r *http.Request) {
	websocketService := serviceprovider.GetService("websocketService").(*service.WebsocketService)

	// Register new websocket connection made from front-end.
	websocketService.RegisterConnection(w, r)
}
