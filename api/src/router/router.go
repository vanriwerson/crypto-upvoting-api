package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// GenerateRoutes entrega as rotas criadas com suas respectivas handlers
func GenerateRoutes() *mux.Router {
	r := mux.NewRouter()

	return routes.SetRoutes(r)
}
