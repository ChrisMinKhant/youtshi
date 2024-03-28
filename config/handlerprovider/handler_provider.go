package handlerprovider

import (
	"log"
	"v1/handler"
)

// By invoking " BindHandler " function from
// " main " function, regiteration for all handlers
// can be done.
func BindHandler() {

	/*
	* All the implemented handlers must be
	* registered here with their intended paths to
	* expose.
	 */
	RegisterHandler("/notify", handler.NewNotifyHandler())
	RegisterHandler("/ws", handler.NewWebsocketHandler())

	log.Println("All the handlers are binded.")
}
