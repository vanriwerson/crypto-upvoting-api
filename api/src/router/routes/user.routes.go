package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Controller:   controllers.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Controller:   controllers.FindAllUsers,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{id}",
		Method:       http.MethodGet,
		Controller:   controllers.FindUserById,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{id}",
		Method:       http.MethodPut,
		Controller:   controllers.UpdateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{id}",
		Method:       http.MethodDelete,
		Controller:   controllers.DeleteUser,
		RequiresAuth: false,
	},
}
