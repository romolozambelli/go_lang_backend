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
}
