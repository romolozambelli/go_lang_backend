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
}
