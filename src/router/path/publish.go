package path

import (
	"backend/src/controllers"
	"net/http"
)

var pathPost = []Path{

	{
		URI:      "/post",
		Method:   http.MethodPost,
		Function: controllers.CreatePost,
		Auth:     true,
	},
	{
		URI:      "/post",
		Method:   http.MethodGet,
		Function: controllers.GetPost,
		Auth:     true,
	},

	{
		URI:      "/post/{postid}",
		Method:   http.MethodGet,
		Function: controllers.GetPostByID,
		Auth:     true,
	},

	{
		URI:      "/post/{postid}",
		Method:   http.MethodPut,
		Function: controllers.UpdatePost,
		Auth:     true,
	},
	{
		URI:      "/post/{postid}",
		Method:   http.MethodDelete,
		Function: controllers.DeletePost,
		Auth:     true,
	},
	{
		URI:      "/users/{userid}/posts",
		Method:   http.MethodGet,
		Function: controllers.GetUsersPost,
		Auth:     true,
	},
	{
		URI:      "/post/{postid}/like",
		Method:   http.MethodPost,
		Function: controllers.LikePost,
		Auth:     true,
	},
	{
		URI:      "/post/{postid}/unlike",
		Method:   http.MethodPost,
		Function: controllers.UnlikePost,
		Auth:     true,
	},
}
