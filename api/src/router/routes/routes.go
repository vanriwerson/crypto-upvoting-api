package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI          string
	Method       string
	Controller   func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}

// SetRoutes cria as funções handler para cada rota criada
func SetRoutes(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, cryptoRoutes...)
	routes = append(routes, voteRoutes...)

	for _, route := range routes {
		if route.RequiresAuth { // aplicação dos middlewares por aninhamento para rotas protegidas
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Controller)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Controller)).Methods(route.Method)
		}
	}

	return r
}
