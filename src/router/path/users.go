package path

import (
	"api/src/router/controllers"
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
		Auth:     false,
	},

	{
		URI:      "/users/{userID}",
		Method:   http.MethodGet,
		Function: controllers.GetUser,
		Auth:     false,
	},

	{
		URI:      "/users/{userID}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		Auth:     false,
	},

	{
		URI:      "/users/{userID}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		Auth:     false,
	},
}
