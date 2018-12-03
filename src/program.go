package main

import (
	"fmt"
	"os"

	"./web"
)

func main() {
	defer func() {
		problem := recover()

		// There was more likely a problem that has managed to escape previous modules.
		if problem != nil {
			fmt.Println(problem)

			// Return an error.
			os.Exit(1)
		}
	}()

	// Start the web server.
	web.InitServer(":5554")
	problem := web.StartServer()

	// Perhaps a problem was caused.
	if problem != nil {
		fmt.Println(problem)
	}
}
