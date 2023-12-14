package util

import (
	"log"
	"net/http"
)

func Response(w http.ResponseWriter, responsePayload *[]byte) (int, error) {
	w.Header().Add("content-type", "application/json")

	signal, err := w.Write(*responsePayload)

	if err != nil {
		log.Printf("Found error in pacakge pkg -> Response() : %s", err)
		return signal, err
	}

	return signal, err
}
