package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoutes = []Route{
	{
		URI:            "/create-user",
		Method:         http.MethodGet,
		Func:           controllers.LoadUserRegistrationPage,
		IsAuthRequired: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Func:           controllers.CreateUser,
		IsAuthRequired: false,
	},
	{
		URI:            "/search-users",
		Method:         http.MethodGet,
		Func:           controllers.LoadUsersPage,
		IsAuthRequired: true,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodGet,
		Func:           controllers.LoadUserProfilePage,
		IsAuthRequired: true,
	},
	{
		URI:            "/users/{userId}/unfollow",
		Method:         http.MethodPost,
		Func:           controllers.UnfollowUser,
		IsAuthRequired: true,
	},
	{
		URI:            "/users/{userId}/follow",
		Method:         http.MethodPost,
		Func:           controllers.FollowUser,
		IsAuthRequired: true,
	},
	{
		URI:            "/profile",
		Method:         http.MethodGet,
		Func:           controllers.LoadLoggedUserProfilePage,
		IsAuthRequired: true,
	},
	{
		URI:            "/edit-user",
		Method:         http.MethodGet,
		Func:           controllers.LoadEditUserPage,
		IsAuthRequired: true,
	},
	{
		URI:            "/edit-user",
		Method:         http.MethodPut,
		Func:           controllers.EditUser,
		IsAuthRequired: true,
	},
	{
		URI:            "/update-password",
		Method:         http.MethodGet,
		Func:           controllers.LoadUpdatePasswordPage,
		IsAuthRequired: true,
	},
	{
		URI:            "/update-password",
		Method:         http.MethodPost,
		Func:           controllers.UpdatePassword,
		IsAuthRequired: true,
	},
	{
		URI:            "/delete-user",
		Method:         http.MethodDelete,
		Func:           controllers.DeleteUser,
		IsAuthRequired: true,
	},
}
