package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseResponse(w http.ResponseWriter, responsePayload any, httpStatusCode int) {
	w.Header().Add("content-type", "application/json")
	w.Header().Add("access-control-allow-origin", "http://localhost:3000")
	w.Header().Add("access-control-allow-methods", "GET, POST, OPTIONS")

	w.WriteHeader(200)

	log.Printf("Response Payload : %v", responsePayload)

	err := json.NewEncoder(w).Encode(responsePayload)

	if err != nil {
		log.Fatalf("There was an error encoding json : %v", err)
	}
}
