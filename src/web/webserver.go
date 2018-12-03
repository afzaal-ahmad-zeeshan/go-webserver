package web

import (
	"fmt"
	"net/http"
	"time"

	"../handlers"
	"../handlers/api"
)

var server *http.Server

// InitServer starts the web server and adds the specific handlers to it.
func InitServer(address string) {
	defer func() {
		problem := recover()
		if problem != nil {
			// Parse it here
			fmt.Println(problem)
		}
	}()

	if address == "" {
		address = ":5555"
	}

	requestHandler := HTTPHandler{}
	requestHandler.AttachHandler("/", handlers.HomeHandler{})
	requestHandler.AttachHandler("/person", api.PersonAPIHandler{})

	requestHandler.AttachFileSystem(".")

	server = &http.Server{
		Addr:           address,
		Handler:        requestHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

// StartServer starts the web service provided it has been configured.
func StartServer() error {
	fmt.Println("Handling HTTP traffic at localhost" + server.Addr)
	problem := server.ListenAndServe()

	// Handle the problems
	if problem != nil {

		// Log the problem
		fmt.Println("Server had a problem starting.\n", problem.Error())
	}

	// Return it so that other functions can handle it too.
	return problem
}
