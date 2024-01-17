package handlerprovider

import (
	"log"
	"v1/handler"
)

/*
* ' BindHandler() ' is the place where you can
* bind all the handlers which implement Handler interface.
**/

func BindHandler() {

	/*
	* Bind Handler
	**/
	log.Print("Binding handler...")
	RegisterHandler("/notify", handler.NewNotifyHandler())
	RegisterHandler("/ws", handler.NewWebsocketHandler())
}
