package web

import "net/http"

// RequestHandler interface specifies how a function can be used to handle a request.
type RequestHandler interface {
	Handle(writer http.ResponseWriter, request *http.Request)
}

// RouteHandler is the handler for a specific route in the application.
type RouteHandler struct {
	path     string
	handlers []RequestHandler
}
