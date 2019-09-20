package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// decodeRequest decodes the body into req
// ReadCloser is an extension of two interfaces, the Reader and Closer interfaces
// req is an empty interface, every type "implements" an empty interface (also non-pointer values)
func decodeRequest(w http.ResponseWriter, body io.ReadCloser, req interface{}) error {
	err := decodeJSON(body, req)
	// direct error handling (love it or hate it)
	if err != nil {
		log.Printf("unable to decode request: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError) // sets the reply on the the response
	}
	return err
}

func decodeJSON(body io.Reader, obj interface{}) error {
	// create a new decoder of the given body
	decoder := json.NewDecoder(body)

	// decode the body into the given interface (ATTENTION: because req is an empty interface it could also be a non-pointer value)
	return decoder.Decode(obj)
}
