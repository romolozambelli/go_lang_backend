package path

import (
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

	for _, path := range paths {
		r.HandleFunc(path.URI, path.Function).Methods(path.Method)
	}

	return r
}
