package handlers

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

//verifyRequest is the definition of our request body. Because the first letter is lower it's a private struct (you cannot access it from outside this package)
// Text is the value you passed into the sign request
// Signed is the response of the sign request
type verifyRequest struct {
	Text   string `json:"text"`
	Signed string `json:"signed"`
}

//Verify checks if the signed string is correctly signed
func Verify(w http.ResponseWriter, r *http.Request) {
	req := new(verifyRequest)
	err := decodeRequest(w, r.Body, req)
	if err != nil {
		return
	}

	// create a byte array from the string
	message := []byte(req.Text)
	// decode string returns the signature and an error, but we ignore the error this time with the "_"
	signature, _ := hex.DecodeString(req.Signed)
	hashed := sha256.Sum256(message)

	err = rsa.VerifyPKCS1v15(&key.PublicKey, crypto.SHA256, hashed[:], signature)
	// handle error and print it into the response
	if err != nil {
		fmt.Fprintf(w, "Error from verification: %s\n", err)
		return
	}
	//if there is no error the signed string is valid and we can print this into the response
	fmt.Fprint(w, "valid token")
}
