package handlers

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testObject struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func TestDecodeJSON(t *testing.T) {
	//you don't have to use the new-function to initialize an object:
	someObject := testObject{}
	someJSON := strings.NewReader(`
{
	"id": "webnesday",
	"text": "webnesday is awsome!"
}
`)

	// here we pass the string reader and the memory address (pointer) of someObject
	err := decodeJSON(someJSON, &someObject)
	assert.NoError(t, err)
}
