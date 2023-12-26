package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseResponse(w http.ResponseWriter, responsePayload any, httpStatusCode int) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(httpStatusCode)

	log.Printf("Response Payload : %v", responsePayload)

	err := json.NewEncoder(w).Encode(responsePayload)

	if err != nil {
		log.Fatalf("There was an error encoding json : %v", err)
	}
}
