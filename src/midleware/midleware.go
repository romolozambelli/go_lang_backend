package midleware

import (
	"fmt"
	"log"
	"net/http"
)

// Function used to log the requests
func Logger(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s \n", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Check the user making the request is authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc { //Similar to func (w http.ResponseWrite, r *http.Request)

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("midleware/midleware - Authenticating")
		next(w, r)
	}
}
