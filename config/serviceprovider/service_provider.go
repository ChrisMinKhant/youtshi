package serviceprovider

import (
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

	RegisterService("notifyService", service.NotifyService{})
	RegisterService("websocketService", service.WebsocketService{})
}
