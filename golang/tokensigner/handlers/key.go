package handlers

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

const privateKey = "im so secure"

var key *rsa.PrivateKey

func init() {
	var err error
	key, err = rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Panicf("unable to generate private key: %s", err)
	}
}
