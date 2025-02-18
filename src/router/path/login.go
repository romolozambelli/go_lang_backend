package path

import (
	"backend/src/controllers"
	"net/http"
)

var routeLogin = Path{

	URI:      "/login",
	Method:   http.MethodPost,
	Function: controllers.Login,
	Auth:     false,
}
