package routes

import (
	"net/http"

	"api/src/controllers"
)

var postRoutes = []Route{
	{
		URI:            "/posts",
		Method:         http.MethodPost,
		Func:           controllers.CreatePost,
		IsAuthRequired: true,
	},
	{
		URI:            "/posts",
		Method:         http.MethodGet,
		Func:           controllers.GetPosts,
		IsAuthRequired: true,
	},
	{
		URI:            "/posts/{id}",
		Method:         http.MethodGet,
		Func:           controllers.GetPost,
		IsAuthRequired: true,
	},
	{
		URI:            "/posts/{id}",
		Method:         http.MethodPut,
		Func:           controllers.EditPost,
		IsAuthRequired: true,
	},
	{
		URI:            "/posts/{id}",
		Method:         http.MethodDelete,
		Func:           controllers.DeletePost,
		IsAuthRequired: true,
	},
	{
		URI:            "/users/{userId}/posts",
		Method:         http.MethodGet,
		Func:           controllers.GetUserPosts,
		IsAuthRequired: true,
	},
	{
		URI:            "/posts/{id}/upvote",
		Method:         http.MethodPost,
		Func:           controllers.Upvote,
		IsAuthRequired: true,
	},
	{
		URI:            "/posts/{id}/downvote",
		Method:         http.MethodPost,
		Func:           controllers.Downvote,
		IsAuthRequired: true,
	},
}
