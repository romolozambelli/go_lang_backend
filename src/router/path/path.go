package path

import (
	"backend/src/midleware"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct to define the pattern for the paths
type Path struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

// Setup the paths on the routers
func Setup(r *mux.Router) *mux.Router {
	paths := pathUsers
	paths = append(paths, routeLogin)
	paths = append(paths, pathPost...)

	for _, path := range paths {

		if path.Auth {
			r.HandleFunc(path.URI, midleware.Logger(midleware.Authenticate(path.Function))).Methods(path.Method)
		} else {
			r.HandleFunc(path.URI, midleware.Logger(path.Function)).Methods(path.Method)
		}
	}

	return r
}
