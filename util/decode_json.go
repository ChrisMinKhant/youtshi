package util

import (
	"encoding/json"
	"io"
	"log"
)

func DecodeJson(requestBody io.ReadCloser, dst any) {

	err := json.NewDecoder(requestBody).Decode(dst)

	if err != nil {
		log.Fatalf("Error found while decoding json request >>> %s", err)
	}
}
