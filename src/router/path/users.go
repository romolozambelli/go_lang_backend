package path

import (
	"backend/src/controllers"
	"net/http"
)

var pathUsers = []Path{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		Auth:     false,
	},

	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controllers.GetUsers,
		Auth:     true,
	},

	{
		URI:      "/users/{userID}",
		Method:   http.MethodGet,
		Function: controllers.GetUserById,
		Auth:     true,
	},

	{
		URI:      "/users/{userID}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		Auth:     true,
	},

	{
		URI:      "/users/{userID}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		Auth:     true,
	},

	{
		URI:      "/users/{userID}/follow",
		Method:   http.MethodPost,
		Function: controllers.FollowUser,
		Auth:     true,
	},

	{
		URI:      "/users/{userID}/unfollow",
		Method:   http.MethodPost,
		Function: controllers.UnFollowUser,
		Auth:     true,
	},

	{
		URI:      "/users/{userID}/followers",
		Method:   http.MethodGet,
		Function: controllers.GetFollowers,
		Auth:     true,
	},

	{
		URI:      "/users/{userID}/following",
		Method:   http.MethodGet,
		Function: controllers.GetFollowingUsers,
		Auth:     true,
	},

	{
		URI:      "/users/{userID}/change-password",
		Method:   http.MethodPost,
		Function: controllers.UpdatePassword,
		Auth:     true,
	},
}
