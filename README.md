# Go Web Server
A minimal, simple and flexible web server written in Go language. 

## How to start
Repository contains executables for Windows (**run.bat**) and Linux (**run.sh**), these files have the `go run` commands added and you can run these files to start the web server. 

That code in turn executes the web server package's code to start the web server,  
``` go
// Start the web server.
web.InitServer(":5554")
problem := web.StartServer()
```

## How it works
This project uses Go's net packages and http packages to provide an interface to HTTP development. The overall structure somewhat is relatable to Node.js' development stack, of modules, routers and request handlers. 

The HTTPHandler of this web server is basically a normal type defined in Go, as,  
``` go
type HTTPHandler struct {
    name            string
    requestHandlers map[string]RequestHandler
}
```
You attach other handlers for the request to this type, and this is further used down the stream to forward the requests to a specific handler in your own scope. 

### What's next?
Well, as this goes, the project requires an extension API to be written. I also need to expose this web server as a Docker package. Some logging, monitoring and other extensibility packages are needed, and will write them later. 

This package also requires to use the Go language channel feature for concurrency support. 
