package handlers

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"net/http"
)

type signRequest struct {
	Text string `json:"text"`
}

func Sign(w http.ResponseWriter, r *http.Request) {
	req := new(signRequest)
	err := decodeRequest(w, r.Body, req)
	if err != nil {
		return
	}

	hashed := sha256.Sum256([]byte(req.Text))
	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed[:])
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to sign text: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "signated: %x", signature)
}
