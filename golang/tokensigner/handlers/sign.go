package handlers

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"net/http"
)

// signRequest is the definition of our request body. Because the first letter is lower it's a private struct (you cannot access it from outside this package)
// It only needs a text
type signRequest struct {
	Text string `json:"text"` // `json:"text"` is a so called tag, tags are used to add more information on a field.
}

//Sign handels the requests if /sign is called
func Sign(w http.ResponseWriter, r *http.Request) {
	// initialize req with a pointer to a signRequest
	req := new(signRequest)
	err := decodeRequest(w, r.Body, req)
	if err != nil {
		return
	}

	// hash the text from the request
	hashed := sha256.Sum256([]byte(req.Text))
	// sign the hashed text (hashed[:] makes a slice from an array)
	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed[:])
	// handle error and print it into the response
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to sign text: %s", err), http.StatusInternalServerError)
		return
	}
	// print the signed text into the response
	fmt.Fprintf(w, "signated: %x", signature)
}
