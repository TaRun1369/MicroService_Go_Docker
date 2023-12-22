package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// "context"

// "fmt"

func main() {
	// what is handlefunc
	// when you use HandleFunc to register a handler function for a specific pattern in the DefaultServeMux,
	// you're essentially saying, "When this type of request with a certain pattern comes in,
	// follow these instructions I've provided in the DefaultServeMux to handle it appropriately."
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello")
		data, _ := io.ReadAll(r.Body)
		log.Printf("data - %s", data)

		// write im responseWriter
		fmt.Fprintf(w, "write se mila - %s", data)
	})
	// responsewriter - Think of the http.ResponseWriter as a container or a box. This box is used by your server to package up the information that it wants to send back to the client. Inside this box, you put things like the web page content, images, or any other data that the client requested.
	// request - jo actual request thi wo (put get wo)
	// The * before http.Request indicates that it is a pointer to an instance of the http.Request struct. In Go, passing a pointer allows functions to modify the original data efficiently. So, when you pass *http.Request to your handler function, you're giving it access to all the details of the incoming request.

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye")
	})
	http.ListenAndServe(":9090", nil)
	// :9090 is for all the ip addresses at port 9090 // can be set to some ip as - 127.0.0.1:9090 i.e. localhost ...
}
