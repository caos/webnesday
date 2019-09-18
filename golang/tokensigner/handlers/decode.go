package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func decodeRequest(w http.ResponseWriter, body io.ReadCloser, req interface{}) error {
	decoder := json.NewDecoder(body)

	err := decoder.Decode(req)
	if err != nil {
		log.Printf("unable to decode request: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return err
}
