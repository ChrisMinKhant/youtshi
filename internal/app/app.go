package app

import (
	"log"
	"net/http"
	"v1/internal/controller"
	"v1/util"
)

func Init() {
	serverPort := ":8080"

	http.HandleFunc("/", controller.HandleTest)
	http.HandleFunc("/sec", controller.HandleSecondTest)

	if util.GetEvnValue("SERVER_PORT") != "" {
		log.Println(" Custom server port was found! ")
		serverPort = util.GetEvnValue("SERVER_PORT")
	}

	log.Printf("Starting the server and listening at port >>> %s", serverPort)
	http.ListenAndServe(serverPort, nil)
}
