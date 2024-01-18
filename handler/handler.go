package handler

import (
	"log"
	"net/http"
	"v1/model"
	"v1/util"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

var funcMap = make(map[string]func(w http.ResponseWriter, r *http.Request))

func findTheFunction(requestedMethod string, w http.ResponseWriter, r *http.Request) {
	err := model.Error{}

	log.Printf("Fetched requested method >>> %v", requestedMethod)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Adjust allowed origin as needed

	for method, fun := range funcMap {
		if requestedMethod == method {
			fun(w, r)
			return
		}

		if requestedMethod == "OPTIONS" {

			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

			return
		}
	}

	log.Print("Error Found >>> HANDLER_NOT_FOUND")

	err.SetErrorCode("HANDLER_NOT_FOUND")
	err.SetStatus(404)
	err.SetErrorMessage("There is no handler for the address you requested.")

	util.ParseResponse(w, err, 404)
}

// Giving back respective response
func Response(status *bool, w http.ResponseWriter) {
	response := model.SuccessResponse{}
	err := model.Error{}

	if *status {

		response.SetStatus(200)
		response.SetMessage("OK")

		util.ParseResponse(w, response, 200)
	} else {

		err.SetErrorCode("INTERNAL_SERVER_ERROR")
		err.SetStatus(500)
		err.SetErrorMessage("Something went wrong.")

		util.ParseResponse(w, response, 500)
	}

}
