package web

import (
	"fmt"
	"net/http"
	"time"
)

// HTTPHandler is the default handler for our web server.
type HTTPHandler struct {
}

// ServeHTTP is the default handler type for our HTTPHandler type.
func (httpHandler HTTPHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	currenttime := time.Now()
	fmt.Printf("[INFO] %d:%d: %s '%s'.\n", currenttime.Hour(), currenttime.Minute(), request.Method, request.URL.Path)
	if request.Method == http.MethodGet {
		writer.Write([]byte("A GET request was captured"))
	} else {
		writer.Write([]byte("A POST request was captured"))
	}
}
