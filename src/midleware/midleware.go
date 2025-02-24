package midleware

import (
	"backend/src/answer"
	"backend/src/auth"
	"log"
	"net/http"
)

// Function used to log the requests
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s \n", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Check the user making the request is authenticated
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc { //Similar to func (w http.ResponseWrite, r *http.Request)

	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auth.CheckToken(r); erro != nil {
			answer.Erro(w, http.StatusUnauthorized, erro)
			return
		}

		nextFunction(w, r)
	}
}
