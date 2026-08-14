package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// Logout removes the authentication data saved in the user's browser
func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", 302)
}
