package api

import (
	"fmt"
	"net/http"
)

// PersonAPIHandler is the API handle for /person
type PersonAPIHandler struct {
}

// HandleRoute function handles the request for Person API.
func (handler PersonAPIHandler) HandleRoute(writer http.ResponseWriter, request *http.Request) {
	message := fmt.Sprintf("[PersonAPI] %s '%s'", request.Method, request.URL.Path)
	writer.Write([]byte(message))
}
