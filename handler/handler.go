package handler

import (
	"log"
	"net/http"
	"v1/model"
	"v1/util"
)

/*
* All handlers must implement this interface to be
* invoked by exposed paths.
 */
type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

// In this map, the handlers' functions are mapped with associated http methods.
// It is like grouping the the routes to its' related function.
var funcMap = make(map[string]func(w http.ResponseWriter, r *http.Request))

func findTheFunction(requestedMethod string, w http.ResponseWriter, r *http.Request) {
	err := model.Error{}

	log.Printf("Fetched requested method ::: %v", requestedMethod)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Adjust allowed origin as needed

	// Find the function to invoke according to provided http method.
	for method, fun := range funcMap {
		if requestedMethod == method {
			fun(w, r) // Found function is invoked.
			return
		}

		if requestedMethod == "OPTIONS" {

			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

			return
		}
	}

	// When there is no handler is found from
	// above looping. The following is the error response
	// for HANDLER_NOT_FOUND.
	log.Print("Fetched found error ::: HANDLER_NOT_FOUND")

	err.SetErrorCode("HANDLER_NOT_FOUND")
	err.SetStatus(404)
	err.SetErrorMessage("There is no handler for the path or method requested.")

	util.ParseResponse(w, err, 404)
}

// func Response(status *bool, w http.ResponseWriter) {
// 	response := model.SuccessResponse{}
// 	err := model.Error{}

// 	if *status {

// 		response.SetStatus(200)
// 		response.SetMessage("OK")

// 		util.ParseResponse(w, response, 200)

// 	} else {

// 		err.SetErrorCode("INTERNAL_SERVER_ERROR")
// 		err.SetStatus(500)
// 		err.SetErrorMessage("Something went wrong.")

// 		util.ParseResponse(w, response, 500)

// 	}

// }
