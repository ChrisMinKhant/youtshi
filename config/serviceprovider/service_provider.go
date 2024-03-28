package serviceprovider

import (
	"log"
	"v1/service"
)

// By invoking " BindService " function from
// " main " function, regiteration for all services
// can be done.
func BindService() {

	/*
	* All the implemented services must be
	* registered here with their intended names to
	* be able to called.
	 */
	defer RegisterService("notifyService", service.NewNotifyService())
	defer RegisterService("websocketService", service.NewWebsocketService())
	defer RegisterService("web_socketService", service.NewWeb_SocketService())

	log.Print("All the services are binded.")
}
