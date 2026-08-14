package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger writes request information in terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate checks whether the cookie exists
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := cookies.Read(r); err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		next(w, r)
	}
}
