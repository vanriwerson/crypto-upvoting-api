package routes

import (
	"api/src/controllers"
	"net/http"
)

var cryptoRoutes = []Route{
	{
		URI:          "/cryptos",
		Method:       http.MethodPost,
		Controller:   controllers.CreateCrypto,
		RequiresAuth: true,
	},
	{
		URI:          "/cryptos",
		Method:       http.MethodGet,
		Controller:   controllers.FindAllCryptos,
		RequiresAuth: false,
	},
	{
		URI:          "/cryptos/{id}",
		Method:       http.MethodGet,
		Controller:   controllers.FindCryptoById,
		RequiresAuth: false,
	},
	{
		URI:          "/cryptos/{id}",
		Method:       http.MethodPut,
		Controller:   controllers.UpdateCrypto,
		RequiresAuth: true,
	},
	{
		URI:          "/cryptos/{id}",
		Method:       http.MethodDelete,
		Controller:   controllers.DeleteCrypto,
		RequiresAuth: true,
	},
}
