package routes

import (
	"api/src/middlewares"
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
	var routes []Route

  routes = append(routes, userRoutes...)
	routes = append(routes, loginRoute)
	routes = append(routes, postRoutes...)

  for _, route := range routes {
		if route.IsAuthRequired {
			r.HandleFunc(route.URI, 
			middlewares.Logger(
				middlewares.Authenticate(route.Func),
			),
		).Methods(route.Method)
		} else {
    	r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
  }

  return r
}
