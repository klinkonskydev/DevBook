package routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

// Route represents all API routes
type Route struct {
	URI            string
	Method         string
	Func           func(w http.ResponseWriter, r *http.Request)
	IsAuthRequired bool
}

func Setup(r *mux.Router) *mux.Router {
  routes := userRoutes
	routes = append(routes, loginRoute)

  for _, route := range routes {
    r.HandleFunc(route.URI, route.Func).Methods(route.Method)
  }

  return r
}
