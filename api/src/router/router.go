package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()
	return routes.Setup(r)
}
