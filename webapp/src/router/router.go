package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

// New returns a router with all routes configured
func New() *mux.Router {
	r := mux.NewRouter()
	return routes.Setup(r)
}
