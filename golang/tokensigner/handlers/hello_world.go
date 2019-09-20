package handlers

import (
	"fmt"
	"net/http"
)

//HelloWorld writes "Hello webnesday!" into the response
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello webnesday!")
}
