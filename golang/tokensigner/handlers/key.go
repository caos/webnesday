package handlers

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

// key is used to sign and verify the texts
var key *rsa.PrivateKey

// init is called if a package (handlers) is imported by another one.
// we import hanlders in our main package.
func init() {
	var err error
	key, err = rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Panicf("unable to generate private key: %s", err)
	}
}
