package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Controller:   controllers.Create,
		RequiresAuth: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Controller:   controllers.FindAll,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{id}",
		Method:       http.MethodGet,
		Controller:   controllers.FindOne,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{id}",
		Method:       http.MethodPut,
		Controller:   controllers.Update,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{id}",
		Method:       http.MethodDelete,
		Controller:   controllers.Delete,
		RequiresAuth: false,
	},
}
