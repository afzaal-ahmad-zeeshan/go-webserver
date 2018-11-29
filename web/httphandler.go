package web

import (
	"fmt"
	"net/http"
	"time"
)

// HTTPHandler is the default handler for our web server.
type HTTPHandler struct {
	name            string
	requestHandlers map[string]RequestHandler
}

// AttachHandler is a helper function used to attach request handlers to the web server instance.
func (httpHandler *HTTPHandler) AttachHandler(path string, handler RequestHandler) {
	if httpHandler.requestHandlers == nil {
		httpHandler.requestHandlers = make(map[string]RequestHandler)
	}
	httpHandler.requestHandlers[path] = handler
}

// ServeHTTP is the default handler type for our HTTPHandler type.
func (httpHandler HTTPHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	currenttime := time.Now()
	fmt.Printf("[INFO] %d:%d: %s '%s'.\n", currenttime.Hour(), currenttime.Minute(), request.Method, request.URL.Path)
	if handler := httpHandler.requestHandlers[request.URL.Path]; handler != nil {
		handler.Handle(writer, request)
	} else {
		writer.Write([]byte("[404] Handler not found."))
	}
}
