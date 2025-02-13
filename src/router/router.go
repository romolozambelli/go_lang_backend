package router

import (
	"api/src/router/path"

	"github.com/gorilla/mux"
)

// Generate will create a router with the routes configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	return path.Setup(r)
}
