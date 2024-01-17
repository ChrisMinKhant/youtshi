package serviceprovider

import (
	"log"
	"v1/service"
)

/*
* ' BindHandler() ' is the place where you can
* bind all the handlers which implement Handler interface.
**/

func BindService() {
	/*
	* Bind Handler
	**/
	log.Printf("Binding service...")

	defer RegisterService("notifyService", service.NewNotifyService())
	defer RegisterService("websocketService", service.NewWebsocketService())
}
