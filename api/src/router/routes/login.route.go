package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Route{
	URI:          "/login",
	Method:       http.MethodPost,
	Controller:   controllers.Login,
	RequiresAuth: false,
}
