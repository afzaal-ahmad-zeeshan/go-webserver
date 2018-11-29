package handlers

import (
	"fmt"
	"net/http"
)

// PersonAPIHandler is the API handle for /person
type PersonAPIHandler struct {
}

// Handle function handles the request for Person API.
func (handler PersonAPIHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	message := fmt.Sprintf("[PersonAPI] %s '%s'", request.Method, request.URL.Path)
	writer.Write([]byte(message))
}
