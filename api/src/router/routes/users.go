package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI: "/users",
		Method: http.MethodPost,
		Func: controllers.CreateUser,
		IsAuthRequired: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Func: controllers.GetUsers,
		IsAuthRequired: true,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodGet,
		Func: controllers.GetUser,
		IsAuthRequired: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodPut,
		Func: controllers.EditUser,
		IsAuthRequired: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodDelete,
		Func: controllers.DeleteUser,
		IsAuthRequired: false,
	},
}
