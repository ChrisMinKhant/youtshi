package handlerprovider

import "v1/handler"

var HandlerMap = make(map[string]handler.Handler)

func RegisterHandler(path string, registeredHandler handler.Handler) {
	HandlerMap[path] = registeredHandler
}

func GetHandler() *map[string]handler.Handler {
	return &HandlerMap
}
