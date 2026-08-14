package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postRoutes = []Route{
	{
		URI:            "/posts",
		Method:         http.MethodPost,
		Func:           controllers.CreatePost,
		IsAuthRequired: true,
	},
	{
		URI:            "/posts/{postId}/upvote",
		Method:         http.MethodPost,
		Func:           controllers.UpvotePost,
		IsAuthRequired: true,
	},
	{
		URI:            "/posts/{postId}/downvote",
		Method:         http.MethodPost,
		Func:           controllers.DownvotePost,
		IsAuthRequired: true,
	},
	{
		URI:            "/posts/{postId}/edit",
		Method:         http.MethodGet,
		Func:           controllers.LoadEditPostPage,
		IsAuthRequired: true,
	},
	{
		URI:            "/posts/{postId}",
		Method:         http.MethodPut,
		Func:           controllers.EditPost,
		IsAuthRequired: true,
	},
	{
		URI:            "/posts/{postId}",
		Method:         http.MethodDelete,
		Func:           controllers.DeletePost,
		IsAuthRequired: true,
	},
}
