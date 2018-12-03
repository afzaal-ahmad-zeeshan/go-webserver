package web

import "net/http"

// RequestHandler interface specifies how a function can be used to handle a request.
type RequestHandler interface {
	HandleRequest(writer http.ResponseWriter, request *http.Request)
}
