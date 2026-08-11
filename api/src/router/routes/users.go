package routes

import (
	"net/http"

	"api/src/controllers"
)

var userRoutes = []Route{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Func:           controllers.CreateUser,
		IsAuthRequired: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Func:           controllers.GetUsers,
		IsAuthRequired: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodGet,
		Func:           controllers.GetUser,
		IsAuthRequired: true,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodPut,
		Func:           controllers.EditUser,
		IsAuthRequired: true,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodDelete,
		Func:           controllers.DeleteUser,
		IsAuthRequired: true,
	},
	{
		URI:            "/users/{id}/follow",
		Method:         http.MethodPost,
		Func:           controllers.FollowUser,
		IsAuthRequired: true,
	},
	{
		URI:            "/users/{id}/unfollow",
		Method:         http.MethodPost,
		Func:           controllers.UnfollowUser,
		IsAuthRequired: true,
	},
	{
		URI:            "/users/{id}/followers",
		Method:         http.MethodGet,
		Func:           controllers.Followers,
		IsAuthRequired: true,
	},
	{
		URI:            "/users/{id}/following",
		Method:         http.MethodGet,
		Func:           controllers.Followers,
		IsAuthRequired: true,
	},
}
