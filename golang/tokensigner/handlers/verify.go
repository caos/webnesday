package handlers

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

type verifyRequest struct {
	Text   string `json:"text"`
	Signed string `json:"signed"`
}

func Verify(w http.ResponseWriter, r *http.Request) {
	req := new(verifyRequest)
	err := decodeRequest(w, r.Body, req)
	if err != nil {
		return
	}

	message := []byte(req.Text)
	signature, _ := hex.DecodeString(req.Signed)
	hashed := sha256.Sum256(message)

	err = rsa.VerifyPKCS1v15(&key.PublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Fprintf(w, "Error from verification: %s\n", err)
		return
	}
	fmt.Fprint(w, "valid token")
}
