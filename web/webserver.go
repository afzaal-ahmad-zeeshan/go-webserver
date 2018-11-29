package web

import (
	"fmt"
	"net/http"
	"time"
)

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

	myHandler := HTTPHandler{}
	server := &http.Server{
		Addr:           address,
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Handling HTTP traffic at localhost:5555")
	problem := server.ListenAndServe()
	if problem != nil {
		fmt.Println("Server had a problem starting.\n", problem.Error())
	}
}
