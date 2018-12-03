package web

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// HTTPHandler is the default handler for our web server.
type HTTPHandler struct {
	name         string
	routehandler map[string]RouteHandler
}

// AttachHandler is a helper function used to attach request handlers to the web server instance.
func (httpHandler *HTTPHandler) AttachHandler(path string, handler RouteHandler) {
	if httpHandler.routehandler == nil {
		httpHandler.routehandler = make(map[string]RouteHandler)
	}
	httpHandler.routehandler[path] = handler
}

// AttachFileSystem attaches FileServer to the web server.
func (httpHandler *HTTPHandler) AttachFileSystem(path string) {
	fs := http.FileServer(http.Dir(path))
	http.Handle(path, http.StripPrefix(path, fs))
}

// ServeHTTP is the default handler type for our HTTPHandler type.
func (httpHandler HTTPHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	currenttime := time.Now()
	fmt.Printf("[INFO] %d:%d: %s '%s'.\n", currenttime.Hour(), currenttime.Minute(), request.Method, request.URL.Path)
	if handler := httpHandler.routehandler[request.URL.Path]; handler != nil {
		handler.HandleRoute(writer, request)
	} else {

		// Try serving the file if is a file, and exists
		filePath := strings.TrimLeft(request.URL.Path, "/")
		if _, status := os.Stat(filePath); os.IsExist(status) {
			http.ServeFile(writer, request, request.URL.Path)
			return
		}

		// Remaining job, to send a response back for an unknown route.
		if path, problem := os.Executable(); problem == nil {
			writer.Write([]byte(fmt.Sprintf("[501] Handler was not found for '%s' route, current working directory is '%s'.", request.URL.Path, path)))
		} else {
			writer.Write([]byte(fmt.Sprintf("[501] Handler was not found for '%s' route.", request.URL.Path)))
		}
	}
}
