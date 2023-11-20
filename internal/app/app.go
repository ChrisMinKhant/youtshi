package app

import (
	"log"
	"net/http"
	"v1/util"
)

var routeMap = make(map[string]func(w http.ResponseWriter, r *http.Request))

func Init() {
	serverPort := ":8080"

	if util.GetEvnValue("SERVER_PORT") != "" {
		log.Println(" Custom server port was found! ")
		serverPort = util.GetEvnValue("SERVER_PORT")
	}

	route()

	log.Printf("Starting the server and listening at port >>> %s", serverPort)
	http.ListenAndServe(serverPort, nil)
}

func RegisterRoute(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	routeMap[path] = handler
}

func route() {
	for path, handler := range routeMap {
		http.HandleFunc(path, handler)
	}
}
