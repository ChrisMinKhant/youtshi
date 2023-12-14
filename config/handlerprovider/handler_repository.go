package handlerprovider

import "v1/handler"

/*
* As a name ' handler repository ' this is where all
* the binded handler are stored.
* Handlers are binded to their respective path which is like mapping.
 */

var HandlerMap = make(map[string]handler.Handler)

func RegisterHandler(path string, registeredHandler handler.Handler) {
	HandlerMap[path] = registeredHandler
}

func GetHandler() *map[string]handler.Handler {
	return &HandlerMap
}
