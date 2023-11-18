package app

import (
	"log"
	"net/http"
	"v1/internal/controller"
)

func Init() {
	log.Printf("Starting the server and listening at port >>> %s", "8080")
	http.HandleFunc("/", controller.HandleTest)
	http.HandleFunc("/sec", controller.HandleSecondTest)

	http.ListenAndServe(":8080", nil)
}
