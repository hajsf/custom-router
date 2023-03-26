// Package main demonstrates how to use the customrouter library
// to create an HTTP server with a custom router that routes incoming requests
// to registered handlers based on their URL path and HTTP method.
package main

import (
	"fmt"
	"modl/customrouter"
	"net/http"
)

func main() {
	// Register a new handler function for the GET method and "/" URL path using the Handle method of GlobalRouter.
	customrouter.GlobalRouter.Handle("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from GET /")
	})

	// Register a new handler function for the POST method and "/" URL path using the Handle method of GlobalRouter.
	customrouter.GlobalRouter.Handle("POST", "/", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body as form data.
		err := r.ParseForm()
		if err != nil {
			// If an error occurs while parsing the form data, write an error response and return.
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Get the values of the "key1" and "key2" parameters from the request body.
		key1 := r.FormValue("key1")
		key2 := r.FormValue("key2")
		// Write a response that includes the values of these parameters.
		fmt.Fprintf(w, "Hello from POST / with key1=%s and key2=%s", key1, key2)
	})

	// Start the HTTP server and listen for incoming requests on port 8080.
	// Use GlobalRouter as the request multiplexer to route incoming requests
	// to registered handlers based on their URL path and HTTP method.
	http.ListenAndServe(":8080", customrouter.GlobalRouter)
}
