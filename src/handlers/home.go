package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler is the API handle for /
type HomeHandler struct {
	// requestHandlers map[string]RequestHandler
}

// HandleRoute function handles the request for Person API.
func (handler HomeHandler) HandleRoute(writer http.ResponseWriter, request *http.Request) {
	message := fmt.Sprintf("[Home] %s '%s'", request.Method, request.URL.Path)
	writer.Write([]byte(message))
}
