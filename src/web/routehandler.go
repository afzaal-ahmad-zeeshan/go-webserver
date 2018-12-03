package web

import "net/http"

// RouteHandler is the handler for a specific route in the application.
type RouteHandler interface {
	HandleRoute(writer http.ResponseWriter, request *http.Request)
}
