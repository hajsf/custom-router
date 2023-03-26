// Package customrouter provides a custom HTTP router implementation
// that allows registering multiple handlers for the same URL pattern
// with different HTTP methods.
package customrouter

import (
	"net/http"
)

// Route represents a single route in the router.
type Route struct {
	// Method is the HTTP method of the route.
	Method string
	// Path is the URL path pattern of the route.
	Path string
	// Handler is the function that handles requests for the route.
	Handler http.HandlerFunc
}

// Router represents an HTTP router that routes incoming requests
// to registered handlers based on their URL path and HTTP method.
type Router struct {
	// mux is an instance of http.ServeMux used to match incoming requests
	// based on their URL pattern.
	mux *http.ServeMux
	// handlers is a map of URL patterns to maps of HTTP methods to handler functions.
	handlers map[string]map[string]http.HandlerFunc
}

// GlobalRouter is a global instance of Router that can be used by multiple packages.
var GlobalRouter = &Router{}

// init initializes GlobalRouter by calling its NewRouter method.
func init() {
	GlobalRouter.NewRouter()
}

// NewRouter initializes the router by creating a new instance of http.ServeMux
// and a new map for storing handlers. It returns a pointer to the router.
func (r *Router) NewRouter() *Router {
	// Create a new instance of http.ServeMux and store it in the mux field.
	r.mux = http.NewServeMux()
	// Create a new map for storing handlers and store it in the handlers field.
	r.handlers = make(map[string]map[string]http.HandlerFunc)
	// Return a pointer to the router.
	return r
}

// Handle registers a new handler function for the given URL pattern and HTTP method.
// If multiple handlers are registered for the same URL pattern with different HTTP methods,
// incoming requests are dispatched to the appropriate handler based on their HTTP method.
func (r *Router) Handle(method, path string, handler http.HandlerFunc) {
	// Check if there's already an entry in the handlers map for the given URL pattern.
	if _, ok := r.handlers[path]; !ok {
		// If there isn't, create a new map for it and store it in the handlers map.
		r.handlers[path] = make(map[string]http.HandlerFunc)
		// Register a new handler function for the given URL pattern using the HandleFunc method of mux.
		r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			// Check if there's an entry in the handlers map for the request's URL path and HTTP method.
			if handler, ok := r.handlers[req.URL.Path][req.Method]; ok {
				// If there is, call the corresponding handler function.
				handler(w, req)
			} else {
				// Otherwise, return a 404 Not Found response.
				http.NotFound(w, req)
			}
		})
	}
	// Add a new entry to the handlers map for the given URL pattern and HTTP method.
	r.handlers[path][method] = handler
}

// ServeHTTP dispatches incoming requests to registered handlers based on their URL path and HTTP method.
// If no handler is registered for the request's URL path and HTTP method, it returns a 404 Not Found response.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Call the ServeHTTP method of mux to handle the incoming request.
	r.mux.ServeHTTP(w, req)
}
