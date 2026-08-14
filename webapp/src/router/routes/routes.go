package routes

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Route represents all routes of the Web Application
type Route struct {
	URI            string
	Method         string
	Func           func(http.ResponseWriter, *http.Request)
	IsAuthRequired bool
}

// Setup puts all routes inside the router
func Setup(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, userRoutes...)
	routes = append(routes, homeRoute)
	routes = append(routes, postRoutes...)
	routes = append(routes, logoutRoute)

	for _, route := range routes {

		if route.IsAuthRequired {
			router.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Func)),
			).Methods(route.Method)

		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Func),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
