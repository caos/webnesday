package main

import (
	"log"
	"net/http"

	"github.com/caos/webnesday/golang/tokensigner/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HelloWorld)
	http.HandleFunc("/sign", handlers.Sign)
	http.HandleFunc("/verify", handlers.Verify)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
