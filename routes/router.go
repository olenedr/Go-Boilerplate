package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	// Request method (GET, POST, PUT, DELETE etc)
	Method string
	// Path for the request ("/user")
	Path string
	// A handle method which the router can send the request to
	Handle httprouter.Handle
}

// A map of all the routes
var routes = make(map[int][]Route)

func Add(method, path string, handle httprouter.Handle) {
	AddHandler(0, method, path, handle)
}

// Method for adding a new route
// Recieves a version, request method, path for the request, and a controller method
func AddHandler(version int, method, path string, handle httprouter.Handle) {
	// Strip away first / if there is one present
	path = strings.TrimLeft(path, "/")
	if version > 0 {
		// If the path has a version associated with it, we prepend the version to the route: /v1/entries
		path = fmt.Sprintf("/v%d/%s", version, path)
	} else {
		path = "/" + path
	}

	// Append the new route to the slice with the correct version
	routes[version] = append(routes[version], Route{Method: method, Path: path, Handle: handle})
}

// Creates an instance of the router and sets all the routes
func New() *httprouter.Router {
	router := httprouter.New()

	// For every routes, runs the through the "versions"
	for _, rs := range routes {
		// For every route in that "version"
		for _, r := range rs {
			// Print all the added routes to the console
			fmt.Println(r.Method, r.Path)
			// Attach all the routes to the router
			router.Handle(r.Method, r.Path, r.Handle)
		}
	}
	routes = nil
	// Add a public folder for js and css
	router.NotFound = http.FileServer(http.Dir("public"))

	return router
}
